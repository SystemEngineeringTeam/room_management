<template>
    <div class="log">
      <h1>入退室ログ(ユーザー登録)</h1>
      <table>
				<thead>
					<tr>
						<th>時間</th>
            <th>入退</th>
						<th>名前</th>
						<th>学籍番号</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="log in jsonData" :key="log.CardReadDatetime">
            <th>{{log.CardReadDatetime}}</th>
            <th>
              <div v-if="log.EntryOrExit==true">入室</div>
              <div v-else>退室</div>
            </th>
            <th>
              <div v-if="log.Name===''">
                <button v-on:click="register(log.uid,log.CardReadDatetime)"><span>登録</span><i class="fas fa-user-plus"></i></button>
              </div>
              <div v-else>{{log.Name}}</div>
            </th>
						<th>{{log.StudentNumber}}</th>
					</tr>
				</tbody>
      </table>
      <Regist v-if="uid!=null"></Regist>
    </div>
</template>

<script>
import Regist from './regist.vue';
import axios from 'axios';
export default {
  name: 'log',
  components:{
    Regist,
  },
  data: ()=>({
    jsonData: null,
    uid:null,
    time:null,
  }),
  methods:{
    register(u,t){
      // window.alert(u);
      this.uid=u;
      this.time=t;
    }
  },
  mounted (){
		axios.get(this.host+'/log')
			.then(response => {
				this.jsonData =response.data;
			}).catch((e) => {
				alert(e);
			});
	}
}
</script>
<style>
  .log{
		background-color: #6DD;
	}
  .log form{
    display: flex;
		flex-direction: column;
		justify-content: center;
  }
  .log button{
    color: #EFF;
    padding: 0.25rem 1rem;
    background-color: #2BF;
    border-radius: 0.25rem;
    border-color: #0AD;
    border-width: 3px;
    border-style: solid;
    font-weight: bold;
    font-size: 1.5rem;
  }
  @media(max-width:560px){
    .log button span{
      display:none;
    }
  }
  @media(min-width:561px){
    .log button i{
      display:none;
    }
  }
  .log button:hover{
    background-color: #0BD;
    border-color: #09D;
  }
</style>