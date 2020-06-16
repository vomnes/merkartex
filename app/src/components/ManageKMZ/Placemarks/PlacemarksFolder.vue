<template>
  <div class="placemarks-folder">
    <div class="placemarks-folder--header">
      <div class="placemarks-folder--header--icon">
        <svg class="placemarks-folder--header--folder" v-svg symbol="folder"></svg>
      </div>
      <h1 class="text__title">{{ title }}</h1>
      <div class="placemarks-folder--header--icon" @click="open = !open">
        <svg
          class="placemarks-folder--header--arrow"
          :class="{ 'open__rotate': open }"
          v-svg symbol="arrow-down"></svg>
      </div>
    </div>
    <transition name="list-fade">
      <div v-if="open">
        <slot></slot>
        <div
          class="placemarks-folder__limit"
          title="Close"
          @click="open = false">
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  name: 'PlacemarksFolder',
  props: {
    title: String,
  },
  data() {
    return {
      open: false,
    };
  },
};
</script>

<style lang="scss" scoped>
@import '@/assets/style/_main.scss';

.placemarks-folder {
  &--header {
    $header-height: 2.5rem;

    background: #FAFAFA;

    height: $header-height;
    display: flex;
    align-items: center;

    $space-top-bottom: 2rem;
    padding: $space-top-bottom 0;
    position: sticky;
    top: -.1rem;
    z-index: 1;

    &--icon {
      height: $header-height;
    }

    &--folder {
      width: 2.2rem;
    }

    & > h1 {
      font-weight: 600;
      font-size: 2rem;

      margin: 0 1.5rem 0 1rem;
    }

    &--arrow {
      width: 1.3rem;
      cursor: pointer;
      transition: .5s;
    }
  }

  &__limit {
    width: 5rem;
    height: .3rem;
    border-radius: .15rem;
    background: $color-grey-2;

    margin: 0 2rem;

    left: 50%;
    transform: translateX(-100%);
    position: relative;

    cursor: pointer;
  }
}

.open__rotate {
  transform: rotate(180deg);
  transition: .5s;
}

.list-fade-enter-active {
  transition: all .5s;
}
.list-fade-leave-active {
  transition: all .3s;
}
.list-fade-enter {
  transform: translateX(-100%);
}

.list-fade-leave-to {
  transform: translateY(100%);
}

</style>
