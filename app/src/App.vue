<template>
  <div id="app">
     <!-- oncontextmenu="return false" -> Disable right click -->
    <ManageKMZ/>
    <notifications position="bottom right"/>
  </div>
</template>

<script>
import ManageKMZ from './components/ManageKMZ/Main.vue';

export default {
  name: 'App',
  components: {
    ManageKMZ,
  },
};
</script>

<style lang="scss">
</style>
