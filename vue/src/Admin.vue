<template>
  <div id="admin">
    <div class=adminButtons>
      <router-link class="btadmin" to="/">通常ページに戻る</router-link>
      <router-link class="btadmin" to="/admin">トップ</router-link>
      <router-link class="btadmin" to="/admin/log">アナリティクス</router-link>
    </div>
    <router-view class="adminRouter"/>
  </div>
</template>

<script>
export default {
  name: 'Admin',
  data: () => ({
    
  }),
  mounted(){
      document.title = '管理者用';
  }
}
</script>

<style>
  body{
    margin:0;
  }
  #admin{
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    display:flex;
    flex-direction:column;
  }
  .adminButtons{
    display: flex;
    justify-content:center;
    margin-top:5;
  }
  .btadmin{
    color: #2c3e50;
    margin: 0 0.25rem;
    padding: 0.25rem 1rem;
    border-radius: 0.25rem 0.25rem 0 0;
    border-width: 3px 3px 0;
    border-style: solid;
    font-weight: bold;
    text-decoration: none;
  }
</style>
