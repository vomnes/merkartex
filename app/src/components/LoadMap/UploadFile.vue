<template>
  <button class="primary-button--blue-white box-round-corner"
    @click="click">
    <svg v-svg symbol="arrow"></svg>
    <p class="text__details text--uppercase">{{ content }}</p>
    <input @change="fileSelected" ref="mapsmeFileInput" type="file" accept=".kmz">
  </button>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'UploadFile',
  props: {
    content: String,
    type: String,
  },
  data() {
    return {

    };
  },
  methods: {
    ...mapActions('selectedFile', [
      'setSelectedFileValue',
    ]),
    click() {
      this.$refs.mapsmeFileInput.click();
    },
    fileSelected(event) {
      const file = event.target.files[0];
      if (file && file.size > 100000) {
        this.$emit('event', 'setError', {
          error: 'File size limited to 100MB',
        });
        return;
      }
      if (file && !file.name.endsWith('.kmz')) {
        this.$emit('event', 'setError', {
          error: 'Support only KMZ files type',
        });
        return;
      }
      this.$emit('event', 'clearError', null);
      this.setSelectedFileValue({ type: 'name', value: file.name });
      this.setSelectedFileValue({ type: 'size', value: file.size });
      this.setSelectedFileValue({ type: 'valid', value: true });
      const reader = new FileReader();
      // let uploadedFile = null;
      reader.onloadstart = this.readerEvent;
      reader.onload = this.readerEvent;
      reader.onprogress = this.readerEvent;
      // uploadedFile = reader.result;
      if (file) {
        reader.readAsText(file);
      }
    },
    readerEvent(e) {
      this.setSelectedFileValue({ type: 'loaded', value: Math.ceil((e.loaded * 100) / e.total) });
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
