<template>
  <div id="app">
     <!-- oncontextmenu="return false" -> Disable right click -->
    <!-- <ManageKMZ/> -->
    <LoadMap/>
    <notifications position="bottom right"/>
  </div>
</template>

<script>
// import ManageKMZ from './components/ManageKMZ/Main.vue';
import LoadMap from './components/LoadMap/LoadMap.vue';

export default {
  name: 'App',
  components: {
    // ManageKMZ,
    LoadMap,
  },
};
</script>

<style lang="scss">
</style>
