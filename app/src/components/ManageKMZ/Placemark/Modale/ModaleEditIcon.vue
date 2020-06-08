<template>
  <md-dialog :md-active="open" @md-clicked-outside="$emit('manageOpen', false)">
    <div class="modale">
      <div class="modale__header">
        <h1 class="text__title-level-2">Edit icon</h1>
        <svg v-svg symbol="close" @click="$emit('manageOpen', false)"></svg>
      </div>
      <div class="modale__content text__details">
        <AutocompleteColor/>
        <AutocompleteIconCategory/>
      </div>
      <div class="modale__footer modale__item--right">
        <button class="primary-button--green text__details box-round-corner">Save</button>
      </div>
    </div>
  </md-dialog>
</template>

<script>
import AutocompleteColor from '../../../../assets/components/Autocomplete/Color.vue';
import AutocompleteIconCategory from '../../../../assets/components/Autocomplete/IconCategory.vue';

export default {
  name: 'ModaleEditIcon',
  components: {
    AutocompleteColor,
    AutocompleteIconCategory,
  },
  props: [
    'open',
  ],
  data() {
    return {
    };
  },
};
</script>

<style lang="scss" scoped>
</style>
