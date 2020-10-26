import Vue from 'vue';
import VueRouter from 'vue-router';
import SvgSprite from 'vue-svg-sprite';
import { Datetime } from 'vue-datetime';
import {
  MdDialog,
  MdAutocomplete,
  MdHighlightText,
  MdField,
  MdMenu,
  MdButton,
  MdList,
} from 'vue-material/dist/components';
import Notifications from 'vue-notification';

import 'leaflet/dist/leaflet.css';
import 'vue-datetime/dist/vue-datetime.css';
import 'vue-material/dist/vue-material.min.css';
import 'vue-material/dist/theme/default.css';

import App from './App.vue';
import store from './store';

import './assets/style/_main.scss';
import SvgSpriteFile from './assets/icons/_sprite.svg';

import LoadMap from './components/LoadMap/LoadMap.vue';
import ManageKMZ from './components/ManageKMZ/Main.vue';

Vue.config.productionTip = false;

Vue.use(SvgSprite, {
  url: SvgSpriteFile,
});

Vue.use(Datetime);

[
  MdDialog,
  MdAutocomplete,
  MdHighlightText,
  MdField,
  MdMenu,
  MdButton,
  MdList,
].forEach((x) => Vue.use(x));

Vue.use(Notifications);

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: LoadMap,
    },
    {
      path: '/map',
      component: ManageKMZ,
    },
  ],
});

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
