<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Login from './views/LoginView.vue'
</script>
<script>
  export default {}
</script>

<template>
  <router-view />
  <main>
  </main>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>