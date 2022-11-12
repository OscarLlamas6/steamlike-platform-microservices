import { Request, Response } from 'express'
import axios from 'axios';
import dotenv from 'dotenv';
dotenv.config();

const { API_DEVELOPERS_BASE_URL } = process.env;

const BASE_URL = API_DEVELOPERS_BASE_URL ? API_DEVELOPERS_BASE_URL : "http://localhost:3004";


// ADDING NEW DEVELOPER
export async function addDeveloper(req: Request, res: Response):Promise<Response | void> {
    try {

        axios.post(`${BASE_URL}/developers/create`,{
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

// GET SINGLE DEVELOPER
export async function getDeveloper(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/developers/${req.params.idDeveloper}`,
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

// GET SINGLE DEVELOPER NO AUTH
export async function getDeveloperNoAuth(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/developers/noauth/${req.params.idDeveloper}`,
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

// GET DEVELOPERS LIST
export async function getDevelopersList(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/developers/list`,
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

// GET DEVELOPERS LIST NO AUTH
export async function getDevelopersListNoAuth(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/developers/list/noauth`,
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

// GET DEVELOPERS LIST v2
export async function getDevelopersListV2(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/developers/`,
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


// UPDATE DEVELOPER REGISTER
export async function updateDeveloper(req: Request, res: Response):Promise<Response | void> {
    
    try {
        
        axios.put(`${BASE_URL}/developers/update`,{
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

// DELETE DEVELOPER REGISTER
export async function deleteDeveloper(req: Request, res: Response):Promise<Response | void>{
    try {
        
        axios.delete(`${BASE_URL}/developers/delete/${req.params.idDeveloper}`,{
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

