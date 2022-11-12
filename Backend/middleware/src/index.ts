import { App } from './app';
import dotenv from 'dotenv';


async function main(){
    dotenv.config();
    let port = process.env.MIDDLEWARE_PORT || 5000;
    const app = new App(port);
    await app.listen();
}

main();



