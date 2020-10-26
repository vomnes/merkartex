<template>
  <div class="manage-kmz">
    <div class="header">
      <div class="header--title">
        <svg v-svg symbol="kmz"></svg>
        <h1 @click="toggleEditTitle(true)" v-if="!editTitle.active">{{ title }}</h1>
        <input
          type="text"
          v-model="editTitle.content"
          :style="{width: `${editTitle.content.length * 1.5}rem`}"
          v-else>
        <button
          v-if="editTitle.active"
          @click="toggleEditTitle(false)"
          class="header--title__input--close">
          <svg v-svg symbol="close"></svg>
        </button>
        <button
          v-if="editTitle.active"
          @click="saveTitle"
          class="header--title__input--confirm">
          <svg v-svg symbol="confirm"></svg>
        </button>
      </div>
      <div class="manage-kmz__header--actions">
        <transition name="fade-scale">
          <button
            v-if="getHasChanges"
            class="primary-button--blue text__details box-round-corner"
            @click="persistData">Save</button>
        </transition>
        <button
        class="primary-button--blue text__details box-round-corner"
        @click="exportData">Export</button>
      </div>
    </div>
    <Map :geoCenter="geoCenter" ref="mapComponent"/>
    <Placemarks/>
    <Keypress/>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

import Api from 'assets/library/api/index';

import Map from './Map/Map.vue';
import Placemarks from './Placemarks/Placemarks.vue';
import Keypress from './Keypress.vue';

const download = (filename, content, type) => {
  const element = document.createElement('a');
  const file = new Blob([content], { type });
  element.href = window.URL.createObjectURL(file);
  element.download = filename;

  element.style.display = 'none';
  document.body.appendChild(element);

  element.click();

  document.body.removeChild(element);
};

export default {
  name: 'ManageKMZ',
  components: {
    Map,
    Placemarks,
    Keypress,
  },
  data() {
    return {
      title: '',
      editTitle: {
        active: false,
        content: '',
      },
      geoCenter: {},
    };
  },
  beforeMount() {
    let kmzContent;
    const localStorageMap = localStorage.getItem('map_data');
    if (localStorageMap) {
      kmzContent = JSON.parse(localStorageMap);
    } else {
      this.$router.push('/');
    }
    this.title = kmzContent.name;
    this.geoCenter = kmzContent.geoCenter;
    this.setPlacemarks(kmzContent.placemarks);
    this.setPlacemarksFolders(kmzContent.folders);
    this.setTitle(kmzContent.name);
    this.setLength(kmzContent.length);
  },
  methods: {
    toggleEditTitle(value) {
      this.editTitle.active = value;
      if (value) {
        this.editTitle.content = this.title;
      }
    },
    ...mapActions('placemarks', [
      'setPlacemarks',
      'setPlacemarksFolders',
      'setTitle',
      'setLength',
      'toggleHasChanges',
    ]),
    saveTitle() {
      this.title = this.editTitle.content;
      this.setTitle(this.title);
      this.toggleEditTitle(false);
      this.toggleHasChanges(true);
    },
    persistData() {
      const newLocalStorageMap = JSON.parse(localStorage.getItem('map_data'));
      newLocalStorageMap.name = this.getTitle;
      newLocalStorageMap.placemarks = this.getPlacemarks;
      newLocalStorageMap.folders = this.getFolders;
      localStorage.setItem('map_data', JSON.stringify(newLocalStorageMap));
      this.toggleHasChanges(false);
      this.$notify({
        title: 'Data saved locally',
        type: 'success',
      });
    },
    exportData() {
      this.persistData();
      Api.exportData(localStorage.getItem('map_data'))
        .then((res) => {
          const fileData = [];
          const reader = res.body.getReader();
          reader.read()
            .then(({ value }) => {
              fileData.push(value);
            }, fileData)
            .then(() => {
              download(
                `${this.title}_${new Date().toISOString()}_map.kmz`,
                fileData[0],
                'application/octet-binary',
              );
            }, fileData);
        });
    },
  },
  computed: {
    ...mapGetters('placemarks', [
      'getPlacemarks',
      'getFolders',
      'getTitle',
      'getHasChanges',
    ]),
  },
  created() {
    window.addEventListener('beforeunload', () => {
    }, false);
  },
};
</script>

<style scoped lang="scss">
  @import '@/assets/style/_main.scss';

  .manage-kmz {
    $header-height: 5rem;
    $margin-top-bottom: 1rem;
    display: grid;
    grid-template-columns: 3fr 60rem;
    grid-template-rows: 5rem calc(
      100vh - calc(#{$header-height} + calc(#{$margin-top-bottom} * 2))
    );

    margin: $margin-top-bottom 2rem $margin-top-bottom 2rem;

    &__header {
      &--actions {
        & > button:not(:last-child) {
          margin-right: .5rem;
        }
      }
    }

    & .header {
      grid-column-start: 1;
      grid-column-end: 3;
      align-self: center;

      display: flex;
      justify-content: space-between;
      align-items: center;

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
          border-bottom: .1rem solid rgba($color-grey-1, .2);
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

    & .primary-button--blue {
      height: 3rem;
    }
  }

  .fade-scale {
    &-enter-active {
      transition: transform .5s !important;
    }
    &-leave-active {
      transition: transform .15s !important;
    }
    &-enter, &-leave-to {
      transform: scale(0) !important;
    }
  }
</style>
