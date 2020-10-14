<template>
    <div class="hoge">
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
    <!-- <script>
        // // jsonの取得(/user)
        // var reqest = new XMLHttpRequest();
        // var userData;
        // reqest.open('GET','/user',true);// apiに送るリクエストのURL指定
        // reqest.responseType = 'json';
        // reqest.onload = function () {
        //     useData = this.response;
        // }
        // reqest.send();
    </script> -->
    <!-- <script>
        new Vue({
            el:"#app",
            data:{
                // POSTで送信するJSONの要素
                frm:{
                    Studentnumber:"",
                    Email:"",
                    Password:"",
                    Name:"",
                    UID:"",
                },
            },
            methods:{
                regist(){
                    // jsonを送信する
                    var reqPost = new XMLHttpRequest();
                    reqPost.open('POST','/user',true);// apiに送るリクエストのURL指定
                    reqPost.setRequestHeader('Content-Type','application/json');
                    reqPost.send(JSON.stringify(this.frm));
                }
            },
        })
    </script> -->
</template>
<script>
export default {
  name: 'Hoge',
  data: ()=>({
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
      reqPost.open('POST','./user:8081',true);// apiに送るリクエストのURL指定
      reqPost.setRequestHeader('Content-Type','application/json');
      reqPost.send(JSON.stringify(this.frm));
    },
    getjson(){
      // jsonの取得(/user)
      var reqest = new XMLHttpRequest();
      var userData;
      reqest.open('GET','localhost:8081/user',true);// apiに送るリクエストのURL指定
      reqest.responseType = 'json';
      reqest.onload = function () {
          userData = this.response;
          window.alert(userData);
      }
      reqest.send();
    }
  }
  ,
}
</script>