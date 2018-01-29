<template>
  <section class="login-bg gray-bg">
  <el-form  :model="ruleForm2" :rules="rules2" ref="ruleForm2" label-position="left" label-width="0px" class="demo-ruleForm login-container">
    <h4 class="title">用户登录</h4>
    <el-form-item prop="account">
      <el-input type="text" v-model="ruleForm2.account" auto-complete="off" placeholder="账号"></el-input>
    </el-form-item>
    <el-form-item prop="checkPass">
      <el-input type="password" v-model="ruleForm2.checkPass" @keyup.enter.native="handleSubmit2"  auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-checkbox v-model="checked" checked class="remember">记住密码</el-checkbox>
    <el-form-item>
      <el-button type="primary"  style="width:100%;"  @click.native.prevent="login" :loading="logining">登录</el-button>
    </el-form-item>
    <div class="login-footer">
      <a href="/forget" class="text-gray">忘记密码</a> | <a href="/reg" class="text-gray">直接注册</a>

    </div>
    
    <div class="login-footer">

    </div>
  </el-form>
  <div class="clearfix"></div>
  </section>
</template>

<script>
	 import { requestLogin } from '../api/api';
  export default {
    data() {
      return {
        logining: false,
        ruleForm2: {
          account: '',
          checkPass: ''
        },
        rules2: {
          account: [
            { required: true, message: '请输入账号', trigger: 'blur' },
          ],
          checkPass: [
            { required: true, message: '请输入密码', trigger: 'blur' },
          ]
        },
        checked: true
      };
    },
    methods: {
      handleReset2() {
        this.$refs.ruleForm2.resetFields();
      },
      test() {
        alert(123);
      },
      login(ev) {
        var _this = this;
        this.$refs.ruleForm2.validate((valid) => {
          if (valid) {
            this.logining = true;
            var loginParams = { username: this.ruleForm2.account, password: this.ruleForm2.checkPass,orgName:"org1" };
            requestLogin(loginParams).then(res => {
            	   if(res.success) {
                    this.$message({
                  message: "登录成功",
                  type: 'success'
                });
                    localStorage.setItem('username', loginParams.username);
                    localStorage.setItem('token', res.token);
                    this.$router.push({path: '/admin'});
                 }else {
                  this.$message({
                  message: "登录失败",
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
      console.log("user",user);
      if (user) {
        this.$router.push({ path: '/' });
      }
      this.ruleForm2.account = localStorage.getItem('name');
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