<template>
  <div class="load-map--background">
    <div class="load-map">
      <div class="load-map__head">
        <h1>Welcome on Merkartex</h1>
        <h2 class="text__details text--uppercase">Load your map 🚀</h2>
      </div>
      <transition name="fade-upload"></transition>
      <UploadedFile
      v-if="selectedFile.valid"
      :filename="selectedFile.name"
      :loaded="selectedFile.loaded"
      @clear="clearFile"/>
      <div class="load-map__upload-section" v-if="!selectedFile.valid">
        <h3 class="text__body">1. MapsMe</h3>
        <button class="primary-button--blue-white box-round-corner"
          @click="uploadMapsMeFile">
          <svg v-svg symbol="arrow"></svg>
          <p class="text__details text--uppercase">Upload MapsMe KMZ Map</p>
        </button>
        <input @change="fileSelected" ref="mapsmeFileInput"
        type="file" accept=".kmz">
      </div>
      <div class="load-map__upload-section" v-if="!selectedFile.valid">
        <h3 class="text__body">2. Google Maps</h3>
        <button class="primary-button--blue-white box-round-corner">
          <svg v-svg symbol="arrow"></svg>
          <p class="text__details text--uppercase">Upload Google Maps KMZ Map</p>
        </button>
        <button class="primary-button--blue-white box-round-corner" @click="gmapsPage">
          <svg v-svg symbol="link"></svg>
          <p class="text__details text--uppercase">Upload map from Google Maps Page</p>
        </button>
      </div>
      <div class="load-map__submit" v-if="selectedFile.valid">
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
import UploadedFile from './UploadedFile.vue';

export default {
  name: 'LoadMap',
  components: {
    UploadedFile,
  },
  data() {
    return {
      selectedFile: {
        name: '', loaded: 0, size: 0, valid: false,
      },
      error: null,
    };
  },
  methods: {
    uploadMapsMeFile() {
      this.$refs.mapsmeFileInput.click();
    },
    fileSelected(event) {
      const file = event.target.files[0];
      if (file && file.size > 100000) {
        this.setError('File size limited to 100MB');
        return;
      }
      if (file && !file.name.endsWith('.kmz')) {
        this.setError('Support only KMZ files type');
        return;
      }
      this.clearError();
      this.selectedFile.name = file.name;
      this.selectedFile.size = file.size;
      this.selectedFile.valid = true;
      const reader = new FileReader();
      // let uploadedFile = null;
      reader.onloadstart = this.readerEvent;
      reader.onload = this.readerEvent;
      reader.onprogress = this.readerEvent;
      // uploadedFile = reader.result;
      if (file) {
        reader.readAsText(file);
      }
    },
    readerEvent(e) {
      this.selectedFile.loaded = Math.round((e.loaded * 100) / e.total);
    },
    clearFile() {
      this.selectedFile = {
        name: '', loaded: 0, size: 0, valid: false,
      };
    },
    setError(error) {
      this.error = error;
    },
    clearError() {
      this.error = null;
    },
    gmapsPage() {
      if (this.error) {
        this.clearError();
      } else {
        this.setError('Not available');
      }
    },
  },
};
</script>

<style lang="scss" scoped>
  @import "LoadMap";
</style>
