import express, { Application } from 'express';
import morgan from 'morgan';
import cors from 'cors';

import IndexRoutes from './routes/index.routes';
import UserRoutes from './routes/user.routes';
import CheckRegionRoutes from './routes/check-region.routes';
import WishlistRoutes from './routes/wishlist.routes';
import LibraryRoutes from './routes/library.routes';
import DeveloperRoutes from './routes/developer.routes';
import CategoryRoutes from './routes/category.routes';
import RegionRoutes from './routes/region.routes';
import GameRoutes from './routes/game.routes';
import DLCRoutes from './routes/dlc.routes';
import DiscountsRoutes from './routes/discounts.routes';
import SalesRoutes from './routes/sales.routes';

export class App{

    private app: Application;

    constructor(private port?: number | string){
        this.app = express();
        this.settings();
        this.middlewares();
        this.routes();
    }

    settings(){
        this.app.set('port', this.port || process.env.PORT || 5000);       
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
        this.app.use('/user', UserRoutes);
        this.app.use('/check-region', CheckRegionRoutes);
        this.app.use('/wishlist', WishlistRoutes);
        this.app.use('/library', LibraryRoutes);
        this.app.use('/developers', DeveloperRoutes);
        this.app.use('/category', CategoryRoutes);
        this.app.use('/region', RegionRoutes);
        this.app.use('/games', GameRoutes);
        this.app.use('/dlc', DLCRoutes);
        this.app.use('/discount', DiscountsRoutes);
        this.app.use('/sales', SalesRoutes);
    }

    async listen(){
        await this.app.listen(this.app.get('port'));
        console.log(`Server on port ${this.app.get('port')} :D`);
    }

}