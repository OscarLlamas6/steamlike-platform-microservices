import { Request, Response } from 'express'
import axios from 'axios';
import dotenv from 'dotenv';
dotenv.config();

const { API_REGION_BASE_URL } = process.env;

const BASE_URL = API_REGION_BASE_URL ? API_REGION_BASE_URL : "http://localhost:3002";

//CHECK REGION
export async function checkRegion(req: Request, res: Response): Promise<Response | void> {
    try {
            
        axios.get(`${BASE_URL}/region/`,
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

