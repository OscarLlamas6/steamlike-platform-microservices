
import request from "supertest";
import dotenv from 'dotenv';
import { App } from '../src/app';
dotenv.config();
const server = new App();


describe('Pruebas de Api Usuario',()=>{


    test('Testing Check-Region Microservice', async () => {
        await server.listen();
        const res = await request(server.app).get('/');
        server.close()
        // COMPARAR ESTADO 200
        expect(res.statusCode).toEqual(200);
    });

    afterAll(done => {        
        done();
    });
})


