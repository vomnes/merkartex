<template>
  <div class="manage-kmz">
    <div class="header">
      <div class="header--title">
        <svg v-svg symbol="kmz"></svg>
        <h1 @click="toggleEditTitle(true)" v-if="!editTitle.active">{{ title }}</h1>
        <input
          type="text"
          v-model="editTitle.content"
          :style="{width: `${editTitle.content.length * 1.75}rem`}"
          v-else>
        <button
          v-if="editTitle.active"
          class="header--title__input--close">
          <svg v-svg symbol="close"></svg>
        </button>
        <button
          v-if="editTitle.active"
          class="header--title__input--confirm">
          <svg v-svg symbol="confirm"></svg>
        </button>
      </div>
    </div>
    <Map/>
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
        <Placemark v-for="index in placemarksList" :key="index" :index="index"/>
      </div>
    </div>
    <Keypress/>
  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex';

import Placemark from './Placemark/Placemark.vue';
// import PlacemarkEdit from './Placemark/PlacemarkEdit.vue';
import Map from './Map/Map.vue';
import Keypress from './Keypress.vue';
import ModaleEditMultiple from './ModaleEditMultiple.vue';

export default {
  name: 'ManageKMZ',
  components: {
    Placemark,
    Map,
    Keypress,
    ModaleEditMultiple,
  },
  props: {
    msg: String,
  },
  data() {
    return {
      open: {
        editMultiple: false,
      },
      title: 'Shanghai',
      editTitle: {
        active: false,
        content: '',
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
    toggleEditTitle(value) {
      this.editTitle.active = value;
      if (value) {
        this.editTitle.content = this.title;
      }
    },
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
      grid-column-end: 3;
      align-self: center;

      &--title {
        display: flex;
        align-items: center;

        & > h1, & > input {
          font-family: $font-3;
          font-weight: 400;
          font-size: 2.4rem;
          background: none;
          border-bottom: .1rem solid transparent;
        }

        & > input {
          border-bottom: .1rem dashed $color-primary;
          min-width: 15rem;
          max-width: 50rem;
          padding-bottom: .5rem;
        }

        & > svg {
          width: 3rem;
          height: 4.2rem;
          margin-right: 1.5rem;
        }

        & > button {
          margin-left: .75rem;
          color: grey;
          cursor: pointer;
          transition: .5s;

          &:hover {
            transition: .5s;
          }
        }

        &__input {
          &--confirm {
            width: 2rem;
            height: 2rem;
            &:hover, &:focus {
              color: rgba($color-primary-green, .4);
            }
          }
          &--close {
            width: 1.5rem;
            height: 1.5rem;
            &:hover, &:focus {
              color: rgba($color-warning, .4);
            }
          }
        }
      }
    }
  }
</style>
