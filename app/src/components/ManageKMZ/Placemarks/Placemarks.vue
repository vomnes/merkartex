<template>
  <div class="placemarks">
    <div class="placemarks__header">
      <h1 class="text__title">Placemarks</h1>
      <div class="placemarks__header--actions">
        <button
          :class="[
            hasPlacemarksSelection ? 'primary-button--blue' : 'button-passive',
            'box-circle',
          ]"
          :data-title="hasPlacemarksSelection ?
          'Edit multiple placemarks' :
          'Edit multiple placemarks using Windows / ⌘ (multi) or Shift (range) keys'"
          data-title-position="left"
          @click="hasPlacemarksSelection ? manageOpenEditMultiple(true) : ''">
          <svg v-svg symbol="layers"></svg>
        </button>
        <button class="primary-button--blue box-circle">
          <svg v-svg symbol="plus"></svg>
        </button>
      </div>
      <ModaleEditMultiple
        :open="open.editMultiple"
         @manageOpen="manageOpenEditMultiple"/>
    </div>
    <!-- <PlacemarkEdit/> -->
    <div class="placemarks__list custom-scrollbar">
      <!-- <Placemark v-for="index in placemarksList" :key="index" :index="index"/> -->
      <div class="placemarks__folder">
        <div class="placemarks__folder--header">
          <div class="placemarks__folder--header--icon">
            <svg class="placemarks__folder--header--folder" v-svg symbol="folder"></svg>
          </div>
          <h1 class="text__title">Shanghai - Old Town</h1>
          <div class="placemarks__folder--header--icon">
            <svg class="placemarks__folder--header--arrow" v-svg symbol="arrow-down"></svg>
          </div>
        </div>
        <Placemark/>
        <Placemark/>
        <Placemark/>
      </div>
      <Placemark/>
      <Placemark/>
      <Placemark/>
      <Placemark/>
      <Placemark/>
      <Placemark/>
    </div>
  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex';

import Placemark from './Placemark/Placemark.vue';
// import PlacemarkEdit from './Placemark/PlacemarkEdit.vue';
import ModaleEditMultiple from './ModaleEditMultiple.vue';

export default {
  name: 'Placemarks',
  components: {
    Placemark,
    ModaleEditMultiple,
  },
  data() {
    return {
      open: {
        editMultiple: false,
      },
    };
  },
  computed: {
    ...mapState({
      placemarksList: (state) => state.placemarks.list,
    }),
    ...mapGetters('placemarks', [
      'hasPlacemarksSelection',
    ]),
  },
  methods: {
    manageOpenEditMultiple(value) {
      this.open.editMultiple = value;
    },
  },
};
</script>

<style lang="scss" scoped>
  @import '@/assets/style/_main.scss';
  @import './Placemarks.scss';
</style>
