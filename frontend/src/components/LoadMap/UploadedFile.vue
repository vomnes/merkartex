<template>
  <div class="uploaded-file">
    <svg v-svg symbol="kmz"></svg>
    <div class="uploaded-file--content">
      <div class="uploaded-file--content__up">
        <h3 class="text__body" :title="stateSelectedFile.name">{{ filenameDisplay }}</h3>
        <svg v-svg symbol="close" @click="clear"></svg>
      </div>
      <progress
      :value="selectedFileLoadedValue" max="100"
      v-if="selectedFileLoadedValue < 100"
      :key="selectedFileLoadedValue"></progress>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  name: 'UploadedFile',
  data() {
    return {

    };
  },
  computed: {
    ...mapGetters('selectedFile', [
      'stateSelectedFile',
      'selectedFileLoadedValue',
    ]),
    filenameDisplay() {
      const lessExtension = this.stateSelectedFile.name.replace('.kmz', '');
      return lessExtension.length < 17
        ? this.stateSelectedFile.name
        : `${lessExtension.substring(0, 17)}... .kmz`;
    },
  },
  methods: {
    ...mapActions('selectedFile', [
      'selectedFileReset',
    ]),
    clear() {
      this.selectedFileReset();
      this.$emit('event', 'clearError', null);
    },
  },
};
</script>

<style lang="scss" scoped>
@import "@/assets/style/_main.scss";

.uploaded-file {
  display: flex;
  align-items: flex-end;
  height: 5rem;

  & > svg {
    height: 100%;
    max-width: 4.5rem;
    margin-right: 1.4rem;
  }

  &--content {
    width: 100%;

    display: flex;
    flex-direction: column;

    justify-content: flex-end;

    &__up {
      display: flex;
      justify-content: space-between;
      height: 2rem;

      margin-bottom: .5rem;

      & > h3 {
        font-weight: 400;
        max-width: 25rem;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }

      & > svg {
        padding: .25rem;
        width: 2rem;
        border-radius: .15rem;
        transition: all .3s;

        &:hover {
          background: $color-blue-2;
          color: $color-primary-white;
          cursor: pointer;
          padding: .35rem;
          transition: all .3s;
        }
      }
    }

    & progress {
      height: .7rem;
      width: 100%;

      border: none;

      &::-webkit-progress {
        &-bar {
          background: #A1C3E6;
        }
        &-value {
          height: 100%;
          background: linear-gradient(90deg, $color-blue 0%, #A1C3E6 200%);
          transition-delay: .4s;
          transition-duration: .5s;
          transition-timing-function: cubic-bezier(0.22, 1, 0.36, 1);
        }
      }

      &::-moz-progress {
        &-bar {
          height: 100%;
          background: linear-gradient(90deg, $color-blue 0%, #A1C3E6 200%);
        }
      }
    }

    &__bar {
      height: .7rem;
      width: 100%;

      background: #A1C3E6;

      &--filled {
        height: 100%;
        background: linear-gradient(90deg, $color-blue 0%, #A1C3E6 200%);
      }
    }
  }

  margin: 2rem 0;
}
</style>
