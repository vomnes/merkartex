<template>
  <div class="keypress">
    <p class="text__details" v-if="shiftPressed">Shift</p>
    <p class="text__details" v-if="windowPressed">Windows / ⌘</p>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  name: 'Keypress',
  computed: {
    ...mapState({
      shiftPressed: (state) => state.keypress.shift,
      windowPressed: (state) => state.keypress.window,
    }),
  },
  methods: {
    ...mapActions('keypress', ['setKeypressStatus']),
  },
  mounted() {
    window.addEventListener('keydown', (e) => {
      if (e.keyCode === 16) {
        this.setKeypressStatus({ type: 'shift', status: true });
      } else if (e.keyCode === 91 || e.keyCode === 224) {
        this.setKeypressStatus({ type: 'window', status: true });
      }
    }, this);
    window.addEventListener('keyup', (e) => {
      if (e.keyCode === 16) {
        this.setKeypressStatus({ type: 'shift', status: false });
      } else if (e.keyCode === 91 || e.keyCode === 224) {
        this.setKeypressStatus({ type: 'window', status: false });
      }
    }, this);
  },
};
</script>

<style lang="scss" scoped>
  @import '@/assets/style/_main.scss';

  .keypress {
    position: fixed;
    bottom: 1rem;
    left: 1rem;

    z-index: 9999;

    display: flex;

    & p {
      color: $color-primary-white;
      background: $color-primary;
      white-space: nowrap;

      @extend .text__details;
      font-size: 1.1rem;
      @extend .box-round-corner;
      padding: .55rem;
    }

    & > * {
      &:not(:last-child) {
        margin-right: 1rem;
      }

      cursor: default;
    }
  }
</style>
