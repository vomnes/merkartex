import { Router, Request, Response } from "express";
import * as multer from 'multer';
import * as fs from 'fs';
import * as unzip from 'unzip';
import * as xml from 'xml-parse';

const router = Router();

const upload = multer({ dest: './public/data/uploads/' });

const readFile = async (file: fs.PathLike) => {
  return await new Promise((resolve, reject) => {
    let readStream = fs.createReadStream(file)
    let chunks = [];
    readStream.on('error', (error: Error) => {
      reject(error);
    });
    readStream.on('data', chunk => {
        chunks.push(chunk);
    });
    readStream.on('close', () => {
      resolve(Buffer.concat(chunks).toString());
    });
  })
}

const deleteFile = async (file: fs.PathLike) => {
  return await new Promise((_, reject) => {
    fs.unlink(file, (err) => {
      if (err) {
        reject('Failed to delete file : ' + err);
      };
    });
  });
}

const readKMZ = async (kmz_path: fs.PathLike) => {
  const UNZIP_PATH = './public/data/unzip/';
  return await new Promise<fs.PathLike>((resolve, reject) => {
    fs.createReadStream(kmz_path)
      .pipe(unzip.Parse())
      .on('error', (_: Error) => {
        reject('Invalid file');
      })
      .on('entry', (entry: any) => {
        var fileName = entry.path;
        var type = entry.type; // 'Directory' or 'File'
        var size = entry.size;
        entry.pipe(fs.createWriteStream(UNZIP_PATH+fileName));
        resolve(UNZIP_PATH+fileName);
      })
  });
}

const manageData = (xmlData: string) => {
  try {
    let jsonObject = xml.parse(xmlData);
    return jsonObject;
  } catch (error) {
    console.log('Failed to parse XML - ' + error);
    return
  }
}

router.post('/file', upload.single('map'), (req: Request, res: Response) => {
  if (!req.file) {
    res.send('Empty file\n');
    return
  }
  readKMZ(req.file.path)
    .then(async extractedFilePath => {
      readFile(extractedFilePath)
      .then((data: string) => {
        let jsonFormat = manageData(data);
        res.send(jsonFormat);
      })
        // .finally(() => deleteFile(extractedFilePath).catch((err) => console.log(err)));
    })
    .catch(error => {
      res.send(error);
    })
    .finally(() => deleteFile(req.file.path).catch((err) => console.log(err)));
});

// https://node.readthedocs.io/en/latest/api/zlib/

export { router as Upload };
