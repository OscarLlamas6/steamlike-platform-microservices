import express, { Application } from 'express';
import morgan from 'morgan';
import cors from 'cors';

import IndexRoutes from './routes/index.routes';
import RegionRoutes from './routes/region.routes';


export class App{

    public app: Application;
    public myServer: any;

    constructor(private port?: number | string){
        this.app = express();
        this.settings();
        this.middlewares();
        this.routes();
    }

    settings(){
        this.app.set('port', this.port || process.env.PORT || 3002);       
    }

    middlewares(){
        let corsOptions = { origin: true, optionsSuccessStatus: 200 };
        this.app.use(morgan('dev'));
        this.app.use(cors(corsOptions));
        this.app.use(express.json({limit: '50mb'}));
        this.app.use(express.json());  
    }
    
    routes(){
        this.app.use(IndexRoutes);
        this.app.use('/region', RegionRoutes);
    }

    async listen(){
        this.myServer = await this.app.listen(this.app.get('port'));
        console.log(`Server on port ${this.app.get('port')} :D`);
    }

    close() {
        this.myServer.close()
    }


}