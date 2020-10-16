<template>
    <div id="app">
        <Room></Room>
        <Log></Log>
        <!-- <component :is="currentView"></component> -->
    </div>
</template>

<script>
import Room from './components/inRoom.vue';
import Log from './components/log.vue';

export default {
  name: 'App',
  data: () => ({
    // apiのホストアドレス
    host: 'http://172.16.6.4',
  }),
  components: {
    Room,
    Log
  },
  methods:{
    viewlog(){

    }
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
