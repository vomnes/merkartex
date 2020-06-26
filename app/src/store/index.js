import Vue from 'vue';
import Vuex from 'vuex';

import placemarks from './modules/placemarks';
import keypress from './modules/keypress';
import selectedFile from './modules/selectedFile';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  modules: {
    placemarks,
    keypress,
    selectedFile,
  },
  strict: debug,
});
