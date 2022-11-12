package utils

import (
	"fmt"
	"os"
	"users-service/models"

	encoder "encoding/json"

	"github.com/streadway/amqp"
)

func SendVerifyEmail(email string, userName string, verifyToken string) {

	RABBITMQ_QUEUE := os.Getenv("RABBITMQ_QUEUE")
	RABBITMQ_URL := os.Getenv("RABBITMQ_URL")

	conn, errmq := amqp.Dial(RABBITMQ_URL)
	if errmq != nil {
		fmt.Printf("Error: %v\n", errmq)
	}

	defer conn.Close()

	ch, errmq := conn.Channel()
	if errmq != nil {
		fmt.Printf("Error: %v\n", errmq)
	}
	defer ch.Close()

	// We create a Queue to send the message to.
	q, errmq := ch.QueueDeclare(
		RABBITMQ_QUEUE, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	if errmq != nil {
		fmt.Printf("Error: %v\n", errmq)
	}

	// We set the payload for the message
	body := models.VerifyMail{
		Email:       email,
		UserName:    userName,
		VerifyToken: verifyToken,
	}

	jsonObj, err := encoder.Marshal(body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(jsonObj),
		})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
