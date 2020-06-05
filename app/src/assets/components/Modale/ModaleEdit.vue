<template>
  <div>
  <div class="modale-background" @click.stop.prevent="clickBackground">
  </div>
  <div class="modale" :class="{ 'modale--move': modaleMove }">
    <div class="modale__header">
      <h1 class="text__title-level-2">Edit - Description</h1>
      <svg v-svg symbol="close"></svg>
    </div>
    <div class="modale__block">
      <textarea class="text__body custom-scrollbar" v-model="description"></textarea>
    </div>
    <div class="modale__footer">
      <p class="modale__footer--detail text__details">Description as texte or HTML</p>
      <div>
        <button class="primary-button--blue text__details box-round-corner">Preview</button>
        <button class="primary-button--green text__details box-round-corner">Set</button>
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: 'ModaleEdit',
  data() {
    return {
      description: 'No content',
      modaleMove: false,
    };
  },
  methods: {
    clickBackground() {
      this.modaleMove = true;
      setTimeout(() => { this.modaleMove = false; }, 250);
    },
  },
};
</script>

<style lang="scss" scoped>
  @import "ModaleEdit";
</style>
