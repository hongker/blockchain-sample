<template>
	<section>
		<el-menu
  :default-active="activeIndex2"
  class="el-menu-demo"
  mode="horizontal"
  background-color="#545c64"
  text-color="#fff"
  active-text-color="#ffd04b">
  <el-menu-item index="1">竞拍列表</el-menu-item>
  <el-submenu index="2">
    <template slot="title">我的主页</template>
    <el-menu-item index="2-1">已参与竞拍</el-menu-item>
    <el-menu-item index="2-2">已拍下商品</el-menu-item>
    <el-menu-item index="2-3">已售出商品</el-menu-item>
  </el-submenu>
  <el-menu-item index="3"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
</el-menu>
	<section class="container">
		<el-row>
  <el-col :span="4" v-for="(item, index) in goods" :key="index" :offset="index > 0 ? 2 : 0">
    <el-card :body-style="{ padding: '0px' }">
      <img v-bind:src="item.picture" class="image">
      <div style="padding: 14px;">
        <span>{{item.title}}</span>
        <span class="price"><strong class="red">{{item.price}}</strong></span>
       
        <div class="bottom clearfix">
        	
        	<el-button v-if="user.id==item.userId && item.state==0" type="danger" class="button" size="mini" @click="up(index)">上架</el-button>

        	<el-button v-if="user.id==item.userId && item.state==1" type="danger" class="button" size="mini" @click="down(index)">下架</el-button>
          <el-button v-if="user.id==item.userId && item.state==1" type="info" class="button" size="mini" @click="start(index)">开始</el-button>

          	<el-button v-if="user.id==item.userId && item.state==2" type="danger" class="button" size="mini" @click="end(index)">结束</el-button>
          <el-button v-if="item.state==2" type="danger" class="button" size="mini" @click="auction(index)">出价</el-button>

          	<p v-if="item.state==3">已结束</p>
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
				user: {
					id:10001,
				},
				goods: [
					{
						id: 1,
						picture: "http://element-cn.eleme.io/static/hamburger.50e4091.png",
						title: "好吃的汉堡",
						date:"02-08 14:50:00",
						price: 200.00,
						userId: 10001,
						state: 0,
					},
					{
						id: 2,
						picture: "http://element-cn.eleme.io/static/hamburger.50e4091.png",
						title: "好吃的汉堡",
						date:"02-08 14:50:00",
						price: 200.00,
						userId: 10002,
						state: 0,
					}
				],
			}
		},
		methods: {
			up(index) {
				this.$message({
		            type: 'success',
		            message: '上架成功'
		          }); 
				this.goods[index].state = 1;
        		
			},
			start(index) {
				this.$message({
		            type: 'success',
		            message: '启动成功'
		          }); 
				this.goods[index].state = 2;
        		
			},
			down(index) {
				this.$message({
		            type: 'success',
		            message: '下架成功'
		          }); 
				this.goods[index].state = 0;
			},
			end(index) {
				this.$message({
		            type: 'success',
		            message: '结束成功'
		          }); 
				this.goods[index].state = 3;
			},
      auction(index) {
        this.$prompt('请输入你的价格', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
        }).then(({ value }) => {
        	let oldPrice = this.goods[index].price;
        	if(value <= oldPrice) {
        		this.$message({
		            type: 'error',
		            message: '不能低于当前价:'+oldPrice
		          }); 
        		return;
        	}
          this.$message({
            type: 'success',
            message: '你的出价是: ' + value
          });
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '取消输入'
          });       
        });
      }
    }
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
    margin-left: 5px;
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
