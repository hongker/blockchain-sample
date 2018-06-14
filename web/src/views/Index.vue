<template>
  
    <div class="wrapper">

      <div class="bg"></div>
      <nav class="navbar navbar-expand-lg navbar-light bg-light">
  <a class="navbar-brand" href="#"><i class="el-icon-menu"></i></a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>

  <div class="collapse navbar-collapse" id="navbarSupportedContent">
    <ul class="navbar-nav mr-auto">
      <li class="nav-item active">
        <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
      </li>
      
      <li class="nav-item dropdown">
        <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
          产品
        </a>
        <div class="dropdown-menu" aria-labelledby="navbarDropdown">
          <a class="dropdown-item" href="#">Product One</a>
          <a class="dropdown-item" href="#">Product Two</a>
          <a class="dropdown-item" href="#">Product Three</a>
        </div>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">博客</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">关于我们</a>
      </li>

      
    </ul>
    <ul class="navbar-nav navbar-right">
        <li class="nav-item">
        <a class="nav-link" href="#">登录</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">注册</a>
      </li>
      </ul>
    
  </div>
</nav>
     

      <div class="content">  
          <el-carousel :interval="5000" arrow="never" indicator-position="none">
            <el-carousel-item  :key="1">
            <p class="title">区块链是什么</p>
        <p class="desc">区块链是分布式数据存储、点对点传输、共识机制、加密算法等计算机技术的新型应用模式。</p>
            </el-carousel-item>
            <el-carousel-item  :key="2">
            <p class="title">区块链的特点</p>
        <p class="desc">数据去中心化、系统开放性、信息不可篡改、节点对等。</p>
            </el-carousel-item>
            <el-carousel-item  :key="3">
            <p class="title">区块链的应用领域</p>
        <p class="desc">艺术行业、法律行业、开发行业、房地产行业、物联网、保险行业、金融行业等。</p>
            </el-carousel-item>
        </el-carousel>
        
        
      </div>
    </div>
    


</template>

<script>
  export default {
    data() {
      return {
        activeIndex: '1',
      };
    },
    methods: {
      handleSelect(key, keyPath) {
        console.log(key, keyPath);
      }
    }
  }
</script>
<style>
.menu {
    padding: 0px 30px;
}
.bg{
  position:fixed;
  top: 0;
  left: 0;
  width:100%;
  height:100%;
  min-width: 1000px;
  z-index:-10;
  zoom: 1;
  background-color: #fff;
  background-repeat: no-repeat;
  background-size: cover;
  -webkit-background-size: cover;
  -o-background-size: cover;
  background-position: center 0;
  background: url(../assets/slide2.jpg);
  background-size: 100% 100%;
}


.el-menu-demo {
  padding: 0px 80px;
  border-bottom: 0px;
}

.el-menu--horizontal>.right {
  float: right;
}

.content {
  height: 200px;
  width: 600px;
  margin: 12% auto;
  text-align: center;
  color: #eee;
  
}
@media screen and (max-width: 800px){
  .content { width: 300px; }
}
.content .title {
  font-size: 48px;
  
}
.content .desc {
  
  font-size: 20px;
}

.bg-light {
  background-color: #fdfdfd00!important;
}
.navbar-light .navbar-nav .active>.nav-link, .navbar-light .navbar-nav .nav-link.active, .navbar-light .navbar-nav .nav-link.show, .navbar-light .navbar-nav .show>.nav-link {
  color: #fff;
}

.navbar-light .navbar-nav .nav-link,.navbar-light .navbar-brand, .navbar-light .navbar-toggler{
    color: #f9eeee;
}


.navbar-light ul li a {
  font-size: 16px;
  font-weight: normal;
  transition: 0.5s;
}
.navbar-light .navbar-nav .nav-link:focus, .navbar-light .navbar-nav .nav-link:hover {
  color: #fff;
}

</style>


