<template>
  <article
    class="placemark"
    @click.stop="placemarkActive = !placemarkActive"
    :class="[ placemarkActive ? 'inner-box-shadow' : 'outer-box-shadow']">
    <div class="placemark--side"></div>
    <div class="placemark--content">
      <div class="header">
        <h2 class="text__title">Jardin Yuyuan</h2>
        <div class="header--icon" v-on:click.prevent="toggleMoreOptionsOpen">
          <svg v-svg symbol="ellipsis"
          tabindex="0" v-on:keyup.enter="toggleMoreOptionsOpen"></svg>
          <div
            v-if="moreOptionsOpen"
            class="placemark--more-options">
            <ul>
              <li tabindex="0">
                <svg v-svg symbol="pencil"></svg>
                <p>Edit</p>
              </li>
              <li tabindex="0">
                <svg v-svg symbol="trash"></svg>
                <p>Remove</p>
              </li>
              <li tabindex="0">
                <svg v-svg symbol="duplicate"></svg>
                <p>Duplicate</p>
              </li>
            </ul>
          </div>
        </div>
      </div>
      <p>
        {{ descriptionContent }}
        <span
          v-if="hasSeeMoreDescription"
          @click="toggleDescriptionOpen"
          @keyup.enter="toggleDescriptionOpen"
          tabindex="0">
          {{ !descriptionOpen ? 'See more' : 'Close' }}
        </span>
      </p>
      <div class="footer">
        <p>Quartier</p>
        <p>05/31/2019</p>
      </div>
    </div>
  </article>
</template>

<script>
const LIMIT_SIZE = 256;

export default {
  name: 'Placemark',
  data() {
    return {
      description: 'Le jardin Yuyuan est un jardin de deux hectares datant du XVIe siècle situé au centre de la Vieille Ville près de Chenghuangmiao à Shanghai, en Chine. Le jardin Yuyuan est un jardin de deux hectares datant du XVIe siècle situé au centre de la Vieille Ville près de Chenghuangmiao à Shanghai, en Chine.',
      descriptionOpen: false,
      moreOptionsOpen: false,
      placemarkActive: false,
    };
  },
  computed: {
    hasSeeMoreDescription() {
      return this.description.length > LIMIT_SIZE;
    },
    descriptionContent() {
      return this.hasSeeMoreDescription && !this.descriptionOpen
        ? `${this.description.slice(0, LIMIT_SIZE)}...`
        : this.description;
    },
  },
  methods: {
    toggleDescriptionOpen() {
      this.descriptionOpen = !this.descriptionOpen;
    },
    toggleMoreOptionsOpen() {
      this.moreOptionsOpen = !this.moreOptionsOpen;
    },
  },
};
</script>

<style lang="scss" scoped>
  @import 'Placemark';
</style>
