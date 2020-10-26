<template>
  <md-dialog :md-active="openEdit" @md-clicked-outside="$emit('manageOpen', false)">
    <form class="placemark placemark-edit" @submit.prevent="saveChanges">
      <div class="header">
        <input class="text__title" type="text" name="title" v-model="editData.title">
        <div class="header--icon" @click="$emit('manageOpen', false)">
          <svg v-svg symbol="close"></svg>
        </div>
      </div>
      <div class="placemark-edit__row">
        <div
          class="placemark-edit__row--icon"
          data-title="Select color & category"
          data-title-position="top"
          @click="manageOpenEditIcon(true)">
          <svg v-svg :color="editData.placemark.color.color" symbol="location"></svg>
          <p class="text__details">{{ editData.placemark.category.name }}</p>
        </div>
        <ModaleEditIcon
          :open="open.editIcon" @manageOpen="manageOpenEditIcon"
          :placemark="editData.placemark"
          @sendEditIconValue="getEditIconValue"/>
        <div data-title="Latitude, Longitude" data-title-position="top">
          <input
            type="text"
            :class="{ 'text--warning': isInvalidLocation }"
            class="text__details text--center"
            v-model="editData.location">
        </div>
      </div>
      <p class="text__body" @click="manageOpenEditDescription(true)">
        {{ descriptionContent }}
      </p>
      <ModaleEditDescription
      :open="open.editDescription"
      :description="editData.description"
      @manageOpen="manageOpenEditDescription"
      @sendDescriptionValue="getDescriptionValue"/>
      <div class="footer">
        <p class="text__details text--uppercase">{{ editData.featureType }}</p>
        <datetime
          class="custom--datetime text--uppercase"
          type="datetime"
          v-model="editData.datetime"
          use24-hour></datetime>
      </div>
      <div class="placemark-edit__actions">
        <button type="submit"
          :class="[
            isValidForm ? 'primary-button--blue' : 'button-passive button-disabled',
            'text__details box-round-corner',
            ]"
            :disabled="!isValidForm">
          Save changes
        </button>
      </div>
    </form>
  </md-dialog>
</template>

<script>
import { mapActions } from 'vuex';
import { Datetime } from 'vue-datetime';
import placemarksDesign from 'assets/data/placemarks-design.json';
import ModaleEditDescription from './Modale/ModaleEditDescription.vue';
import ModaleEditIcon from './Modale/ModaleEditIcon.vue';

const cloneDeep = require('clone-deep');

const LIMIT_SIZE = 256;

export default {
  name: 'PlacemarkEdit',
  props: {
    openEdit: Boolean,
    data: Object,
    type: String,
  },
  components: {
    Datetime,
    ModaleEditIcon,
    ModaleEditDescription,
  },
  data() {
    const editData = this.initData();
    return {
      editData,
      open: {
        editIcon: false,
        editDescription: false,
      },
    };
  },
  computed: {
    descriptionContent() {
      if (!this.editData.description) {
        return 'No description';
      }
      return this.editData.description.length > LIMIT_SIZE
        ? `${this.editData.description.slice(0, LIMIT_SIZE)}...`
        : this.editData.description;
    },
    isInvalidLocation() {
      const raw = this.editData.location;
      const array = raw.split(',');
      const validFloat = (str) => /^[ ]*?[+-]?\d+(\.\d+)?$/.test(str);
      if (array.length !== 2) {
        return true;
      }
      if (!validFloat(array[0]) || !validFloat(array[1])) {
        return true;
      }
      if (Number.isNaN(parseFloat(array[0])) || Number.isNaN(parseFloat(array[1]))) {
        return true;
      }
      return false;
    },
    isValidForm() {
      return !this.isInvalidLocation;
    },
  },
  watch: {
    openEdit() {
      this.editData = this.initData();
    },
  },
  methods: {
    ...mapActions('placemarks', [
      'updatePlacemark',
      'pushPlacemark',
    ]),
    initData() {
      let datetime = this.data.updatedAt;
      if (this.type !== 'edit') {
        datetime = new Date().toISOString();
      }
      return {
        title: this.data.name,
        description: this.data.description ? this.data.description : null,
        datetime,
        location: `${this.data.location.latitude},${this.data.location.longitude}`,
        placemark: {
          color: this.getFormatedColor(this.data.icon.style),
          category: {
            name: this.data.icon.category ? this.data.icon.category : 'None',
          },
        },
        featureType: this.data.featureType,
      };
    },
    manageOpenEditIcon(value) {
      this.open.editIcon = value;
    },
    manageOpenEditDescription(value) {
      this.open.editDescription = value;
    },
    getEditIconValue(value) {
      this.editData.placemark = value;
    },
    getDescriptionValue(value) {
      this.editData.description = value;
    },
    getFormatedColor(rawColor) {
      for (let i = 0; i < placemarksDesign.colors.length; i += 1) {
        if (placemarksDesign.colors[i].name === rawColor) {
          return placemarksDesign.colors[i];
        }
      }
      return placemarksDesign.colors[0];
    },
    saveChanges() {
      if (!this.isValidForm) {
        return;
      }
      this.$emit('manageOpen', false);
      const newData = cloneDeep(this.data);
      newData.name = this.editData.title;
      newData.description = this.editData.description;
      newData.updatedAt = this.editData.datetime;
      newData.icon = {
        style: this.editData.placemark.color.name,
        category: this.editData.placemark.category.name,
      };
      const location = this.editData.location.split(',');
      newData.location = {
        latitude: parseFloat(location[0]),
        longitude: parseFloat(location[1]),
      };
      switch (this.type) {
        case 'edit':
          this.updatePlacemark({ id: this.data.id, data: newData });
          break;
        case 'new':
          this.pushPlacemark(newData);
          break;
        default:
      }
    },
  },
};
</script>

<style lang="scss" scoped>
  @import "Placemark";
  @import "PlacemarkEdit";
</style>

<style lang="scss">
  @import '@/assets/style/_main.scss';
  .custom--datetime .vdatetime-input {
    @extend .text__details;
    text-align: right;
    width: 20rem;
    cursor: pointer;
  }
</style>
