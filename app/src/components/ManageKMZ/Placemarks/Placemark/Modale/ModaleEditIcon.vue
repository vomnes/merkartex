<template>
  <md-dialog :md-active="open" @md-clicked-outside="manageCancelClose">
    <div class="modale">
      <div class="modale__header">
        <h1 class="text__title-level-2">Edit icon</h1>
        <svg v-svg symbol="close" @click="manageCancelClose"></svg>
      </div>
      <div class="modale__content text__details">
        <AutocompleteColor
        :colors="placemarksDesign.colors"
        :current="placemark.color"
        @sendValue="receiveNewValue"/>
        <AutocompleteIconCategory
        :categories="placemarksDesign.categories"
        :current="placemark.category"
        @sendValue="receiveNewValue"/>
      </div>
      <div class="modale__footer" :class="{ 'modale__item--right': !invalid }">
        <p v-if="invalid" class="text__details" style="color: #ed213a">Field(s) can't be empty</p>
        <button
          class="primary-button--green text__details box-round-corner"
          @click="manageChangeClose">Change</button>
      </div>
    </div>
  </md-dialog>
</template>

<script>
import AutocompleteColor from 'assets/components/Autocomplete/Color.vue';
import AutocompleteIconCategory from 'assets/components/Autocomplete/IconCategory.vue';
import placemarksDesign from 'assets/data/placemarks-design.json';

const cloneDeep = require('clone-deep');

export default {
  name: 'ModaleEditIcon',
  components: {
    AutocompleteColor,
    AutocompleteIconCategory,
  },
  props: {
    open: Boolean,
    placemark: Object,
  },
  data() {
    return {
      placemarksDesign,
      currentPlacemark: cloneDeep(this.placemark),
      invalid: false,
    };
  },
  watch: {
    open(value) {
      if (value) {
        // Fix child take ownership on parent's values
        this.currentPlacemark = cloneDeep(this.placemark);
      }
    },
  },
  methods: {
    receiveNewValue(type, value) {
      this.currentPlacemark[type] = value;
    },
    manageCancelClose() {
      this.$emit('manageOpen', false);
    },
    manageChangeClose() {
      if (this.currentPlacemark.color.name == null || this.currentPlacemark.category.name == null) {
        this.invalid = true;
        return;
      }
      this.$emit('sendEditIconValue', this.currentPlacemark);
      this.$emit('manageOpen', false);
      this.invalid = false;
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
