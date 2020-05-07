<template>
  <div>
    <el-carousel :interval="4000" type="card" style="width:100%;height:100%">
      <el-carousel-item v-for="item in urls" :key="item">
         <el-image class="medium" :src="item" style="width:100%;height:100%"></el-image>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>
<script>
import pictureApi from "@/api/picture";
  export default {
    data() {
        return {
          urls: []
        };
      },
    mounted: function() {
      pictureApi.getBackPicList().then(response => {
        if(!response.code) {
            const resp = response.data;
            // console.log(resp.data.rows)
            this.urls = resp.data;
          } 
      });
    }
    }   
</script>      

<style>
  .el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 500px;
    margin: 0;
  }
  
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
</style>