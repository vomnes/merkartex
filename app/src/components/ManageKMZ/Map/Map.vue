<template>
  <div class="map-area">
    <l-map
      ref="map"
      :zoom="zoom"
      :center="center"
      :options="mapOptions"
    >
      <PlaceMarkerCluster>
        <PlaceMarker
        v-for="placemark in getPlacemarks" :key="placemark.key"
        :position="{ lat: placemark.location.latitude, lng: placemark.location.longitude }"
        :title="placemark.name" :color="getColor(placemark.iconStyle)" icon="home"/>
      </PlaceMarkerCluster>
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
import { mapGetters } from 'vuex';

import placemarksDesign from 'assets/data/placemarks-design.json';

import PlaceMarkerCluster from './PlaceMarkerCluster.vue';
import PlaceMarker from './PlaceMarker.vue';


export default {
  name: 'Map',
  components: {
    LMap,
    LTileLayer,
    LControlZoom,
    PlaceMarker,
    PlaceMarkerCluster,
  },
  props: {
    geoCenter: Object,
  },
  data() {
    const placemarks = [
      {
        position: {
          lat: 31.22222,
          lng: 121.45806,
        },
        color: '#e31a23',
        icon: 'home',
      },
      {
        position: {
          lat: 31.21,
          lng: 121.42,
        },
        color: '#fb622a',
        icon: 'pencil',
      },
      {
        position: {
          lat: 31.229,
          lng: 121.451,
        },
        color: '#0171c4',
        icon: 'navigation',
      },
    ];
    return {
      zoom: 4,
      center: L.latLng(this.geoCenter.latitude, this.geoCenter.longitude),
      url: 'https://api.maptiler.com/maps/voyager/{z}/{x}/{y}.png?key=1Yg83v5zhwYytD6ZRJrP',
      attribution: '<a href="https://carto.com/" target="_blank">&copy; CARTO</a> <a href="https://www.openstreetmap.org/copyright" target="_blank">&copy; OSM</a>',
      mapOptions: { zoomControl: false, attributionControl: true, zoomSnap: true },
      placemarks,
      bounds: {
        northEast: {
          lat: 0,
          lng: 0,
        },
        southWest: {
          lat: 0,
          lng: 0,
        },
      },
    };
  },
  methods: {
    getColor(iconColor) {
      const color = iconColor.replace('#placemark-', '');
      for (let i = 0; i < placemarksDesign.colors.length; i += 1) {
        if (placemarksDesign.colors[i].name === color) {
          return placemarksDesign.colors[i].color;
        }
      }
      return '';
    },
  },
  computed: {
    ...mapGetters('placemarks', [
      'getPlacemarks',
    ]),
  },
};
</script>

<style lang="scss" scoped>
  @import '@/assets/style/_main.scss';
  @import "Map";
</style>

<style lang="scss">
  @import "PlaceMarker";
</style>
