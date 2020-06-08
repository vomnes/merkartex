<template>
  <div class="autocomplete-color">
    <div class="autocomplete-color__current-color">
      <span class="selected-color" :style="`background-color: ${value.color}`"></span>
    </div>
    <md-autocomplete
      v-model="value"
      :md-options="colors.map(x => formatValue(x))">
      <label>
        Color
      </label>

      <template slot="md-autocomplete-item" slot-scope="{ item }">
        <span class="selected-color" :style="`background-color: ${item.color}`"></span>
        <md-highlight-text :md-term="item.name">{{ item.name }}</md-highlight-text>
      </template>

      <template slot="md-autocomplete-empty" slot-scope="{ term }">
        Color "{{ term }}" not handled
      </template>
    </md-autocomplete>
  </div>
</template>

<script>
const formatValue = ({ name, color }) => {
  const toLowerCase = () => name.toLowerCase();
  const toString = () => name;
  return {
    name,
    color,
    toLowerCase,
    toString,
  };
};

export default {
  name: 'AutocompleteColor',
  data: () => ({
    value: formatValue({ name: 'red', color: '#e31a23' }),
    colors: [
      { name: 'red', color: '#e31a23' },
      { name: 'pink', color: '#fa4281' },
      { name: 'purple', color: '#9e28af' },
      { name: 'deep_purple', color: '#9e28af' },
      { name: 'blue', color: '#0171c4' },
      { name: 'light_blue', color: '#1aa0f6' },
      { name: 'cyan', color: '#12c3c5' },
      { name: 'teal', color: '#07a18d' },
      { name: 'green', color: '#378c3c' },
      { name: 'lime', color: '#93bb36' },
      { name: 'yellow', color: '#fbc60a' },
      { name: 'orange', color: '#fe9f00' },
      { name: 'deep_orange', color: '#fb622a' },
      { name: 'brown', color: '#7d5342' },
      { name: 'gray', color: '#737373' },
      { name: 'blue_gray', color: '#577480' },
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
  .autocomplete-color {
    display: flex;

     &__current-color {
      display: flex;
      align-items: flex-end;
      margin-bottom: .5rem;

      .selected-color {
        width: 1.75rem;
        height: 1.75rem;
        border-radius: 1.75rem;
      }
    }
  }
</style>

<style lang="scss">
  @import '../../style/main.scss';

  .selected-color {
    width: 1.6rem;
    height: 1.6rem;
    display: inline-block;
    margin-right: 1.6rem;
    border-radius: 50%;
    box-shadow: $box-shadow-1;
    color: currentColor;
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
