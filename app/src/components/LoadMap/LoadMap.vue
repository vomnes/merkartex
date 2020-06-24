<template>
  <div class="load-map--background">
    <div class="load-map">
      <div class="load-map__head">
        <h1>Welcome on Merkartex</h1>
        <h2 class="text__details text--uppercase">Load your map 🚀</h2>
      </div>
      <div class="load-map__upload-section">
        <h3 class="text__body">1. MapsMe</h3>
        <button class="primary-button--blue-white box-round-corner">
          <svg v-svg symbol="arrow"></svg>
          <p class="text__details text--uppercase">Upload MapsMe KMZ Map</p>
        </button>
      </div>
      <div class="load-map__upload-section">
        <h3 class="text__body">2. Google Maps</h3>
        <button class="primary-button--blue-white box-round-corner">
          <svg v-svg symbol="arrow"></svg>
          <p class="text__details text--uppercase">Upload Google Maps KMZ Map</p>
        </button>
        <button class="primary-button--blue-white box-round-corner">
          <svg v-svg symbol="link"></svg>
          <p class="text__details text--uppercase">Upload map from Google Maps Page</p>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: '',
  data() {
    return {

    };
  },
};
</script>

<style lang="scss" scoped>
  @import "LoadMap";
</style>
