<template>
		<div class="room">
			<h1>入室者一覧</h1>
			<table>
				<thead>
					<tr>
						<th>入室時間</th>
						<th>名前</th>
						<th>学籍番号</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="person in jsonData" :key="person.EntryDatetime">
						<th>{{person.EntryDatetime}}</th>
						<th>{{person.Name}}</th>
						<th>{{person.StudentNumber}}</th>
					</tr>
				</tbody>
			</table>
		</div>
</template>

<script>
import axios from 'axios';

export default{
	data: ()=>({
		jsonData: null,
	}),
	mounted (){
		axios.get(this.$parent.host+'/user')
			.then(response => {
				this.jsonData =response.data;
			}).catch((e) => {
				alert(e);
			});
	}
}
</script>
<style>
	.room{
		background-color: #CCC;
		padding: 1rem 1rem 2rem;
		display: flex;
		flex-direction: column;
		justify-content: center;
	}
</style>