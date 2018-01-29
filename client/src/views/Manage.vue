<template>
	<section>
		<el-menu
  :default-active="activeIndex2"
  class="el-menu-demo"
  mode="horizontal"
  background-color="#545c64"
  text-color="#fff"
  active-text-color="#ffd04b">
  <el-menu-item index="1">处理中心</el-menu-item>
  <el-submenu index="2">
    <template slot="title">我的工作台</template>
    <el-menu-item index="2-1">选项1</el-menu-item>
    <el-menu-item index="2-2">选项2</el-menu-item>
    <el-menu-item index="2-3">选项3</el-menu-item>
  </el-submenu>
  <el-menu-item index="3"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
</el-menu>

<el-form label-width="80px">
  <el-form-item label="账户名称">
    <div>{{username}}</div>
  </el-form-item>
  <el-form-item label="账户余额">
    <div >{{amount}}</div>
  </el-form-item>

  
  
</el-form>

<el-form :inline="true" :model="payment">
	<el-form-item label="收款人">
  	<el-input v-model="payment.receiver" placeholder="请输入内容"></el-input>
  </el-form-item>
  <el-form-item label="金额">
  	
  	<el-input v-model="payment.amount" placeholder="请输入内容"></el-input>
  </el-form-item>
  <el-form-item>
    <el-button type="primary" @click="pay">付款</el-button>
  </el-form-item>
</el-form>



	</section>

</template>
<script>
import {
	 query, invokeChaincode
} from '../api/api';
export default {
	data() {
		return {
			username: "a",
			amount: "0.00",
			activeIndex2: '1',
			other:"b",
			payment: {
				receiver: "b",
				amount: "1",
			}
		}
	},
	methods: {
		pay() {
			let params = {
				channelName: "mychannel",
				chaincodeName: "mycc",
				fcn: "move",
				args: [this.username,this.payment.receiver,this.payment.amount]
			}

			invokeChaincode(params).then(res => {
				console.log(res);
			});


		},
	
		
	},
	mounted() {
		this.username = localStorage.getItem('username');
		let params = {
			channelName: "mychannel",
			chaincodeName: "mycc",
			peer: "peer1",
			fcn: "query",
			args: "['"+this.username+"']",
		};

		query(params).then(res => {
			this.amount = res.amount;
		});
	}
}
</script>