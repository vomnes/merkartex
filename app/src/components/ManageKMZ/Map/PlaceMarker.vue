<template>
  <l-marker :lat-lng="placemark.position" :icon="placemark.icon"/>
</template>

<script>
import L from 'leaflet';
import { LMarker } from 'vue2-leaflet';
import placemarkIcon from '../../../assets/icons/_sprite.svg';

export default {
  name: 'PlaceMarker',
  components: {
    LMarker,
  },
  props: {
    position: Object,
    color: {
      type: String,
      default: '#3A7BD5',
    },
    icon: {
      type: String,
    },
  },
  data() {
    const content = `
      <div class="pin__content">
        <h3>Jardin Yuyuan</h3>
        <div class="pin__content--actions">
          <svg>
            <use xlink:href='${placemarkIcon}#pencil'></use>
          </svg>
          <svg>
            <use xlink:href='${placemarkIcon}#trash'></use>
          </svg>
          <svg>
            <use xlink:href='${placemarkIcon}#duplicate'></use>
          </svg>
        </div>
      </div>
    `;
    const html = `
    <div class="pin" style="background-color: ${this.color}"/>
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
};
</script>

<style lang="scss" scoped>
</style>
