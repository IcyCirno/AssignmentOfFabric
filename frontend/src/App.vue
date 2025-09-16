<template>
  <div id="app">
    <NavBar />
    <router-view />
    <Alert ref="alertComponent" />
  </div>
</template>

<script>
import NavBar from './components/NavBar.vue'
import Alert from "./components/Alert.vue";

export default {
  name: 'App',
  components: { Alert ,NavBar},
  methods: {
    showAlert(type, message) {
      this.$refs.alertComponent.createAlert(type, message);
    }
  }
}

</script>

<style>
/* 可以在这里加全局样式 */
body {
  margin: 0;
  font-family: Arial, sans-serif;
}
</style>
