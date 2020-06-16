<template>
  <article
    class="placemark"
    @click.stop="select"
    :class="[ placemarkIsSelected(this.index) ? 'placemark--selected' : 'placemark--default']">
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
            <li tabindex="0">
              <svg v-svg symbol="pencil"></svg>
              <p>Edit</p>
            </li>
            <li tabindex="0">
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
      <p class="text__details text--uppercase">{{ data.updatedAt }}</p>
    </div>
  </article>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';
import placemarksDesign from 'assets/data/placemarks-design.json';

const LIMIT_SIZE = 256;

export default {
  name: 'Placemark',
  props: {
    index: Number,
    data: Object,
  },
  computed: {
    ...mapGetters('placemarks', [
      'placemarkIsSelected',
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
        ? `${this.description.slice(0, LIMIT_SIZE)}...`
        : this.description;
    },
    getColor() {
      const color = this.data.iconStyle.replace('#placemark-', '');
      for (let i = 0; i < placemarksDesign.colors.length; i += 1) {
        if (placemarksDesign.colors[i].name === color) {
          return placemarksDesign.colors[i].color;
        }
      }
      return '';
    },
  },
  data() {
    return {
      description: this.data.description ? this.data.description : '',
      descriptionOpen: false,
      moreOptionsOpen: false,
    };
  },
  methods: {
    ...mapActions('placemarks', [
      'setPlacemarkSelectStatus',
      'unselectAllPlacemarks',
      'selectPlacemarksRange',
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
      } else {
        this.unselectAllPlacemarks(this.index);
      }
    },
  },
};
</script>

<style lang="scss" scoped>
  @import 'Placemark';
</style>
