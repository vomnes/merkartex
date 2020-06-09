<template>
  <div class="manage-kmz">
    <div class="header">
      <div class="header--title">
        <svg v-svg symbol="kmz"></svg>
        <h1>Shanghai.kmz</h1>
      </div>
    </div>
    <Map/>
    <div class="placemarks">
      <div class="placemarks__header">
        <h1 class="text__title">Placemarks</h1>
        <div class="placemarks__header--actions">
          <button class="primary-button--blue box-circle placemarks__add">
            <svg v-svg symbol="plus"></svg>
          </button>
        </div>
      </div>
      <!-- <PlacemarkEdit/> -->
      <div class="placemarks__list custom-scrollbar">
        <Placemark v-for="index in placemarksList" :key="index" :index="index"/>
      </div>
    </div>
    <Keypress/>
  </div>
</template>

<script>
import { mapState } from 'vuex';

import Placemark from './Placemark/Placemark.vue';
import Map from './Map/Map.vue';
import Keypress from './Keypress.vue';

export default {
  name: 'ManageKMZ',
  components: {
    Placemark,
    Map,
    Keypress,
  },
  props: {
    msg: String,
  },
  computed: {
    ...mapState({
      placemarksList: (state) => state.placemarks.list,
    }),
  },
};
</script>

<style scoped lang="scss">
  @import '../../assets/style/_main.scss';
  @import './Placemarks.scss';

  .manage-kmz {
    $header-height: 5rem;
    $margin-top-bottom: 1rem;
    display: grid;
    grid-template-columns: 3fr 60rem;
    grid-template-rows: 5rem calc(
      100vh - calc(#{$header-height} + calc(#{$margin-top-bottom} * 2))
    );

    margin: $margin-top-bottom 2rem $margin-top-bottom 2rem;

    & .header {
      grid-column-start: 1;
      grid-column-end: 2;
      align-self: center;

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
  }
</style>
