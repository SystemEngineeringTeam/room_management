<template>
    <div class="regist">
      <form @submit.prevent="regist">
        <table>
          <thead></thead>
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
                <input type="text" v-model="frm.Password" required readonly>
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
      {{frm}}
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
      Email:'',
      UID:''
    }
  }),
  methods:{
    regist(){
      
      // jsonを送信する
      var reqPost = new XMLHttpRequest();
      reqPost.open('POST',this.$parent.host+'/user',true);// apiに送るリクエストのURL指定
      reqPost.setRequestHeader('Content-Type','application/json');
      reqPost.send(JSON.stringify(this.frm));
      location.reload(true);
    },
    onChangeStuNum(){
      var gstr=this.frm.Studentnumber.slice(0,1);
      this.frm.Email=this.frm.Studentnumber +gstr+gstr+"@aitech.ac.jp";
    }
  },
  mounted(){
    this.frm.UID=this.$parent.uid;
    this.time = this.$parent.time;
  },
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
</style>