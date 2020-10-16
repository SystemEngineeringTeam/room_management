<template>
    <div class="log">
      {{jsonData}}
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
                <button v-on:click="register(log.uid,log.CardReadDatetime)">登録</button>
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
    host: null,
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
    this.host = this.$parent.host;
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
		padding: 1rem 1rem 2rem;
		display: flex;
		flex-direction: column;
		justify-content: center;
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
  .log button:hover{
    background-color: #0BD;
    border-color: #09D;
  }
</style>