<template>
  <div class="manage-kmz">
    <div class="map"></div>
    <div class="placemarks">
      <h1 class="text__title">Placemarks</h1>
      <div class="placemarks__list">
        <Placemark/>
        <Placemark/>
        <Placemark/>
      </div>
    </div>
  </div>
</template>

<script>
import Placemark from './Placemark/Placemark.vue';

export default {
  name: 'ManageKMZ',
  components: {
    Placemark,
  },
  props: {
    msg: String,
  },
};
</script>

<style scoped lang="scss">
  @import '../../assets/style/_main.scss';
  @import './Map.scss';
  @import './Placemarks.scss';

  .manage-kmz {
    display: flex;

    margin: 2rem 2rem 0 2rem;
  }
</style>
