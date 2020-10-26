<template>
  <article
    class="placemark"
    :id="`placemark-${data.location.latitude}-${data.location.longitude}`"
    @click.stop="select"
    :class="[ placemarkIsSelected(index)
    ? 'placemark--selected'
    : 'placemark--default']">
    <div class="placemark--side" :style="{ backgroundColor: getColor }"></div>
    <div class="header">
      <h2 class="text__title">{{ data.name }}</h2>
      <div class="header--icon">
        <svg v-svg symbol="ellipsis"
        tabindex="0"
        v-on:click.stop.prevent="toggleMoreOptionsOpen"
        v-on:keyup.enter="toggleMoreOptionsOpen"></svg>
        <div
          v-if="moreOptionsOpen"
          class="placemark--more-options">
          <ul>
            <li tabindex="0" @click.stop="editMode">
              <svg v-svg symbol="pencil"></svg>
              <p>Edit</p>
            </li>
            <li tabindex="0" @click.stop="removePlacemark(index); toggleMoreOptionsOpen()">
              <svg v-svg symbol="trash"></svg>
              <p>Remove</p>
            </li>
            <li tabindex="0">
              <svg v-svg symbol="duplicate"></svg>
              <p>Duplicate</p>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <p class="text__body">
      {{ descriptionContent }}
      <span
        v-if="hasSeeMoreDescription"
        @click.stop="toggleDescriptionOpen"
        @keyup.enter="toggleDescriptionOpen"
        tabindex="0">
        {{ !descriptionOpen ? 'See more' : 'Close' }}
      </span>
    </p>
    <div class="footer">
      <p class="text__details text--uppercase">{{ data.featureType }}</p>
      <p class="text__details text--uppercase">{{ moment(data.updatedAt).format('LLL') }}</p>
    </div>
    <PlacemarkEdit
      :data="data"
      :openEdit="open.edit"
      type="edit"
      @manageOpen="manageToggleEditMode"/>
  </article>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

import placemarksDesign from 'assets/data/placemarks-design.json';

import PlacemarkEdit from './PlacemarkEdit.vue';

const moment = require('moment');

const LIMIT_SIZE = 256;

export default {
  name: 'Placemark',
  props: {
    index: Number,
    data: Object,
  },
  components: {
    PlacemarkEdit,
  },
  data() {
    return {
      description: this.data.description ? this.data.description : '',
      descriptionOpen: false,
      moreOptionsOpen: false,
      open: {
        edit: false,
      },
      moment,
    };
  },
  computed: {
    ...mapGetters('placemarks', [
      'placemarkIsSelected',
      'hasPlacemarksSelection',
    ]),
    ...mapGetters('keypress', [
      'modeSelectRangeOn',
      'modeSelectMultiOn',
      'modeSelectOn',
    ]),
    hasSeeMoreDescription() {
      return this.description.length > LIMIT_SIZE;
    },
    descriptionContent() {
      return this.hasSeeMoreDescription && !this.descriptionOpen
        ? `${this.data.description.slice(0, LIMIT_SIZE)}...`
        : this.data.description;
    },
    getColor() {
      const color = this.data.icon.style;
      for (let i = 0; i < placemarksDesign.colors.length; i += 1) {
        if (placemarksDesign.colors[i].name === color) {
          return placemarksDesign.colors[i].color;
        }
      }
      return '';
    },
  },
  methods: {
    ...mapActions('placemarks', [
      'setPlacemarkSelectStatus',
      'unselectAllPlacemarks',
      'selectPlacemarksRange',
      'removePlacemark',
    ]),
    toggleDescriptionOpen() {
      this.descriptionOpen = !this.descriptionOpen;
    },
    toggleMoreOptionsOpen() {
      this.moreOptionsOpen = !this.moreOptionsOpen;
    },
    select() {
      if (this.modeSelectOn) {
        if (this.modeSelectMultiOn) {
          this.setPlacemarkSelectStatus({
            index: this.index,
            status: !this.placemarkIsSelected(this.index),
          });
        } else if (this.modeSelectRangeOn) {
          this.selectPlacemarksRange(this.index);
        }
      } else if (this.hasPlacemarksSelection) {
        this.unselectAllPlacemarks(this.index);
      } else {
        this.$root.$emit('flyToLocation', {
          lat: this.data.location.latitude,
          lng: this.data.location.longitude,
        });
      }
    },
    editMode() {
      this.open.edit = true;
      this.toggleMoreOptionsOpen();
    },
    manageToggleEditMode(value) {
      this.open.edit = value;
    },
  },
};
</script>

<style lang="scss" scoped>
  @import 'Placemark';
</style>
