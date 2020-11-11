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
            <!-- <th v-for="aday in jsonData" :key="aday.Day">
              <span v-if="aday.IsOnce">true</span>
              <span v-else>false</span>
            </th> -->
          </tr>
        </tbody>
      </table>
      {{jsonData}}
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
  }),
  methods:{
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
