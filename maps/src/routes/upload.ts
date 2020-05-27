import { Router, Request, Response } from "express";
import * as multer from 'multer';

const router = Router();

const upload = multer({ dest: './public/data/uploads/' });

router.post('/file', upload.single('map'), (req: Request, res: Response) => {
  console.log('->', req);
  console.log('=>', res);
  res.send("_Upload_");
})

export { router as Upload };
