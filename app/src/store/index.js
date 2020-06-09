import Vue from 'vue';
import Vuex from 'vuex';

import placemarks from './modules/placemarks';
import keypress from './modules/keypress';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  modules: {
    placemarks,
    keypress,
  },
  strict: debug,
});
