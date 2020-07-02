/* eslint "import/prefer-default-export": "off" */
import importData from './importData';
import uploadFile from './uploadFile';
import uploadFromMyMaps from './uploadFromMyMaps';

import config from './config';

const apiConf = {
  BACK_URL: config.apiDomain,
};

export default {
  importData: () => importData(),
  uploadFile: (data) => uploadFile(apiConf, data),
  uploadFromMyMaps: (myMapsURL) => uploadFromMyMaps(apiConf, myMapsURL),
};
