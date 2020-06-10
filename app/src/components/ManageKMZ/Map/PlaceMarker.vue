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
    return {
      placemark: {
        position: L.latLng(this.position.lat, this.position.lng),
        icon: L.divIcon({
          labelAnchor: [-3, 0],
          popupAnchor: [0, -23],
          html: `
          <div class="pin" style="background-color: ${this.color}"/>
            <svg>
              <use xlink:href='${placemarkIcon}#${this.icon}'></use>
            </svg>
          </div>`,
        }),
      },
    };
  },
};
</script>

<style lang="scss" scoped>
</style>
