<template>
  <div class="map-area">
    <l-map
      ref="map"
      :zoom="zoom"
      :center="center"
      :options="mapOptions"
    >
      <PlaceMarkerCluster>
        <template v-if="render">
          <PlaceMarker
          v-for="placemark in getPlacemarks"
          :key="placemark.id" :index="placemark.id"
          :position="{ lat: placemark.location.latitude, lng: placemark.location.longitude }"
          :title="placemark.name"
          :color="getColor(placemark.icon.style)"
          :icon="placemark.icon.category ? placemark.icon.category.toLowerCase() : ''"/>
        </template>
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
    return {
      zoom: 4,
      center: L.latLng(this.geoCenter.latitude, this.geoCenter.longitude),
      url: 'https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png',
      attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors &copy; <a href="https://carto.com/attributions">CARTO</a>',
      mapOptions: { zoomControl: false, attributionControl: true, zoomSnap: true },
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
      render: true,
    };
  },
  mounted() {
    this.$nextTick(() => {
      this.map = this.$refs.map.mapObject;
    });
    this.$root.$on('flyToLocation', this.flyTo);
  },
  watch: {
    getPlacemarks() {
      // Force rerender on placemarker change
      this.render = false;
      this.$nextTick(() => {
        this.render = true;
      });
    },
  },
  methods: {
    getColor(iconColor) {
      for (let i = 0; i < placemarksDesign.colors.length; i += 1) {
        if (placemarksDesign.colors[i].name === iconColor) {
          return placemarksDesign.colors[i].color;
        }
      }
      return '';
    },
    flyTo({ lat, lng }) {
      const defaultZoom = 16;
      const newCenter = [lat, lng];
      this.map.flyTo(newCenter, defaultZoom);
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

<!-- const placemarks = [
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
]; -->
