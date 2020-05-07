<template>
  <div class="demo-image">
    <div class="new">
    <div class="pics" v-for="url in urls" :key="url.url">
      <span class="demonstration" style="display:inline-block">{{url.title}}</span>
      <el-image :src="url.url" lazy></el-image>
    </div>
    </div>
    <!-- 分页组件 -->
    <div class="block">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :page-count="5"
        :current-page="currentPage"
        :page-sizes="[9, 18, 45]"
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
      pageSize: 10,
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
      pictureApi.getPicList(this.currentPage, this.pageSize).then(response => {
        if(!response.code) {
          const resp = response.data;
          // console.log(resp.data.rows)
          this.urls = resp.data.rows;
          this.total = resp.data.total;
        }
      });
    }
  },
  mounted: function() {
    pictureApi.getPicList(1, 9).then(response => {
      if(!response.code) {
          const resp = response.data;
          // console.log(resp.data.rows)
          this.urls = resp.data.rows;
          this.total = resp.data.total;
        } 
    });
  }
};
</script>

<style scoped>
.pics {
  text-align: center;
  width: 31%;
  height: 31%;
  display: inline-block;
  padding: 1%;
}
.demonstration {
  margin: 2%;
}
</style>>




</style>