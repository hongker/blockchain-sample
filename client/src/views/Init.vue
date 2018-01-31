<template>
	<el-button v-if="step == 1" type="primary" size="small" @click="handleChannel()">部署通道</el-button>
        <el-button v-else-if="step == 2" type="primary" size="small" @click="handlePeer()">加入节点</el-button>
        <el-button v-else-if="step == 3" type="primary" size="small" @click="installChaincode()">部署链码</el-button>
        <el-button v-else-if="step == 4" type="primary" size="small" @click="initChaincode()">初始化</el-button>
        
        <el-button v-else type="primary" size="small" @click="completed()">已完成</el-button>
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
         username: "a",
         peer: "peer1",
         channelName: "mychannel",
         chaincodeName: "mycc",
       };
     },

     methods: {
       completed() {
         this.$message({
               message: "操作成功",
               type: 'success'
             });
       },
       handleChannel() {
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
       handlePeer() {
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
       installChaincode() {
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
       initChaincode() {
         let params = {
           channelName: this.channelName,
           chaincodeName: this.chaincodeName,
           chaincodeVersion: "v0",
           args: [this.username, "100"]
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
      console.log(this.step);
    }
}
</script>