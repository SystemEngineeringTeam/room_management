<template>
    <div class="regist">
            <label for="name">名前</label>
            <input type="text" id="name" v-model="frm.Name">
            <br>
            <label for="pass">パスワード</label>
            <input type="password" v-model="frm.Password">
            <br>
            <label for="strNum">学籍番号</label>
            <input type="text" id="strNum" v-model="frm.Studentnumber">
            <br>
            <label for="Email">メアド</label>
            <input type="email" id="email" v-model="frm.Email">
            <br>
            <!-- <input type="hidden" v-model="frm.UID"> -->
            <label>UID(通常は隠しておく)</label>
            <input type="text" v-model="frm.UID">
            <br>
            <button @click="regist">登録</button>
            <button @click="getjson">json取得</button>
    </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'regist',
  data: ()=>({
    jsonData: null,
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
    getjson(){
      axios.get('http://localhost:8081/user')
      .then(response => {
        this.jsonData =response.data;
        // window.alert(JSON.stringify(this.jsonData));
      }).catch((e) => {
        alert(e);
      });
    }
  },
}
</script>