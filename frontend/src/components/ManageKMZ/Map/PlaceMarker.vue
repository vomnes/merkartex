<template>
  <l-marker
    :lat-lng="placemark.position"
    :icon="placemark.icon"
    @click="scrollToPlacemark"
    @mouseenter="hoveredEnter"
    @mouseleave="hoveredLeave"
  />
</template>

<script>
import L from 'leaflet';
import { LMarker } from 'vue2-leaflet';
import placemarkIcon from 'assets/icons/_sprite.svg';

export default {
  name: 'PlaceMarker',
  components: {
    LMarker,
  },
  props: {
    position: Object,
    title: String,
    color: {
      type: String,
      default: '#3A7BD5',
    },
    icon: {
      type: String,
    },
    index: Number,
  },
  data() {
    const content = `
      <div class="pin__content">
        <h3>${this.title}</h3>
      </div>
    `;
    // <div class="pin__content--actions">
    //   <svg>
    //     <use xlink:href='${placemarkIcon}#pencil'></use>
    //   </svg>
    //   <svg>
    //     <use xlink:href='${placemarkIcon}#trash'></use>
    //   </svg>
    //   <svg>
    //     <use xlink:href='${placemarkIcon}#duplicate'></use>
    //   </svg>
    // </div>
    const html = `
    <div class="pin" id="pin-1" style="color: ${this.color}"/>
      <svg>
        <use xlink:href='${placemarkIcon}#${this.icon}'></use>
      </svg>
      ${content}
    </div>`;
    return {
      placemark: {
        position: L.latLng(this.position.lat, this.position.lng),
        icon: L.divIcon({
          labelAnchor: [-3, 0],
          popupAnchor: [0, -23],
          html,
        }),
      },
    };
  },
  methods: {
    scrollToPlacemark() {
      const elem = document
        .getElementById(`placemark-${this.position.lat}-${this.position.lng}`);
      if (elem) {
        elem
          .scrollIntoView({
            behavior: 'smooth',
            block: 'start',
          });
      }
    },
    hoveredEnter() {
      const elem = document
        .getElementById(`placemark-${this.position.lat}-${this.position.lng}`);
      if (elem) {
        elem
          .classList
          .add('placemark--hovered');
      }
    },
    hoveredLeave() {
      const elem = document
        .getElementById(`placemark-${this.position.lat}-${this.position.lng}`);
      if (elem) {
        elem
          .classList
          .remove('placemark--hovered');
      }
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
