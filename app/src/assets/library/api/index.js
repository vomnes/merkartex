/* eslint "import/prefer-default-export": "off" */
import importData from './importData';
import uploadFile from './uploadFile';

import config from './config';

const apiConf = {
  BACK_URL: config.apiDomain,
};

export default {
  importData: () => importData(),
  uploadFile: (params) => uploadFile(params, apiConf),
};
