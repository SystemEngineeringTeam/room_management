<template>
    <div class="resetDays">
      <h2>リセットする日を指定する</h2>
      <table>
        <thead>
          <tr>
            <th></th>
            <th>日</th>
            <th>月</th>
            <th>火</th>
            <th>水</th>
            <th>木</th>
            <th>金</th>
            <th>土</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <th>リセット</th>
            <td v-for="(aday,index) in weeks" :key="aday">
              <div v-if="aday==null" class="resetContainer">
                <input type="checkbox" v-on:click="resetOn(index)">
              </div>
              <div v-else class="resetContainer">
                <input type="checkbox" v-on:click="resetOff(index,aday)" checked>
                <span v-if="aday" v-on:click="changeIsOnce(index,false)">一度切り</span>
                <span v-else v-on:click="changeIsOnce(index,true)">毎週</span>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div style="display:none;">{{jsonData}}</div>
    </div>
</template>

<script>

import axios from "axios";

export default {
  name: 'week',
  data: ()=>({
    host:null,
    Email:null,
    jsonData: null,
    reqestData:{
      Email:null,
      Day:null,
      IsOnce:true,
    },
    weeksName:['Sun','Mon','Tue','Wed','Thu','Fri','Sat'],
    weeks:[null,null,null,null,null,null,null],
    res:null,
  }),
  methods:{
    resetOn(dayNum){
      this.reqestData.Day=this.weeksName[dayNum];
      axios.post(this.$parent.host+'/reset',this.reqestData)
			.then(response => {
        this.res =response.data;
        this.loadWeeks();
			}).catch((e) => {
				alert(e);
      });
    },
    resetOff(dayNum,IsOnce){
      this.reqestData.Day=this.weeksName[dayNum];
      this.reqestData.IsOnce=IsOnce;
      axios.put(this.$parent.host+'/reset',this.reqestData)
			.then(response => {
        this.res =response.data;
        this.loadWeeks();
			}).catch((e) => {
				alert(e);
      });
    },
    changeIsOnce(dayNum,IsOnce){
      this.reqestData.Day=this.weeksName[dayNum];
      this.reqestData.IsOnce=!IsOnce;
      axios.put(this.$parent.host+'/reset',this.reqestData)
			.then(response => {
        this.res =response.data;
        this.reqestData.IsOnce=IsOnce;
        axios.post(this.$parent.host+'/reset',this.reqestData)
        .then(response => {
          this.res =response.data;
          this.loadWeeks();
        }).catch((e) => {
          alert(e);
        });
			}).catch((e) => {
				alert(e);
      });
    },
    loadWeeks(){
      this.$parent.reloadWeekKey++;
    },
  },
  mounted(){
    this.Email = this.$parent.userData.Email;
    this.host = this.$parent.host;
    this.reqestData.Email=this.Email;
    axios.get(this.host+'/reset?Email='+this.Email)
			.then(response => {
				this.jsonData =response.data;
        for(let index in this.jsonData){
          let aday = this.jsonData[index];
        switch (aday.Day) {
          case 'Sun':
            this.weeks[0]=aday.IsOnce;
            break;
          case 'Mon':
            this.weeks[1]=aday.IsOnce;
            break;
          case 'Tue':
            this.weeks[2]=aday.IsOnce;
            break;
          case 'Wed':
            this.weeks[3]=aday.IsOnce;
            break;
          case 'Thu':
            this.weeks[4]=aday.IsOnce;
            break;
          case 'Fri':
            this.weeks[5]=aday.IsOnce;
            break;
          case 'Sat':
            this.weeks[6]=aday.IsOnce;
            break;
          default:
            break;
        }
      }
			}).catch((e) => {
				alert(e);
      });
    
  },
}
</script>
<style>
  .resetDays{
  left: 0;
  top: 20vh;
	padding: 1rem 1rem 2rem;
  display: flex;
  justify-content: center;
  background-color: #FD7;
  flex-direction:column;
  }
  .resetContainer{
    align-items:center;
    display: flex;
    flex-direction:column;
  }
  .resetContainer span{
    font-weight:bold;
  }
  .resetContainer span:hover{
    color:#7DF;
  }
</style>
