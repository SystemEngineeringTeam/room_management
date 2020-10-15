<template>
    <div id="app">
        <Room></Room>
        <Log></Log>
    </div>
</template>

<script>
import Room from './components/inRoom.vue';
import Log from './components/log.vue';

export default {
  name: 'App',
  data: () => ({
    host: 'http://localhost:8081',
    page:'room',
  }),
  components: {
    Room,
    Log
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
