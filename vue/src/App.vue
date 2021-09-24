<template>
    <div id="app">
      <div class="buttonsContainer">
        <!-- <div id="btReload" v-on:click="reloadComponent"><i class="fas fa-sync-alt"></i></div> -->
        <button id="btLog" v-on:click="viewLog"><i class="fas fa-list-ul"></i></button>
        <button id="btRoom" v-on:click="viewRoom"><i class="fas fa-door-open"></i></button>
        <button id="btUser" v-on:click="viewUser"><i class ="fas fa-user"></i></button>
      </div>
      <div :key="reloadKey">
      <component :is="view" class="mainContainer"></component>
      </div>
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
    view:"room",
    reloadKey:0,
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
    },
    reloadComponent(){
      this.reloadKey++;
    }
  },
  mounted(){
      document.title = '入室者一覧';
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
  display:flex;
  flex-direction:column;
}
.mainContainer{
  padding: 1rem 1rem 2rem;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}
@media(max-width:560px){
  h1{
    font-size: 5.5vw;
  }
  :root{
    font-size: 4vw;
  }
  .buttonsContainer button{
    font-size: 8.5vw;
    width: 100%;
  }
  .buttonsContainer{
    position:fixed;
    bottom:0;left:0;
    width:calc(100% - 10px);
    padding: 5px 5px 0;
    background-color: white;
    justify-content: space-evenly;
  }
  .mainContainer{
    min-height:calc(100vh - 9vw);
  }
}
@media(min-width:561px){
  .buttonsContainer button{
    font-size: 1.5rem;
  }
  #app{
    margin-top: 30px;
  }
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
#btReload{
  color: #2c3e50;
	background-color: #DDD;
  margin: 0. 0.25rem;
  padding: 0.125rem 0.75rem;
  border-radius: 1rem;
  border-width: 3px;
  border-style: solid;
  font-weight: bold;
  font-size: 1rem;
}
#btReload i{
  vertical-align:middle;
}
</style>
