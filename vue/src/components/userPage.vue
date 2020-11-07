<template>
    <div class="userPage">
      <h1>ユーザー情報</h1>
      <div class="loginContainer" v-if="!logined">
        <form @submit.prevent="loginSubmit">
          <h2>ログイン</h2>
          <h3 v-if="errorMessage!==''">{{errorMessage}}</h3>
          <table>
            <thead></thead>
            <tbody>
              <tr>
                <th>
                  <label for="strNum">学籍番号</label>
                </th>
                <th>
                  <input type="text" id="strNum" pattern="[a-z]\d{5}" :value="loginData.Studentnumber" @input="loginData.Studentnumber = $event.target.value;onChangeStuNum();" maxlength="6" required>
                  <input type="hidden" id="email" v-model="loginData.frm.Email" maxlength='45' required readonly>
                </th>
              </tr>
              <tr>
                <th>
                  <label for="pass">パスワード</label>
                </th>
                <th>
                  <input type="text" v-model="loginData.frm.Password" required readonly>
                </th>
              </tr>
            </tbody>
          </table>
          <input class="bt" type="submit" value="ログインする">
        </form>
      </div>
      <div v-else class="userPageinner">
        <table>
          <thead></thead>
          <tbody>
            <tr>
              <th>名前</th>
              <th>{{userData.Name}}</th>
            </tr>
            <tr>
              <th>学籍番号</th>
              <th>{{userData.StudentNumber}}</th>
            </tr>
            <tr>
              <th>メアド</th>
              <th>{{userData.Email}}</th>
            </tr>
          </tbody>
        </table>
        <Week></Week>
      </div>
    </div>
</template>

<script>
import Week from './week.vue';
import axios from 'axios';

export default {
  name: 'User',
  components:{
    Week,
  },
  data: ()=>({
    host:null,
    jsonData: null,
    userData:null,
    logined:false,
    loginData:{
      Studentnumber:'',
      frm:{
        Email:'',
        Password:'hoge',
      },
    },
    errorMessage:'',
  }),
  methods:{
    onChangeStuNum(){
      var gstr=this.loginData.Studentnumber.slice(0,1);
      if(gstr.match("[A-Z]")!=null){
        gstr=gstr.toLowerCase();
        this.loginData.Studentnumber=gstr;
      }
      if(gstr==''||gstr.match("[a-z]")==null){
      this.loginData.Studentnumber='';
      this.loginData.frm.Email="example@aitech.ac.jp";
      }else{
        if (this.loginData.Studentnumber.length>1) {
          if(this.loginData.Studentnumber.slice(-1).match("[0-9]")==null){
            this.loginData.Studentnumber=this.loginData.Studentnumber.slice(0,-1);
          }
        }
      this.loginData.frm.Email=this.loginData.Studentnumber +gstr+gstr+"@aitech.ac.jp";
      }
    },
    loginSubmit(){
      axios.post(this.$parent.host+'/login',this.loginData.frm)
			.then(response => {
        this.userData =response.data;
        if(this.userData.Name!==''){
          this.logined=true;
        }else{
          this.errorMessage="ログインに失敗しました";
        }
			}).catch((e) => {
				alert(e);
      });

    },
  },
  mounted(){
    this.host = this.$parent.host;
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
  .userPageinner{
    display: flex;
		flex-direction: column;
		justify-content: center;
  }
</style>
