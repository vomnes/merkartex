<template>
  <md-dialog :md-active="open" @md-clicked-outside="$emit('manageOpen', false)">
    <div class="modale">
      <div class="modale__header">
        <h1 class="text__title-level-2">Edit - Description</h1>
        <svg v-svg symbol="close" @click="$emit('manageOpen', false)"></svg>
      </div>
      <div class="modale__content">
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
  </md-dialog>
</template>

<script>
export default {
  name: 'ModaleEditDescription',
  props: {
    open: Boolean,
  },
  data() {
    return {
      description: 'No content',
    };
  },
  methods: {
  },
};
</script>

<style lang="scss" scoped>
.modale {
  &__content {
    width: 50rem;
    height: 40rem;

    & textarea {
      width: 100%;
      height: 100%;
      resize: none;
    }
  }
}

</style>
