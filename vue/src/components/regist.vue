<template>
    <div class="regist">
      <table>
        <thead></thead>
        <tbody>
          <tr>
            <th>
              <label for="name">名前</label>
            </th>
            <th>
              <input type="text" id="name" v-model="frm.Name">
            </th>
          </tr>
          <tr>
            <th>
              <label for="pass">パスワード</label>
            </th>
            <th>
              <input type="password" v-model="frm.Password">
            </th>
          </tr>
          <tr>
            <th>
              <label for="strNum">学籍番号</label>
            </th>
            <th>
              <input type="text" id="strNum" v-model="frm.Studentnumber">
            </th>
          </tr>
          <tr>
            <th>
              <label for="Email">メアド</label>
            </th>
            <th>
              <input type="email" id="email" v-model="frm.Email">
            </th>
          </tr>
          <input type="hidden" v-model="frm.UID">
        </tbody>
      </table>
      <button @click="regist">登録</button>
    </div>
</template>

<script>
export default {
  name: 'regist',
  data: ()=>({
    jsonData: null,
    time:null,
    frm:{
      Name:'',
      Password:'',
      Studentnumber:'',
      Email:'',
      UID:''
    }
  }),
  methods:{
    regist(){
      // jsonを送信する
      var reqPost = new XMLHttpRequest();
      reqPost.open('POST','http://localhost:8081/user',true);// apiに送るリクエストのURL指定
      reqPost.setRequestHeader('Content-Type','application/json');
      reqPost.send(JSON.stringify(this.frm));
    },
  },
  mounted(){
    this.frm.UID=this.$parent.uid;
    this.time = this.$parent.time;
  }
}
</script>
<style>
  .regist{
    margin-top: 2rem;
    background-color: #9EE;
    border-radius: 1rem;
    padding: 1.5rem 1rem;
    display: flex;
    justify-content: center;
    flex-direction: column;
  }
  .regist>button{
    font-size: 1rem;
    margin-top: 1rem;
  }
</style>