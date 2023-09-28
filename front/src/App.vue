<template>
  <!-- <Home msg="Bienvenido" /> -->
  <router-view></router-view>
  <!-- <h1>Hola</h1> -->
</template>
<script lang="js">

import { defineComponent, inject, computed, ref } from 'vue'

export default defineComponent({
  name: 'App',
  props: {
    msg: String,
  },
  setup(props) {
    const count = ref(0)
    return {
      count,
    }
  },
})

</script>


<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
