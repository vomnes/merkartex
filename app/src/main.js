import Vue from 'vue';
import SvgSprite from 'vue-svg-sprite';

import App from './App.vue';
import store from './store';

import './assets/style/_main.scss';
import SvgSpriteFile from './assets/icons/_sprite.svg';

Vue.config.productionTip = false;

Vue.use(SvgSprite, {
  url: SvgSpriteFile,
});

new Vue({
  store,
  render: (h) => h(App),
}).$mount('#app');
