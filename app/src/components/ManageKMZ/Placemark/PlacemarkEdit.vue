<template>
  <form class="placemark placemark-edit">
    <div class="header">
      <input class="text__title" type="text" name="title" value="Jardin Yuyuan">
      <div class="header--icon">
        <svg v-svg symbol="close"></svg>
      </div>
    </div>
    <div class="placemark-edit__row">
      <div
        class="placemark-edit__row--icon"
        data-title="Select color & category"
        @click="manageOpenEditIcon(true)">
        <svg v-svg color="green" symbol="location"></svg>
        <p class="text__details">Sport</p>
      </div>
      <ModaleEditIcon :open="open.editIcon" @manageOpen="manageOpenEditIcon"/>
      <div data-title="Latitude, Longitude">
        <input
          type="text"
          class="text__details text--center"
          value="62.208, 6.67296">
      </div>
    </div>
    <p class="text__body" @click="manageOpenEditDescription(true)">
      {{ descriptionContent }}
    </p>
    <ModaleEditDescription :open="open.editDescription" @manageOpen="manageOpenEditDescription"/>
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
import ModaleEditDescription from './Modale/ModaleEditDescription.vue';
import ModaleEditIcon from './Modale/ModaleEditIcon.vue';

const LIMIT_SIZE = 256;

export default {
  name: 'PlacemarkEdit',
  components: {
    Datetime,
    ModaleEditIcon,
    ModaleEditDescription,
  },
  data() {
    return {
      description: 'Le jardin Yuyuan est un jardin de deux hectares datant du XVIe siècle situé au centre de la Vieille Ville près de Chenghuangmiao à Shanghai, en Chine. Le jardin Yuyuan est un jardin de deux hectares datant du XVIe siècle situé au centre de la Vieille Ville près de Chenghuangmiao à Shanghai, en Chine.',
      datetime: '2019-06-21',
      open: {
        editIcon: false,
        editDescription: false,
      },
    };
  },
  computed: {
    descriptionContent() {
      return this.description.length > LIMIT_SIZE
        ? `${this.description.slice(0, LIMIT_SIZE)}...`
        : this.description;
    },
  },
  methods: {
    manageOpenEditIcon(value) {
      this.open.editIcon = value;
    },
    manageOpenEditDescription(value) {
      this.open.editDescription = value;
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
