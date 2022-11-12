import { Request, Response } from 'express'
import axios from 'axios';
import dotenv from 'dotenv';
dotenv.config();

const { API_CATALOGS_BASE_URL } = process.env;

const BASE_URL = API_CATALOGS_BASE_URL ? API_CATALOGS_BASE_URL : "http://localhost:3005";


// ADDING NEW CATEGORY
export async function addCategory(req: Request, res: Response):Promise<Response | void> {
    try {

        axios.post(`${BASE_URL}/category/create`,{
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

// GET SINGLE CATEGORY
export async function getCategory(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/category/${req.params.idCategory}`,
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

// GET CATEGORIES LIST
export async function getCategoriesList(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/category/list`,
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

// GET CATEGORIES LIST NO AUTH
export async function getCategoriesListNoAuth(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/category/list/noauth`,
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


// UPDATE CATEGORY REGISTER
export async function updateCategory(req: Request, res: Response):Promise<Response | void> {
    
    try {
        
        axios.put(`${BASE_URL}/category/update`,{
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

// DELETE CATEGORY REGISTER
export async function deleteCategory(req: Request, res: Response):Promise<Response | void>{
    try {
        
        axios.delete(`${BASE_URL}/category/delete/${req.params.idCategory}`,{
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

