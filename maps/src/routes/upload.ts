import { Router, Request, Response } from "express";
import * as multer from 'multer';
import * as fs from 'fs';

const router = Router();

const upload = multer({ dest: './public/data/uploads/' });

router.post('/file', upload.single('map'), (req: Request, res: Response) => {
  if (!req.file) {
    res.send('Empty file\n');
    return
  }
  const data = fs.readFileSync(req.file.path, 'utf8');
  res.send(data);
});

export { router as Upload };
