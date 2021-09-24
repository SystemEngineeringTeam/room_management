<template>
    <div class="registContainer">
      <div class="regist">
        <div>
          <div class="backBt" v-on:click="hideRegist"></div>
        </div>
        <form @submit.prevent="regist">
          <table>
            <thead>
              <tr>
                <th>入室時間</th>
                <th>{{time}}</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <th>
                  <label for="name">名前</label>
                </th>
                <th>
                  <input type="text" id="name" v-model="frm.Name" required>
                </th>
              </tr>
              <tr>
                <th>
                  <label for="pass">パスワード</label>
                </th>
                <th>
                  <input type="password" v-model="frm.Password" required>
                </th>
              </tr>
              <tr>
                <th>
                  <label for="strNum">学籍番号</label>
                </th>
                <th>
                  <input type="text" id="strNum" pattern="[a-z]\d{5}" :value="frm.Studentnumber" @input="frm.Studentnumber = $event.target.value;onChangeStuNum();" maxlength="6" required>
                </th>
              </tr>
              <tr>
                <th>
                  <label for="Email">メアド</label>
                </th>
                <th>
                  <input type="email" id="email" v-model="frm.Email" maxlength='45' required readonly>
                </th>
              </tr>
            </tbody>
          </table>
          <input class="bt" type="submit" value="登録">
        </form>
      </div>
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
      Password:'hoge',
      Studentnumber:'',
      Email:'example@aitech.ac.jp',
      UID:''
    }
  }),
  methods:{
    regist(){
      
      // jsonを送信する
      var reqPost = new XMLHttpRequest();
      reqPost.open('POST',this.host+'/user',true);// apiに送るリクエストのURL指定
      reqPost.setRequestHeader('Content-Type','application/json');
      reqPost.send(JSON.stringify(this.frm));
      location.reload(true);
    },
    onChangeStuNum(){
      var gstr=this.frm.Studentnumber.slice(0,1);
      if(gstr.match("[A-Z]")!=null){
        gstr=gstr.toLowerCase();
        this.frm.Studentnumber=gstr;
      }
      if(gstr==''||gstr.match("[a-z]")==null){
      this.frm.Studentnumber='';
      this.frm.Email="example@aitech.ac.jp";
      }else{
        if (this.frm.Studentnumber.length>1) {
          if(this.frm.Studentnumber.slice(-1).match("[0-9]")==null){
            this.frm.Studentnumber=this.frm.Studentnumber.slice(0,-1);
          }
        }
      this.frm.Email=this.frm.Studentnumber +gstr+gstr+"@aitech.ac.jp";
      }
    },
    hideRegist(){
      this.frm.UID=null;
      this.$parent.uid=null;
    }
  },
  mounted(){
    this.frm.UID=this.$parent.uid;
    this.time = this.$parent.time;
  },
}
</script>
<style>
  .registContainer{
    left: 0;
    top: 20vh;
    width: 100%;
    position: fixed;
    display: flex;
    justify-content: center;
  }
  .regist{
    box-shadow: 5px 10px 10px #0BB;
    background-color: #9EE;
    border-radius: 1rem;
    padding: 0.75rem 1rem 1.5rem;
    display: flex;
    justify-content: center;
    flex-direction: column;
    z-index: 2;
  }
  .regist .bt{
    color: #EFF;
    padding: 0.25rem 1rem;
    background-color: #2BF;
    border-radius: 0.25rem;
    border-color: #0AD;
    border-width: 3px;
    border-style: solid;
    font-weight: bold;
    font-size: 1rem;
    margin-top: 1rem;
  }
  .regist .bt:hover{
    background-color: #0BD;
    border-color: #09D;
  }
  .backBt{
    height: 1rem;
    width: 1rem;
    border-radius: 50%;
    background-color: #5CC;
    margin-left: auto;
    margin-bottom: 0.5rem;
  }
</style>
