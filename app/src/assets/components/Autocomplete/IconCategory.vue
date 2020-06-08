<template>
  <div class="autocomplete-icon-category">
    <div class="autocomplete-icon-category__current-icon">
      <span class="selected-icon" >{{ value.icon }}</span>
    </div>
    <md-autocomplete
      v-model="value"
      :md-options="list.map(x => formatValue(x))"
      :md-selected="selected(value)">
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
  props: ['categories', 'current'],
  data() {
    return {
      value: formatValue(this.current),
      list: this.categories,
      formatValue,
    };
  },
  methods: {
    selected(value) {
      this.$emit('sendValue', 'category', { name: value.name, color: value.color });
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
