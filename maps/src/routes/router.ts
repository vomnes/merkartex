import { Request, Response } from "express";

export class Routes {
  public routes(app): void {
    app.get('/', (_: Request, res: Response) => {
      res.send('Hello World!')
    })
  }
}
