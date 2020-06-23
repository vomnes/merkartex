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
        <button class="primary-button--blue box-circle" @click="open.addNewPlacemark = true">
          <svg v-svg symbol="plus"></svg>
        </button>
      </div>
      <ModaleEditMultiple
        :open="open.editMultiple"
         @manageOpen="manageOpenEditMultiple"/>
    </div>
    <div class="placemarks__list custom-scrollbar">
      <!-- <PlacemarksFolder title="Shanghai - Old town"></PlacemarksFolder> -->
      <Placemark
        v-for="(placemarkContent, index) in getPlacemarks"
        :key="index" :index="index" :data="placemarkContent"/>
    </div>
    <PlacemarkEdit
      :data="emptyPlacemark"
      :openEdit="open.addNewPlacemark"
      type="new"
      @manageOpen="manageToggleNewPlacemark"/>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

import emptyPlacemark from 'assets/data/empty-placemark.json';

import Placemark from './Placemark/Placemark.vue';
import ModaleEditMultiple from './ModaleEditMultiple.vue';
import PlacemarkEdit from './Placemark/PlacemarkEdit.vue';
// import PlacemarksFolder from './PlacemarksFolder.vue';

export default {
  name: 'Placemarks',
  components: {
    Placemark,
    ModaleEditMultiple,
    PlacemarkEdit,
  },
  data() {
    return {
      open: {
        editMultiple: false,
        addNewPlacemark: false,
      },
      emptyPlacemark,
    };
  },
  computed: {
    ...mapGetters('placemarks', [
      'hasPlacemarksSelection',
      'getPlacemarks',
    ]),
  },
  methods: {
    manageOpenEditMultiple(value) {
      this.open.editMultiple = value;
    },
    manageToggleNewPlacemark(value) {
      this.open.addNewPlacemark = value;
    },
  },
};
</script>

<style lang="scss" scoped>
  @import '@/assets/style/_main.scss';
  @import './Placemarks.scss';
</style>
