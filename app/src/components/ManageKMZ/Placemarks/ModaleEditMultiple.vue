<template>
  <md-dialog :md-active="open" @md-clicked-outside="$emit('manageOpen', false)">
    <div class="modale">
      <div class="modale__header">
        <h1 class="text__title-level-2">Edit placemarks</h1>
        <svg v-svg symbol="close" @click="$emit('manageOpen', false)"></svg>
      </div>
      <div class="modale__content">
        <div class="edit-multiple__section">
          <h2 class="text__title-level-2">Title</h2>
          <div class="edit-multiple__section--row">
            <input
              class="text__details edit-multiple__section--row--before"
              type="text" name="" value="" placeholder="Before" v-model="title.before">
            <p class="edit-multiple__section--row--main">Placemark Title</p>
            <input
              class="text__details edit-multiple__section--row--after"
              type="text" name="" value="" placeholder="After" v-model="title.after">
          </div>
        </div>
        <div class="edit-multiple__section">
          <h2 class="text__title-level-2">Icon style</h2>
          <div class="edit-multiple__section--row">
            <AutocompleteColor
            :colors="placemarksDesign.colors"
            :current="placemarkStyle.color"
            @sendValue="receiveNewValue"/>
            <AutocompleteIconCategory
            :categories="placemarksDesign.categories"
            :current="placemarkStyle.category"
            @sendValue="receiveNewValue"/>
          </div>
        </div>
      </div>
      <div class="modale__footer">
        <p class="modale__footer--detail text__details">
          Update the selected placemarks
        </p>
        <div>
          <button
            class="primary-button--green text__details box-round-corner"
            @click="manageChangeClose"
            >Update</button>
        </div>
      </div>
    </div>
  </md-dialog>
</template>

<script>
import { mapActions } from 'vuex';
import AutocompleteColor from 'assets/components/Autocomplete/Color.vue';
import AutocompleteIconCategory from 'assets/components/Autocomplete/IconCategory.vue';
import placemarksDesign from 'assets/data/placemarks-design.json';

export default {
  name: 'ModaleEditMultiple',
  components: {
    AutocompleteColor,
    AutocompleteIconCategory,
  },
  props: {
    open: Boolean,
  },
  data() {
    return {
      placemarksDesign,
      title: {
        before: null,
        after: null,
      },
      placemarkStyle: {
        color: null,
        category: null,
      },
    };
  },
  watch: {
    open(value) {
      if (value) {
        this.placemarkStyle = {
          color: null,
          category: null,
        };
      }
    },
  },
  methods: {
    ...mapActions('placemarks', [
      'editMultiplePlacemarks',
      'unselectAllPlacemarks',
    ]),
    manageChangeClose() {
      this.$emit('manageOpen', false);
      this.editMultiplePlacemarks({
        title: this.title,
        icon: {
          style: this.placemarkStyle.color.name,
          category: this.placemarkStyle.category.name,
        },
      });
      this.unselectAllPlacemarks(0);
    },
    receiveNewValue(type, value) {
      this.placemarkStyle[type] = value;
    },
  },
};
</script>

<style lang="scss" scoped>
.edit-multiple {
  &__section {
    & h2 {
      text-align: center;
      margin-bottom: 1.5rem;
    }

    &--row {
      display: flex;
      justify-content: center;

      &--before {
        text-align: right;
        width: 15rem;
      }

      &--main {
        color: #888888;
        background: #f0f0f0;

        margin: 0 .75rem 0 .75rem;
        padding: .25rem .75rem .25rem .75rem;
        border-radius: 2rem;

        cursor: default;
      }

      &--after {
        text-align: left;
        width: 15rem;
      }
    }

    &:not(:first-child) {
      margin-top: 2rem;
    }
  }
}
</style>
