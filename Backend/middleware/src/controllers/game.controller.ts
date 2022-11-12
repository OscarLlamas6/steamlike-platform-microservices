import { Request, Response } from 'express'
import axios from 'axios';
import dotenv from 'dotenv';
dotenv.config();

const { API_GAMES_BASE_URL } = process.env;

const BASE_URL = API_GAMES_BASE_URL ? API_GAMES_BASE_URL : "http://localhost:3006";


// ADDING NEW GAME
export async function addGame(req: Request, res: Response):Promise<Response | void> {
    try {

        axios.post(`${BASE_URL}/games/create`,{
            ...req.body
        }, {
            headers: {
                        "Content-Type": "application/json",
                        "Authorization": req.header('Authorization') || "",
                        "Accept": "*/*",
                    },
            timeout: 15000,
            })
            .then( (response) => {
            return res.status(response.status).json(response.data);   
        }).catch( (error: any) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                message: `${error.message} - ${error.stack}`,
                status: 404
            }); 
            }                 
        });
    
    } catch (e:any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// GET SINGLE GAME BY ID
export async function getGameByID(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/games/single/${req.params.idGame}`,
            {
                headers: {
                            "Content-Type": "application/json",
                            "Authorization": req.header('Authorization') || "",
                            "Accept": "*/*",
                        },
                timeout: 15000,
            })
        .then( (response) => {
            return res.status(response.status).json(response.data);       
        })
        .catch( (error) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }          
        });      
    }
    catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// GET SINGLE GAME BY ID NO AUTH
export async function getGameByIDNoAuth(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/games/single/noauth/${req.params.idGame}`,
            {
                headers: {
                            "Content-Type": "application/json",
                            "Authorization": req.header('Authorization') || "",
                            "Accept": "*/*",
                        },
                timeout: 15000,
            })
        .then( (response) => {
            return res.status(response.status).json(response.data);       
        })
        .catch( (error) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }          
        });      
    }
    catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}


// GET SINGLE GAME (ADMIN)
export async function getGameAdmin(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/games/info/${req.params.idGame}`,
            {
                headers: {
                            "Content-Type": "application/json",
                            "Authorization": req.header('Authorization') || "",
                            "Accept": "*/*",
                        },
                timeout: 15000,
            })
        .then( (response) => {
            return res.status(response.status).json(response.data);       
        })
        .catch( (error) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }          
        });      
    }
    catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// GET GAMES LIST
export async function getGamesList(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/games/list`,
            {
                headers: {
                            "Content-Type": "application/json",
                            "Authorization": req.header('Authorization') || "",
                            "Accept": "*/*",
                        },
                timeout: 15000,
            })
        .then( (response) => {
            return res.status(response.status).json(response.data);       
        })
        .catch( (error) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }          
        });      
    }
    catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// GET GAMES LIST v2
export async function getGamesListV2(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/games/`,
            {
                headers: {
                            "Content-Type": "application/json",
                            "Authorization": req.header('Authorization') || "",
                            "Accept": "*/*",
                        },
                timeout: 15000,
            })
        .then( (response) => {
            return res.status(response.status).json(response.data);       
        })
        .catch( (error) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }          
        });      
    }
    catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// GET GAMES WITH AGE RESTRICTION
export async function getGamesListByAgeRestriction(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/games/list/age`,
            {
                headers: {
                            "Content-Type": "application/json",
                            "Authorization": req.header('Authorization') || "",
                            "Accept": "*/*",
                        },
                params : {
                    ...req.query
                },
                timeout: 15000,
            })
        .then( (response) => {
            return res.status(response.status).json(response.data);       
        })
        .catch( (error) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }          
        });      
    }
    catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// GET ALL GAMES
export async function getAllGames(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/games/list/all`,
            {
                headers: {
                            "Content-Type": "application/json",
                            "Authorization": req.header('Authorization') || "",
                            "Accept": "*/*",
                        },
                params : {
                    ...req.query
                },
                timeout: 15000,
            })
        .then( (response) => {
            return res.status(response.status).json(response.data);       
        })
        .catch( (error) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }          
        });      
    }
    catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// GET GAMES FOR HOME PAGE
export async function getGameForHomePage(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/games/list/home`,
            {
                headers: {
                            "Content-Type": "application/json",
                            "Authorization": req.header('Authorization') || "",
                            "Accept": "*/*",
                        },
                params : {
                    ...req.query
                },
                timeout: 15000,
            })
        .then( (response) => {
            return res.status(response.status).json(response.data);       
        })
        .catch( (error) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }          
        });      
    }
    catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// UPDATE GAME REGISTER
export async function updateGame(req: Request, res: Response):Promise<Response | void> {
    
    try {
        
        axios.put(`${BASE_URL}/games/update`,{
            ...req.body
        },{
            headers: {
                        "Content-Type": "application/json",
                        "Authorization": req.header('Authorization') || "",
                        "Accept": "*/*",
                    },
            timeout: 15000,
            }).then( (response) => {
                return res.status(response.status).json(response.data);   
        }).catch( (error) => {
    
            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }                          
        });
            
    } catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// DELETE GAME REGISTER
export async function deleteGame(req: Request, res: Response):Promise<Response | void>{
    try {
        
        axios.delete(`${BASE_URL}/games/delete/${req.params.idGame}`, {
            headers: {
                        "Content-Type": "application/json",
                        "Authorization": req.header('Authorization') || "",
                        "Accept": "*/*",
                    },
            timeout: 15000,
            })
            .then( (response) => {
                return res.status(response.status).json(response.data);      
            })
            .catch( (error) => {

                if (error.response) {
                    res.status(error.response.status);
                    return res.json(error.response.data);              
                } else {
                    return res.status(404).json({
                        message: `${error.message} - ${error.stack}`,
                        status: 404
                    }); 
                } 
            
            });

    } catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

// RESTORE GAME
export async function restoreGame(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/games/restore/${req.params.idGame}`,
            {
                headers: {
                            "Content-Type": "application/json",
                            "Authorization": req.header('Authorization') || "",
                            "Accept": "*/*",
                        },
                timeout: 15000,
            })
        .then( (response) => {
            return res.status(response.status).json(response.data);       
        })
        .catch( (error) => {

            if (error.response) {
                res.status(error.response.status);
                return res.json(error.response.data);              
            } else {
                return res.status(404).json({
                    message: `${error.message} - ${error.stack}`,
                    status: 404
                }); 
            }          
        });      
    }
    catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}