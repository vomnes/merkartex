<template>
  <div class="autocomplete-icon-category">
    <div class="autocomplete-icon-category__current-icon">
      <span class="selected-icon" >{{ value.icon }}</span>
    </div>
    <md-autocomplete
      v-model="value"
      :md-options="list.map(x => formatValue(x))">
      <label>
        Category
      </label>

      <template slot="md-autocomplete-item" slot-scope="{ item }">
        <span class="autocomplete-icon-category__list--item">{{ item.icon }}</span>
        <md-highlight-text :md-term="item.name">{{ item.name }}</md-highlight-text>
      </template>

      <template slot="md-autocomplete-empty" slot-scope="{ term }">
        Category "{{ term }}" not handled
      </template>
    </md-autocomplete>
  </div>
</template>

<script>
const formatValue = ({ name, icon }) => {
  const toLowerCase = () => name.toLowerCase();
  const toString = () => name;
  return {
    name,
    icon,
    toLowerCase,
    toString,
  };
};

export default {
  name: 'AutocompleteIconCategory',
  data: () => ({
    value: formatValue({ name: 'None', icon: '⨀' }),
    list: [
      { name: 'None', icon: '⨀' },
      { name: 'Hotel', icon: '🏨' },
      { name: 'Animals', icon: '🐾' },
      { name: 'Buddhism', icon: '☸️' },
      { name: 'Building', icon: '🏢' },
      { name: 'Christianity', icon: '✝️' },
      { name: 'Entertainment', icon: '🎬' },
      { name: 'Exchange', icon: '💸' },
      { name: 'Food', icon: '🍽' },
      { name: 'Gas', icon: '⛽' },
      { name: 'Judaism', icon: '✡️' },
      { name: 'Medicine', icon: '🏥' },
      { name: 'Mountain', icon: '🏔' },
      { name: 'Museum', icon: '🖼' },
      { name: 'Islam', icon: '☪️' },
      { name: 'Park', icon: '🌳' },
      { name: 'Parking', icon: '🚗' },
      { name: 'Shop', icon: '🛍' },
      { name: 'Sights', icon: '🌄' },
      { name: 'Swim', icon: '🏊🏻' },
      { name: 'Water', icon: '🌊' },
      { name: 'Bar', icon: '🍻' },
      { name: 'Transport', icon: '🚆' },
      { name: 'Viewpoint', icon: '🌅' },
      { name: 'Sport', icon: '🎾' },
      { name: 'Finish', icon: '🏁' },
    ],
    formatValue,
  }),
  methods: {
    noop() {
      window.alert('noop');
    },
  },
};
</script>

<style lang="scss" scoped>
  .autocomplete-icon-category {
    display: flex;

     &__current-icon {
      display: flex;
      align-items: flex-end;
      margin-bottom: 1rem;

      &__list--item {
        margin-right: 1rem;
      }
    }
  }
</style>

<style lang="scss">
  @import '../../style/main.scss';

  .autocomplete-icon-category {
    &__list--item {
      margin-right: 1rem;
    }

    &__current-icon {
      .selected-icon {
        width: 1.6rem;
        height: 1.6rem;
        display: inline-block;
        margin-right: 1.6rem;
        border-radius: 50%;
      }
    }
  }

  .md-menu-content {
    z-index: 11;
  }

  .md-list.md-theme-default .md-autocomplete-items .md-highlight-text-match {
    color: $color-primary;
    @extend .text__details;
    text-transform: capitalize;
  }

  .md-field {
    margin: 0;
  }

  :root {
    --md-theme-default-primary: #3A7BD5;
  }
</style>
