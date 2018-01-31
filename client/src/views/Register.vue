<template>
  <section class="login-bg gray-bg">
  <el-form  :model="regForm" :rules="regRules" ref="regForm" label-position="left" label-width="0px" class="login-container">
    <h4 class="title">用户注册</h4>
    <el-form-item prop="username">
      <el-input type="text" v-model="regForm.username" auto-complete="off" placeholder="账号"></el-input>
    </el-form-item>
    <el-form-item prop="orgName">
      <el-input type="text" v-model="regForm.orgName" auto-complete="off" placeholder="组织"></el-input>
    </el-form-item>
    <el-form-item prop="password">
      <el-input type="password" v-model="regForm.password" @keyup.enter.native="reg"  auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary"  style="width:100%;"  @click.native.prevent="reg" :loading="loading">注册</el-button>
    </el-form-item>
    <div class="login-footer">
      <a href="/forget" class="text-gray">忘记密码</a> | <a href="/login" class="text-gray">直接登录</a>

    </div>
    
    <div class="login-footer">

    </div>
  </el-form>
  <div class="clearfix"></div>
  </section>
</template>

<script>
   import {
     requestRegister,
     invokeChaincode,
   } from '../api/api';
   export default {
     data() {
       return {
         loading: false,
         regForm: {
           username: '',
           orgName: '',
           password: '',
         },
         regRules: {
           username: [{
             required: true,
             message: '请输入账号',
             trigger: 'blur'
           }, ],
           password: [{
             required: true,
             message: '请输入密码',
             trigger: 'blur'
           }, ]
         },
         checked: true
       };
     },
     methods: {

       reg(ev) {
         var _this = this;
         this.$refs.regForm.validate((valid) => {
           if (valid) {
             this.loading = true;
             var params = {
               username: this.regForm.username,
               password: this.regForm.password,
               orgName: this.regForm.orgName
             };
             requestRegister(params).then(res => {
               if (res.success) {

                 this.$message({
                   message: "注册成功",
                   type: 'success'
                 });
                 localStorage.setItem('username', params.username);
                 localStorage.setItem('token', res.token);

                 invokeChaincode({
                    channelName: "mychannel",
                    chaincodeName: "mycc",
                    fcn: "initialize",
                    args: [params.username, "100"],
                  }).then(res => {
                      console.log(res);
                  });
                  
                 this.$router.push({
                   path: '/admin'
                 });
               } else {
                 this.$message({
                   message: "注册失败",
                   type: 'error'
                 });
               }



             });
           } else {
             console.log('error submit!!');
             return false;
           }
         });
       }

     },
     mounted() {
       var user = localStorage.getItem('user');
       console.log("user", user);
       if (user) {
         this.$router.push({
           path: '/'
         });
       }
     }
   }
</script>

<style scoped>
  .login-bg {
    width:100%;
    height:100%;
    background-size:cover;
  }
  .login-container {
    margin: 0px auto;
    top:20px;
    position: relative;
    width: 350px;
    padding: 35px 35px 15px 35px;
    background: #fff;
    margin-bottom: 80px;
  }
  .login-footer {text-align:center;color:#888;}
  .title {
    margin: 0px auto 20px auto;
    text-align: center;
    color: #505458;
  }
  .remember {
    margin: 0px 0px 35px 0px;
  }
  .partner {
    margin-top: 20px;
  }
  .partner ul li a {
    font-size: 20px;
  }
</style>