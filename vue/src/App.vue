<template>
    <div id="app">
      <div class="buttonsContainer">
        <button id="btLog" v-on:click="viewLog">ログ</button>
        <button id="btRoom" v-on:click="viewRoom">入室者一覧</button>
        <button id="btUser" v-on:click="viewUser">ユーザー情報</button>
      </div>
      <component :is="view"></component>
    </div>
</template>

<script>
import Room from './components/inRoom.vue';
import Log from './components/log.vue';
import User from './components/userPage.vue';

export default {
  name: 'App',
  data: () => ({
    // apiのホストアドレス
    // host: 'http://localhost:8081',
    host: 'http://172.16.6.4:8081',
    loginData:{
      Email:'',
      Password:'hoge',
    },
    view:"log",
  }),
  components: {
    Room,
    Log,
    User
  },
  methods:{
    viewLog(){
      this.view='log';
      document.title = 'ログ';
    },
    viewRoom(){
      this.view='room';
      document.title = '入室者一覧';
    },
    viewUser(){
      this.view='user';
      document.title = 'ユーザー情報';
    }
  },
  mounted(){
      document.title = 'ログ';
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
  margin-top: 30px;
}
.buttonsContainer{
  display: flex;
  
}
.buttonsContainer button{
  color: #2c3e50;
  margin: 0 0.25rem;
  padding: 0.25rem 1rem;
  border-radius: 0.25rem 0.25rem 0 0;
  border-width: 3px 3px 0;
  border-style: solid;
  font-weight: bold;
  font-size: 1.5rem;
}
button:focus{
  outline: 0;
}
#btLog{
  background-color: #6DD;
  border-color: #0AD;
}
#btRoom{
  background-color: #5EA;
  border-color: #0C5;
}
#btUser{
	background-color: #FD0;
  border-color:#EC0;
}
</style>
