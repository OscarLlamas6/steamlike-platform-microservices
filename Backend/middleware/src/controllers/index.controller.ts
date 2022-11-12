import { Request, Response } from 'express'

export function indexWelcome(req: Request, res: Response): Response {
   return res.json('Middleware - Steamlike Platform | Grupo 4 :D'); 
}