import * as express from 'express';
import * as bodyParser from 'body-parser';

import { Routes } from './routes/router';

import { Upload } from './routes/upload';

class App {
  public app: express.Application;
  public router: Routes = new Routes();


  constructor() {
    this.app = express();
    this.config();
    this.allRoutes();
    this.router.routes(this.app);
  }

  private config(): void {
    this.app.use(bodyParser.json());
    this.app.use(bodyParser.urlencoded({ extended: true }));
  }

  private allRoutes(): void {
    this.app.use("/upload", Upload);
  }
}

export default new App().app;
