<template>
    <div class="userPage">
      <h1>ユーザー情報</h1>
      <div class="loginContainer" v-if="!logined">
        <form @submit.prevent="loginSubmit">
          <table>
            <h2>ログイン</h2>
            <thead></thead>
            <tbody>
              <tr>
                <th>
                  <label for="strNum">学籍番号</label>
                </th>
                <th>
                  <input type="text" id="strNum" pattern="[a-z]\d{5}" :value="Studentnumber" @input="Studentnumber = $event.target.value;onChangeStuNum();" maxlength="6" required>
                  <input type="hidden" id="email" v-model="frm.Email" maxlength='45' required readonly>
                </th>
              </tr>
              <tr>
                <th>
                  <label for="pass">パスワード</label>
                </th>
                <th>
                  <input type="text" v-model="frm.Password" required readonly>
                </th>
              </tr>
            </tbody>
          </table>
          <input class="bt" type="submit" value="登録">
        </form>
      </div>
      <div v-else>
        
      </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'User',
  data: ()=>({
    jsonData: null,
    logined:false,
    Studentnumber:'',
    frm:{
      Email:'',
      Password:'hoge',
    },
  }),
  methods:{
    onChangeStuNum(){
      var gstr=this.Studentnumber.slice(0,1);
      if(gstr.match("[A-Z]")!=null){
        gstr=gstr.toLowerCase();
        this.Studentnumber=gstr;
      }
      if(gstr==''||gstr.match("[a-z]")==null){
      this.Studentnumber='';
      this.frm.Email="example@aitech.ac.jp";
      }else{
        if (this.Studentnumber.length>1) {
          if(this.Studentnumber.slice(-1).match("[0-9]")==null){
            this.Studentnumber=this.Studentnumber.slice(0,-1);
          }
        }
      this.frm.Email=this.Studentnumber +gstr+gstr+"@aitech.ac.jp";
      }
    },
    loginSubmit(){
      // ---XMLHttpRequestで試す
      // var reqPost = new XMLHttpRequest();
      // reqPost.open('POST',this.$parent.host+'/login',true);// apiに送るリクエストのURL指定
      // reqPost.setRequestHeader('Content-Type','application/json');
      // reqPost.onload=()=>{
      //   this.jsonData=JSON.parse(reqPost.response);
      // }
      // reqPost.send(JSON.stringify(this.frm));

      axios.post(this.$parent.host+'/login',this.frm)
			.then(response => {
				this.jsonData =response.data;
			}).catch((e) => {
				alert(e);
			});

    },
  },
}
</script>
<style>
  .userPage{
		background-color: #FD0;
		padding: 1rem 1rem 2rem;
		display: flex;
		flex-direction: column;
		justify-content: center;
	}
  .loginContainer{
  left: 0;
  top: 20vh;
	padding: 1rem 1rem 2rem;
  display: flex;
  justify-content: center;
  background-color: #FD7;
  }
  .userPage .bt{
    padding: 0.25rem 1rem;
    background-color: #FB0;
    border-radius: 0.25rem;
    border-color: #EA0;
    border-width: 3px;
    border-style: solid;
    font-weight: bold;
    font-size: 1rem;
    margin-top: 1rem;
  }
  .userPage .bt:hover{
    background-color: #EA0;
    border-color: #D80;
  }
</style>
