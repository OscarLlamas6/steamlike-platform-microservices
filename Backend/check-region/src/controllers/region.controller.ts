import { Request, Response } from 'express'
import axios from 'axios';
import geoip from 'geoip-lite'

import dotenv from 'dotenv';
dotenv.config();

export async function checkRegion(req: Request, res: Response): Promise<Response | void> {
    try {

    const { ipAddress } = req.query;

    if (!ipAddress) {
        res.json({
            "region":"GU"
        })
    }

    let response = await axios.get(`http://ip-api.com/json/${ipAddress}`, {
        headers: {
            'Content-Type': 'application/json'
            },
    });

    if (response.status === 200) {
        res.json(response.data)
    } 

    var geo = geoip.lookup(`${ipAddress}`);

    if (!geo) {
        res.json({
            "region":"GU"
        })
    }

    res.json(geo);
                 
    } catch (e: any) {
        console.log(`${e.message} - ${e.stack}`);
        return res.status(404).json({
            message: `${e.message} - ${e.stack}`,
            status: 404
        });
    }
}

