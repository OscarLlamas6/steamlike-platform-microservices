import { Request, Response } from 'express'
import axios from 'axios';
import dotenv from 'dotenv';
dotenv.config();

const { API_DLC_BASE_URL } = process.env;

const BASE_URL = API_DLC_BASE_URL ? API_DLC_BASE_URL : "http://localhost:3007";


// ADDING NEW DLC
export async function addDLC(req: Request, res: Response):Promise<Response | void> {
    try {

        axios.post(`${BASE_URL}/dlc/create`,{
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

// GET SINGLE DLC
export async function getDLC(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/dlc/single/${req.params.idDLC}`,
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

// GET SINGLE DLC NO AUTH
export async function getDLCNoAuth(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/dlc/single/noauth/${req.params.idDLC}`,
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

// GET SINGLE DLC (ADMIN)
export async function getDLCAdmin(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/dlc/info/${req.params.idDLC}`,
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

// GET DLCs LIST
export async function getDLCList(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/dlc/list`,
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

// GET DLCs LIST BY GAME
export async function getDLCListByGame(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/dlc/game/list/${req.params.idGame}`,
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

// GET DLCs LIST BY GAME (NO AUTH)
export async function getDLCListByGameNoAuth(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/dlc/game/list/${req.params.idGame}/noauth`,
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

// GET DLCs LIST NO AUTH
export async function getDLCListNoAuth(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/dlc/list/noauth`,
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


// GET ALL DLC
export async function getAllDLC(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/dlc/list/all`,
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


// UPDATE DLC REGISTER
export async function updateDLC(req: Request, res: Response):Promise<Response | void> {
    
    try {
        
        axios.put(`${BASE_URL}/dlc/update`,{
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

// DELETE DLC REGISTER
export async function deleteDLC(req: Request, res: Response):Promise<Response | void>{
    try {
        
        axios.delete(`${BASE_URL}/dlc/delete/${req.params.idDLC}`,{
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


// RESTORE DLC
export async function restoreDLC(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/dlc/restore/${req.params.idDLC}`,
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
