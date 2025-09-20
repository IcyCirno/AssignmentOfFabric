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
    },
    showAlert(type, message) {
      this.$refs.alertComponent.createAlert(type, message);
    }
  },
  // provide 必须返回一个对象；将 showAlert 绑定到 this，避免调用时 this 丢失
  provide() {
      console.log("App provide 注册 showAlert");

    return {
      showAlert: (type, message) => this.showAlert(type, message)
    };
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
