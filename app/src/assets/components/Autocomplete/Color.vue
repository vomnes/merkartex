<template>
  <div class="autocomplete-color">
    <div class="autocomplete-color__current-color">
      <span class="selected-color" :style="`background-color: ${value.color}`"></span>
    </div>
    <md-autocomplete
      v-model="value"
      :md-options="list.map(x => formatValue(x))"
      :md-selected="selected(value)">
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
  props: ['colors', 'current'],
  data() {
    return {
      value: formatValue(this.current),
      list: this.colors,
      formatValue,
    };
  },
  methods: {
    selected(value) {
      this.$emit('sendValue', 'color', { name: value.name, color: value.color });
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
