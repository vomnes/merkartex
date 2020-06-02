<template>
  <div id="app">
    <ManageKMZ msg="Hello"/>
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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
