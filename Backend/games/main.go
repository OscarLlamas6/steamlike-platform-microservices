package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"games-service/configs"
	"games-service/controllers"
	"games-service/models"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Saludo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Microservicio Juegos - Steamlike Platform | Grupo 4 :D"})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	//os.Setenv("IS_DEV", "TRUE")
	IS_DEV := os.Getenv("IS_DEV")

	if IS_DEV == "TRUE" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	// Se crea el servidor con GIN
	r := gin.Default()

	// Se aplican middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	//Rutas
	var userName = "userName"

	var jwtKEY string = os.Getenv("JWT_TOKEN_KEY")
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(jwtKEY),
		Timeout:     6 * time.Hour,
		MaxRefresh:  6 * time.Hour,
		IdentityKey: userName,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.UserPayload); ok {
				return jwt.MapClaims{
					userName: v.UserName,
					"group":  4,
					"userId": v.Id,
					"email":  v.Email,
					"name":   v.Name + " " + v.LastName,
					"region": v.Region,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.UserPayload{
				UserName: claims[userName].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {

			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			hashString := []byte(loginVals.Password)
			encodedPass := fmt.Sprintf("%x", md5.Sum(hashString))

			db := configs.ConnectDB()

			// Verificando si las credenciales coinciden con algun usuario

			userQuery, err := db.Query("SELECT * FROM User WHERE username = ? AND password = ?", userID, encodedPass)
			if err != nil {
				defer db.Close()
				fmt.Println(err)
				return nil, jwt.ErrFailedAuthentication
			}

			if !userQuery.Next() {
				defer db.Close()
				fmt.Println(err)
				return nil, errors.New("credenciales incorrectas")
			}

			// Verificando si la cuenta ya ha sido verificada o no

			verifyQuery, err := db.Query("SELECT * FROM User WHERE username = ? AND password = ? AND `User`.isActive = 1;", userID, encodedPass)
			if err != nil {
				defer db.Close()
				fmt.Println(err)
				return nil, jwt.ErrFailedAuthentication
			}

			if !verifyQuery.Next() {
				defer db.Close()
				fmt.Println(err)
				return nil, errors.New("la cuenta no ha sido verificada, revise su correo")
			}

			rows, err := db.Query("SELECT idUser, `User`.`name`, LastName, username, email, `Region`.`name` FROM User INNER JOIN Region ON User.idRegion = Region.idRegion WHERE username = ? AND password = ? AND `User`.isDeleted = 0 AND `User`.isActive = 1;", userID, encodedPass)
			if err != nil {
				defer db.Close()
				fmt.Println(err)
				return nil, jwt.ErrFailedAuthentication
			}

			if rows.Next() {

				var id int64

				var name, lastName, username, email, region string

				err = rows.Scan(&id, &name, &lastName, &username, &email, &region)
				if err != nil {
					defer db.Close()
					return nil, jwt.ErrFailedAuthentication
				}

				defer db.Close()
				return &models.UserPayload{
					Id:       id,
					Name:     name,
					LastName: lastName,
					UserName: userID,
					Email:    email,
					Region:   region,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// if v, ok := data.(*UserTest); ok && v.UserName == "admin" {
			// 	return true
			// }

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.GET("/", Saludo)

	game := r.Group("/games")
	{
		game.POST("/create", authMiddleware.MiddlewareFunc(), controllers.CreateGame())
		game.PUT("/update", authMiddleware.MiddlewareFunc(), controllers.UpdateGame())
		game.GET("/", authMiddleware.MiddlewareFunc(), controllers.GetGamesv2())
		game.GET("/list/home", controllers.GetGamesNoAuth())
		game.GET("/single/:idGame", authMiddleware.MiddlewareFunc(), controllers.GetGame())
		game.GET("/single/noauth/:idGame", controllers.GetGame())
		game.GET("/info/:idGame", authMiddleware.MiddlewareFunc(), controllers.GetGameAdmin())
		game.GET("/list", authMiddleware.MiddlewareFunc(), controllers.GetGames())
		game.GET("/list/all", authMiddleware.MiddlewareFunc(), controllers.GetAllGames())
		game.GET("/list/age", authMiddleware.MiddlewareFunc(), controllers.GetGamesByAgeRestriction())
		game.DELETE("/delete/:idGame", authMiddleware.MiddlewareFunc(), controllers.DeleteGame())
		game.GET("/restore/:idGame", authMiddleware.MiddlewareFunc(), controllers.RestoreGame())
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	HOST_PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(HOST_PORT)
}
