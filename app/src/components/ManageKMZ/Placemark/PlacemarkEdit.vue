<template>
  <form class="placemark placemark-edit">
    <div class="header">
      <input class="text__title" type="text" name="title" value="Jardin Yuyuan">
      <div class="header--icon">
        <svg v-svg symbol="close"></svg>
      </div>
    </div>
    <div class="placemark-edit__row">
      <div class="placemark-edit__row--icon" data-title="Select color & category">
        <svg v-svg symbol="location"></svg>
        <p class="text__details">Sport</p>
      </div>
      <p class="text__details" data-title="Latitude, Longitude">• 62.208, 6.67296</p>
    </div>
    <p class="text__body">
      {{ descriptionContent }}
    </p>
    <div class="footer">
      <p class="text__details text--uppercase">Quartier</p>
      <datetime
        class="custom--datetime text--uppercase"
        type="datetime"
        v-model="datetime"
        use24-hour></datetime>
    </div>
    <div class="placemark-edit__actions">
      <button class="primary-button--blue text__details box-round-corner">Save changes</button>
    </div>
  </form>
</template>

<script>
import { Datetime } from 'vue-datetime';

const LIMIT_SIZE = 256;

export default {
  name: 'PlacemarkEdit',
  components: {
    Datetime,
  },
  data() {
    return {
      description: 'Le jardin Yuyuan est un jardin de deux hectares datant du XVIe siècle situé au centre de la Vieille Ville près de Chenghuangmiao à Shanghai, en Chine. Le jardin Yuyuan est un jardin de deux hectares datant du XVIe siècle situé au centre de la Vieille Ville près de Chenghuangmiao à Shanghai, en Chine.',
      datetime: '2019-06-21',
    };
  },
  computed: {
    descriptionContent() {
      return this.description.length > LIMIT_SIZE
        ? `${this.description.slice(0, LIMIT_SIZE)}...`
        : this.description;
    },
  },
};
</script>

<style lang="scss" scoped>
  @import "Placemark";
  @import "PlacemarkEdit";
</style>

<style lang="scss">
  @import '../../../assets/style/_main.scss';
  .custom--datetime .vdatetime-input {
    @extend .text__details;
    text-align: right;
    width: 20rem;
    cursor: pointer;
  }
</style>
