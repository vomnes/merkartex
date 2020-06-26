<template>
  <div class="load-map--background">
    <div class="load-map">
      <div class="load-map__head">
        <h1>Welcome on Merkartex</h1>
        <h2 class="text__details text--uppercase">Load your map 🚀</h2>
      </div>
      <UploadedFile v-if="stateSelectedFile.valid"/>
      <div class="load-map__upload-section" v-if="!stateSelectedFile.valid">
        <h3 class="text__body">1. MapsMe</h3>
        <UploadFile
          content="Upload MapsMe KMZ Map"
          type="mapsme"
          @event="receiveEvent"/>
      </div>
      <div class="load-map__upload-section" v-if="!stateSelectedFile.valid">
        <h3 class="text__body">2. Google Maps</h3>
        <UploadFile
          content="Upload Google Maps KMZ Map"
          type="googlemaps"
          @event="receiveEvent"/>
        <div class="load-map__upload-section--btn-block">
          <button
            class="primary-button--blue-white box-round-corner"
            :class="{ 'load-map__upload-section--small-button': linkActivated }"
            @click="gmapsPage">
            <svg v-svg symbol="link"></svg>
            <p class="text__details text--uppercase">
              Upload map from Google - My Maps Page
            </p>
          </button>
          <input class="load-map__upload-section--link text__body"
          type="text" placeholder="Google - My Maps - URL Link">
        </div>
      </div>
      <div class="load-map__submit" v-if="stateSelectedFile.valid">
        <button
          class="primary-button--green box-round-corner">
          <p class="text__body">Load your map</p>
        </button>
      </div>
      <transition name="fade-error">
        <p class="load-map__error text__body" v-if="error">{{ error }}</p>
      </transition>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import UploadedFile from './UploadedFile.vue';
import UploadFile from './UploadFile.vue';

export default {
  name: 'LoadMap',
  components: {
    UploadedFile,
    UploadFile,
  },
  data() {
    return {
      error: null,
      linkActivated: false,
    };
  },
  methods: {
    gmapsPage() {
      if (!this.linkActivated) {
        this.linkActivated = true;
        // return;
      }
      // console.log('Check link Google Maps validity');
    },
    receiveEvent(type, data) {
      switch (type) {
        case 'setError':
          this.error = data.error;
          break;
        case 'clearError':
          this.error = null;
          break;
        default:
      }
    },
  },
  computed: {
    ...mapGetters('selectedFile', [
      'stateSelectedFile',
    ]),
  },
};
</script>

<style lang="scss">
  @import "LoadMap";
</style>
