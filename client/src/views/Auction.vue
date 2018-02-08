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
	<section class="container">
		<el-row>
  <el-col :span="4" v-for="(o, index) in 2" :key="o" :offset="index > 0 ? 2 : 0">
    <el-card :body-style="{ padding: '0px' }">
      <img src="http://element-cn.eleme.io/static/hamburger.50e4091.png" class="image">
      <div style="padding: 14px;">
        <span>好吃的汉堡</span>
        <span class="price"><strong class="red">200</strong></span>
        <div class="bottom clearfix">
          <time class="time">{{ currentDate }}</time>
          <el-button type="danger" class="button" size="medium">出价</el-button>
        </div>
      </div>
    </el-card>
  </el-col>
</el-row>
	</section>
	
	</section>
</template>
<script>
	export default {
		data() {
			return {
				activeIndex2: '1',
				currentDate: '02-08 14:50:00',
			}
		},
	}
</script>
<style>
  .container {
  	padding: 20px;
  }
  .time {
    font-size: 13px;
    color: #999;
  }
  
  .bottom {
    margin-top: 13px;
    line-height: 12px;
  }

  .button {
    float: right;
  }

  .image {
    width: 100%;
    display: block;
  }

  .clearfix:before,
  .clearfix:after {
      display: table;
      content: "";
  }
  
  .clearfix:after {
      clear: both
  }
  .red {
  	color:#F56C6C;
  }
  .price {
  	margin-left: 20px;
  }
</style>
