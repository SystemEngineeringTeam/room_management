<template>
    <div id="analytics">
      <h1>入退室グラフ</h1>
      <div>
        <input type="date" :value="startTex.date" @input="startDate = inputDate(startDate,$event.target.value);graphData=genGraphData(jsonData,startDate,endDate);dates2Tex();">
        <input type="time" :value="startTex.time" @input="startDate = inputTime(startDate,$event.target.value);graphData=genGraphData(jsonData,startDate,endDate);dates2Tex();">
        ～
        <input type="date" :value="endTex.date" @input="endDate = inputDate(endDate,$event.target.value);graphData=genGraphData(jsonData,startDate,endDate);dates2Tex();">
        <input type="time" :value="endTex.time" @input="endDate = inputTime(endDate,$event.target.value);graphData=genGraphData(jsonData,startDate,endDate);dates2Tex();">
        
        <!-- <div v-for="(log,itr) in jsonData" :key="itr">
          {{itr}}:{{log}}
        </div> -->
        <h3 v-if="Object.keys(graphData).length==0">この期間の入室はありません</h3>
        <div class="graphContainer" v-for="(datum,stuNum) in graphData" :key="stuNum">
          <span>
            <small>{{stuNum}}</small><br>
            {{datum.name}}
          </span>
          <!-- {{datum}} -->
          <div class="graphBar" :style="datum.style">
            <div v-for="n of datum.cnt" :key="n">
            </div>
          </div>
        </div>
      </div>
      
    </div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'analytics',
  data: ()=>({
    jsonData: null,
    graphData:{},
    startDate:null,
    startTex:{
      date:null,
      time:null
    },
    endTex:{
      date:null,
      time:null
    },
    endDate:null,
    hoge:null,
  }),
  methods:{
    initDate(){
      let defTime=new Date();
      defTime.setHours(0);
      defTime.setMinutes(0);
      defTime.setSeconds(0);
      defTime.setMilliseconds(0);
      this.startDate=defTime;
      this.endDate=new Date(defTime);
      this.endDate.setDate(defTime.getDate()+1);
      this.dates2Tex();
    },
    inputDate(date,val){
      let DateArr = val.match(/(\d+)/g).map(str => parseInt(str, 10));
      date.setYear(DateArr[0]);
      date.setMonth(DateArr[1]-1);
      date.setDate(DateArr[2]);
      return date;
    },
    inputTime(date,val){
      let timeArr = val.match(/(\d+)/g).map(str => parseInt(str, 10));
      date.setHours(timeArr[0]);
      date.setMinutes(timeArr[1]);
      return date;
    },
    dates2Tex(){
      let funcDate2Tex = ((e)=>{
        return e.getFullYear()+"-"+('00'+(e.getMonth()+1)).slice(-2)+"-"+('00'+e.getDate()).slice(-2);
      });
      let funcTime2Tex = ((e)=>{
        return ('00'+e.getHours()).slice(-2)+":"+('00'+e.getMinutes()).slice(-2);
      });
      this.startTex.date=funcDate2Tex(this.startDate);
      this.startTex.time=funcTime2Tex(this.startDate);
      this.endTex.date=funcDate2Tex(this.endDate);
      this.endTex.time=funcTime2Tex(this.endDate);
    },
    genGraphData(d,start,end){
      if(start.getTime()>=end.getTime()){
        return {};
      }
      let result= new Map();
      let befEnterMap = new Map();
      let nameMap = new Map();
      let timeRange = end.getTime()/1000-start.getTime()/1000;
      d.forEach(log => {
        let logDateArr = log.CardReadDatetime.match(/(\d+)/g).map(str => parseInt(str, 10));
        let logTime = new Date(logDateArr[0],logDateArr[1]-1,logDateArr[2],logDateArr[3],logDateArr[4],logDateArr[5]);
        // console.log(logTime);
        let deltaTime = logTime.getTime()/1000-start.getTime()/1000;
        let timePoint =deltaTime/timeRange;
        if(timePoint<=1&&timePoint>=0){
          if(result.has(log.StudentNumber)){
            let object = result.get(log.StudentNumber);
            // object.array[object.array.length]=[deltaTime,timePoint,(end-logTime)/1000];
            if(object.cnt%2!=log.EntryOrExit){
              object.style['grid-template-columns']+=(timePoint-object.fr)+'fr ';
              object.cnt++;
              object.fr=timePoint;
            }else{
              // object.style['grid-template-columns']+=log.EntryOrExit+'px ';
            }
            result.set(log.StudentNumber,object);
          }else{
            // result.set(log.StudentNumber,{name:log.Name,array:[[logTime,deltaTime,timePoint,(end-logTime)/1000]],test:[timeRange,start,end],style:{'grid-template-columns':(log.EntryOrExit==1?'':'0fr ')+timePoint+'fr '},cnt:2-log.EntryOrExit,fr:timePoint});
            result.set(log.StudentNumber,{name:log.Name,style:{'grid-template-columns':(log.EntryOrExit==1?'':'0fr ')+timePoint+'fr '},cnt:2-log.EntryOrExit,fr:timePoint});
          }
        }else if(timePoint<0){
          befEnterMap.set(log.StudentNumber,log.EntryOrExit);
        }
        if(!nameMap.has(log.StudentNumber)){
          nameMap.set(log.StudentNumber,log.Name);
        }
      });
      befEnterMap.forEach((val,key) => {
        if(val==1&&!result.has(key)){
          result.set(key,{name:nameMap.get(key),style:{'grid-template-columns':'0fr '},cnt:1,fr:0});
        }
      });
      result.forEach(e => {
        e.style['grid-template-columns']+=(1-e.fr)+'fr';
        e.cnt++;
      });
      return Object.fromEntries(result);
    },

  },
  mounted (){
    this.initDate();
		axios.get(this.host+'/log')
			.then(response => {
				this.jsonData =response.data;
        this.jsonData.sort((a,b)=>{return (new Date(a.CardReadDatetime)-new Date(b.CardReadDatetime))});
        this.graphData=this.genGraphData(this.jsonData,this.startDate,this.endDate);
			}).catch((e) => {
				alert(e);
			});
	}
}
</script>
<style>
  #analytics{
		background-color: #6DD;
    padding-bottom:50px;
	}
  .graphBar{
    display:grid;
    height:1rem;
    width:calc(100% - 150px);
  }
  .graphBar>*:nth-child(2n){
    background-color:red;
  }
  .graphBar>*:nth-child(2n-1){
    /* background-color:blue; */
  }
  .graphContainer{
    display:flex;
    flex-direction:column;
    align-items:center;
  }
</style>