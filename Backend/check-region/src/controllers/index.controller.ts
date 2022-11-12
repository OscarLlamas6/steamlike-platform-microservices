import { Request, Response } from 'express'

export function indexWelcome(req: Request, res: Response): Response {
   return res.json('Region Microservice - Steamlike Platform | Grupo 4 :D'); 
}