<template>
    <div class="resetDays">
      <table>
        <thead>
          <tr>
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
            <th v-for="aday in weeks" :key="aday">
              <div v-if="aday==null">
                リセットしない
              </div>
              <div v-else>
                リセットする
              </div>
            </th>
          </tr>
        </tbody>
      </table>
      <div>
        {{jsonData}}
        <button v-on:click="hoge">日曜リセしてみる</button>
        {{res}}
      </div>

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
    weeksName:['Sun','Mon','Tue','Wed','Thu','Fri','Sat'],
    weeks:[null,null,null,null,null,null,null],
    res:null,
  }),
  methods:{
    hoge(){
      let postData={
        Email:this.Email,
        Day:'Sun',
        IsOnce:true,
      };
      axios.put(this.$parent.host+'/reset',postData)
			.then(response => {
        this.res =response.data;
			}).catch((e) => {
				alert(e);
      });
    },
  },
  mounted(){
    this.Email = this.$parent.userData.Email;
    this.host = this.$parent.host;
    axios.get(this.host+'/reset',{Email:this.Email})
			.then(response => {
				this.jsonData =response.data;
			}).catch((e) => {
				alert(e);
      });
    
    this.jsonData.forEach(aday => {
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
  }
</style>
