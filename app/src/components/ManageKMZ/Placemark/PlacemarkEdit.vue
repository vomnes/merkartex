<template>
  <form class="placemark placemark-edit">
    <div class="header">
      <input class="text__title" type="text" name="title" value="Jardin Yuyuan">
      <div class="header--icon">
        <!-- <svg v-svg symbol="ellipsis"></svg> -->
      </div>
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
  }
</style>
