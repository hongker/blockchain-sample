<template>
  <section class="">
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

<el-table
    :data="tableData"
    border
    align="left"
    style="width: 100%">
    <el-table-column
      prop="name"
      label="应用名称"
      width="180">
    </el-table-column>
    <el-table-column
      prop="desc"
      label="介绍"
      width="280">
    </el-table-column>
    <el-table-column
      prop="date"
      label="时间"
      width="180">
    </el-table-column>
    <el-table-column
      label="操作"
      >
       <template slot-scope="scope">
        
        <el-button v-if="step == 1" type="primary" size="small" @click="handleChannel(scope.row)">部署通道</el-button>
        <el-button v-else-if="step == 2" type="primary" size="small" @click="handlePeer(scope.row)">加入节点</el-button>
        <el-button v-else-if="step == 3" type="primary" size="small" @click="installChaincode(scope.row)">部署链码</el-button>
        <el-button v-else-if="step == 4" type="primary" size="small" @click="initChaincode(scope.row)">初始化</el-button>
        <el-button v-else @click="handleManage(scope.row)" type="primary" size="small" >管理</el-button>
        
      </template>
    </el-table-column>
  </el-table>

  </section>
</template>

<script>
   import {
     installChaincode,
     queryChaincode,
     createChannel,
     joinChannel,
     queryChannel,
     initChaincode
   } from '../api/api';
   import {
     Loading
   } from 'element-ui';
   export default {
     data() {
       return {
         step: 1,
         activeIndex: '1',
         activeIndex2: '1',
         username: "a",
         other: "b",
         peer: "peer1",
         channelName: "mychannel",
         chaincodeName: "mycc",

         tableData: [{
           date: '2016-05-02',
           name: '智能拍卖',
           desc: '基于区块链的智能拍卖'
         }],
       };
     },

     methods: {
       handleManage(row) {
         this.$router.push({
           path: '/manage'
         });
       },
       handleChannel(row) {
         let params = {
           channelName: this.channelName,
           channelConfigPath: "../fixtures/channel/mychannel.tx"
         };

         createChannel(params).then(res => {
           if (res.success) {
             this.$message({
               message: "操作成功",
               type: 'success'
             });
             this.step = 2;
           } else {
             this.$message({
               message: "操作成功",
               type: 'error'
             });
           }
         });
       },
       handlePeer(row) {
         let params = {
           channelName: this.channelName,
           peers: ["peer1", "peer2"]
         };
         joinChannel(params).then(res => {
           if (res.success) {
             this.$message({
               message: "操作成功",
               type: 'success'
             });
             this.step = 3;
           } else {
             this.$message({
               message: "操作成功",
               type: 'error'
             });
           }
         });
       },
       installChaincode(row) {
         var params = {
           peers: ["peer1", "peer2"],
           chaincodeName: this.chaincodeName,
           chaincodePath: "github.com/example_cc",
           chaincodeVersion: "v0",
         };

         installChaincode(params).then(res => {
           if (res.success) {
             this.$message({
               message: "操作成功",
               type: 'success'
             });
             this.step = 4;
           } else {
             this.$message({
               message: "操作成功",
               type: 'error'
             });
           }
         });
       },
       initChaincode(row) {
         let params = {
           channelName: this.channelName,
           chaincodeName: this.chaincodeName,
           chaincodeVersion: "v0",
           args: [this.username, "100", this.other, "200"]
         };
         let loadingInstance = Loading.service({
           text: "loading..."
         });
         initChaincode(params).then(res => {
           loadingInstance.close();
           if (res.success) {
             this.$message({
               type: 'success',
               message: '操作成功!'
             });
             this.step = 5;
           } else {
             this.$message({
               type: 'error',
               message: res.message,
             });
           }
         });
       }
     },
     mounted() {
        this.username = localStorage.getItem('username');
       let params = {
         peer: "peer1",
         type: "installed",
       };

       queryChannel({
         peer: this.peer,
       }).then(res => {


         if (res.indexOf(this.channelName) == -1) {
           return;
         } else {
           queryChaincode({
             peer: "peer1",
             type: "installed",
           }).then(res => {

             if (res.indexOf(this.chaincodeName) == -1) {
               this.step = 3;
             } else {
               queryChaincode({
                 peer: "peer1",
                 type: "instantiated",
               }).then(res => {
                 if (res.indexOf(this.chaincodeName) == -1) {
                   this.step = 4;
                 } else {
                   this.step = 5;
                 }
               });
             }
           });

         }


       });



     },
   }
</script>

<style scoped>
  
</style>