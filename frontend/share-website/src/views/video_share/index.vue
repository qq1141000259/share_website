<template>
  <div class="demo-image">
    <div class="new">
    <div v-for="url in urls" :key="url.url" style="text-align: center;display:inline-block;width: 100%;
  height: 100%;">
      <span class="demonstration" style="display:block">{{url.title}}</span>
      <video class="videos video-js vjs-default-skin vjs-big-play-centered" preload="auto" controls>
        <source :src="url.url" type="video/mp4" >
      </video> 
    </div>
    </div>
    <!-- 分页组件 -->
    <div class="block">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :page-count="5"
        :current-page="currentPage"
        :page-sizes="[3, 9, 18]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
      ></el-pagination>
    </div>
  </div>
</template>

<script>
import pictureApi from "@/api/picture";

export default {
  data() {
    return {
      total: 20,
      currentPage: 1,
      pageSize: 3,
      urls: []
    };
  },
  methods: {
    // 当每页显示条数改变后,被触发 , val是最新的每页显示条数
    handleSizeChange(val) {
      // console.log(val)
      this.pageSize = val;
      this.fetchData();
    },
    // 当页码改变后,被触发 , val 是最新的页面
    handleCurrentChange(val) {
      // console.log(val)
      this.currentPage = val;
      this.fetchData();
    },
    fetchData() {
      // 调用分页查询数据
      pictureApi.getVideoList(this.currentPage, this.pageSize).then(response => {
        if(!response.code) {
          const resp = response.data;
          // console.log(resp.data.rows)
          this.urls = resp.data.rows;
          this.total = resp.data.total;
          this.$refs.videoPlay.src = rthis.$refs.videoPlay.a;
        }
      });
    }
  },
  mounted: function() {
    pictureApi.getVideoList(1, 3).then(response => {
      if(!response.code) {
          const resp = response.data;
          // console.log(resp.data.rows)
          this.urls = resp.data.rows;
          this.total = resp.data.total;
          this.$refs.video.src
        } 
    });
  }
};
</script>

<style scoped>
.videos {
  text-align: center;
  width: 730px;
  height: 420px;
  display: inline-block;
  padding: 1%;
}
.demonstration {
  margin: 2%;
}
</style>>




</style>