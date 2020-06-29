<template>
  <div id="app">
     <!-- oncontextmenu="return false" -> Disable right click -->
    <router-view></router-view>
    <notifications position="bottom right"/>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {},
};
</script>

<style lang="scss">
</style>
