<template>
  <div class="map-area">
    <l-map
      ref="map"
      :zoom="zoom"
      :center="center"
      :options="mapOptions"
    >
      <PlaceMarker :position="placemark.position" color="#e31a23" icon="home"/>
      <l-control-zoom :position="'bottomleft'"></l-control-zoom>
      <l-tile-layer :url="url" :attribution="attribution"></l-tile-layer>
    </l-map>
  </div>
</template>

<script>
import L from 'leaflet';
import {
  LMap,
  LTileLayer,
  LControlZoom,
} from 'vue2-leaflet';
import PlaceMarker from './PlaceMarker.vue';

export default {
  name: 'Map',
  components: {
    LMap, LTileLayer, LControlZoom, PlaceMarker,
  },
  data() {
    return {
      zoom: 10,
      center: L.latLng(31.22222, 121.45806),
      url: 'https://api.maptiler.com/maps/voyager/{z}/{x}/{y}.png?key=1Yg83v5zhwYytD6ZRJrP',
      attribution: '<a href="https://carto.com/" target="_blank">&copy; CARTO</a> <a href="https://www.openstreetmap.org/copyright" target="_blank">&copy; OSM</a>',
      mapOptions: { zoomControl: false, attributionControl: true, zoomSnap: true },
      placemark: {
        position: {
          lat: 31.22222,
          lng: 121.45806,
        },
      },
    };
  },
};
</script>

<style lang="scss" scoped>
  @import '../../../assets/style/_main.scss';
  @import "Map";
</style>

<style lang="scss">
  @import "PlaceMarker";
</style>
