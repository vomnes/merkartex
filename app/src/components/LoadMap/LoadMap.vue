<template>
  <div class="load-map--background">
    <div class="load-map">
      <div class="load-map__head">
        <h1>Welcome on Merkartex</h1>
        <h2 class="text__details text--uppercase">Load your map 🚀</h2>
      </div>
      <UploadedFile v-if="stateSelectedFile.valid" @event="receiveEvent"/>
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
          :class="{ 'input--warning': gmapsCheck }"
          v-model="gmapsValue"
          type="text" placeholder="Google - My Maps - URL Link">
        </div>
      </div>
      <div class="load-map__submit" v-if="stateSelectedFile.valid">
        <button
          class="primary-button--green box-round-corner"
          @click="loadMap">
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

import Api from 'assets/library/api/index';

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
      gmapsValue: '',
    };
  },
  methods: {
    gmapsPage() {
      if (!this.linkActivated) {
        this.linkActivated = true;
        return;
      }
      if (this.gmapsValue === '') {
        this.error = 'Google My Maps can\'t be empty';
        return;
      }
      if (this.gmapsCheck) {
        this.error = 'Google My Maps URL is invalid';
        return;
      }
      this.error = null;
      Api.uploadFromMyMaps(this.gmapsValue)
        .then((res) => this.manageHTTPResponse(res));
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
    loadMap() {
      Api.uploadFile({
        map: this.getSelectedFileContent,
        filename: this.stateSelectedFile.name,
      })
        .then((res) => this.manageHTTPResponse(res));
    },
    manageHTTPResponse(res) {
      if (res.status >= 500) {
        res.json()
          .then((data) => {
            this.error = data.error;
          });
      } else if (res.status >= 400) {
        res.json()
          .then((data) => {
            this.error = data.error;
          });
      } else {
        res.json()
          .then((data) => {
            this.error = null;
            localStorage.setItem('map_data', JSON.stringify(data));
            this.$router.push('/map');
          });
      }
    },
  },
  computed: {
    ...mapGetters('selectedFile', [
      'stateSelectedFile',
      'getSelectedFileContent',
    ]),
    gmapsCheck() {
      if (this.gmapsValue === '') {
        return false;
      }
      return this.gmapsValue.search(/google\.com\/maps\/d\/(edit|viewer)\?.*mid=/g) === -1;
    },
  },
};
</script>

<style lang="scss">
  @import "LoadMap";

  .input--warning {
    border: .1rem solid $color-warning;
  }
</style>
