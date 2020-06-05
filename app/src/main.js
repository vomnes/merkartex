import Vue from 'vue';
import SvgSprite from 'vue-svg-sprite';
import { Datetime } from 'vue-datetime';

import 'leaflet/dist/leaflet.css';
import 'vue-datetime/dist/vue-datetime.css';

import App from './App.vue';
import store from './store';

import './assets/style/_main.scss';
import SvgSpriteFile from './assets/icons/_sprite.svg';

Vue.config.productionTip = false;

Vue.use(SvgSprite, {
  url: SvgSpriteFile,
});

Vue.use(Datetime);

new Vue({
  store,
  render: (h) => h(App),
}).$mount('#app');
