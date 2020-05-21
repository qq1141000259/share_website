<template>
    <div class="navbar">

        <!--  default-active :　默认选中的菜单
            :router="true" true表示开启路由模式,开启之后, index值代表的就是路由地址
          -->
     <el-menu
      :router="true"
      :default-active="$route.path"
      class="el-menu-vertical-demo"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b">
      <el-menu-item index="/home">
        <i class="el-icon-s-home"></i>
        <span slot="title">首页</span>
      </el-menu-item>
      <el-menu-item index="/word">
        <i class="el-icon-notebook-1"></i>
        <span slot="title">单词薄</span>
      </el-menu-item>
      <el-menu-item index="/picture">
        <i class="el-icon-picture"></i>
        <span slot="title">图片</span>
      </el-menu-item>
      <el-menu-item index="/video">
        <i class="el-icon-video-camera-solid"></i>
        <span slot="title">视频</span>
      </el-menu-item>
      <el-menu-item index="/message">
        <i class="el-icon-notebook-2"></i>
        <span slot="title">留言薄</span>
      </el-menu-item>
      <el-menu-item index="/upload">
        <i class="el-icon-upload2"></i>
        <span slot="title">上传</span>
      </el-menu-item>
    </el-menu>

    </div>
</template>


<style scoped>
.el-menu {
    border-right: none;
}
</style>