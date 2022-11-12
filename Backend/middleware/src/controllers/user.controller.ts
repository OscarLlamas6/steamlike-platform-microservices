import { Request, Response } from 'express'
import axios from 'axios';
import dotenv from 'dotenv';
dotenv.config();

const { API_USERS_BASE_URL } = process.env;

const BASE_URL = API_USERS_BASE_URL ? API_USERS_BASE_URL : "http://localhost:3000";

//GET USER PROFILE
export async function getProfile(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/user/profile/${req.params.username}`,
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

//GET ALL USERS
export async function getAllUsers(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/user/list/all`,
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

//SIGNUP USER
export async function createUser(req: Request, res: Response):Promise<Response | void> {
    try {

        console.log(BASE_URL)
    
        axios.post(`${BASE_URL}/user/create`,{
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

//LOGIN USER
export async function loginUser(req: Request, res: Response):Promise<Response | void> {
 try {

    axios.post(`${BASE_URL}/user/login`,{
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
    
 } catch (e:any) {
    console.log(`${e.message} - ${e.stack}`);
    return res.status(404).json({
        message: `${e.message} - ${e.stack}`,
        status: 404
    });
 }
}

//UPDATE USER INFO
export async function updateUser(req: Request, res: Response):Promise<Response | void> {
    
    try {
        
        axios.put(`${BASE_URL}/user/update`,{
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

//DELETE USER
export async function deleteUser(req: Request, res: Response):Promise<Response | void>{
    try {
        
        axios.delete(`${BASE_URL}/user/delete/${req.params.username}`,{
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

//VERIFY USER
export async function verifyUser(req: Request, res: Response):Promise<Response | void> {
    try {

        axios.get(`${BASE_URL}/user/verify`, {
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

// RESEND 
export async function resendVerify(req: Request, res: Response):Promise<Response | void> {
    try {

        axios.get(`${BASE_URL}/user/resend/verify`, {
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

//RESTORE USER
export async function restoreUser(req: Request, res: Response):Promise<Response | void>{
    try {
        
        axios.get(`${BASE_URL}/user/restore/${req.params.username}`,{
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