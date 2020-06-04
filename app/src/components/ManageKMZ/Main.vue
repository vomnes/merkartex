<template>
  <div class="manage-kmz">
    <div class="header">
      <div class="header--title">
        <svg v-svg symbol="kmz"></svg>
        <h1>Shanghai.kmz</h1>
      </div>
    </div>
    <div class="map">
      <!-- https://cloud.maptiler.com/maps/voyager -->
    </div>
    <div class="placemarks">
      <h1 class="text__title">Placemarks</h1>
      <div class="placemarks__list custom-scrollbar">
        <Placemark/>
        <Placemark/>
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
    display: grid;
    grid-template-columns: 3fr 60rem;
    grid-template-rows: 6.5rem auto;

    margin: 1rem 2rem 0 2rem;

    & .header {
      grid-column-start: 1;
      grid-column-end: 2;

      &--title {
        display: flex;
        align-items: center;

        & h1 {
          font-family: $font-3;
          font-weight: 400;
          font-size: 2.4rem;
        }

        & svg {
          width: 3rem;
          height: 4.2rem;
          margin-right: 1.5rem;
        }
      }
    }

    & .map {
      grid-column-start: 2;
      grid-column-end: 1;
    }
  }
</style>
