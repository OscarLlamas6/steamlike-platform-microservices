import { App } from './app';
import dotenv from 'dotenv';


async function main(){
    dotenv.config();
    let port = process.env.PORT || 3002;
    const app = new App(port);
    await app.listen();
}

main();



