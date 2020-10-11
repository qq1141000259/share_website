<template>
  
  <div>
    <el-row style="position: fixed;right: 30px;z-index: 999">
      <el-button type="primary" icon="el-icon-edit" circle @click="onSearch"></el-button>
    </el-row>
    <el-row :gutter="12" class="blog-row">
      <el-col :span="24" >
        <div v-for="blog in blogData" :key="blog.id" @click="queryBlog($event, blog.id)">
          <el-card shadow="hover" class="blog-card" >
          <div>
            <span class="blog-title">标 题: </span>
            <span>{{blog.title}} </span>
          </div>
          <div>
            <span class="blog-title">日 期: </span>
            <span class="blog-text">{{blog.date}} </span>
          </div>
          <div>
            <span class="blog-title">标 签: </span>
            <el-tag effect="dark" size="mini" class="blog-tag" v-for="tag in blog.tag" :key="tag">{{tag}}</el-tag>
          </div>
          <div>
            <span class="blog-title">简 介: </span>
            <span class="blog-text">{{blog.desc}}</span>
          </div>
          <div>
            <p class="blog-auth">———— {{blog.user}}</p>
          </div>
        </el-card>
        </div>
        
        
      </el-col>
    </el-row>

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
import blogApi from "@/api/blog";
export default {
  data() {
    return {
      total: 20,
      currentPage: 1,
      pageSize: 10,
      blogData: []
    };
  },
  methods: {
    // 查看博客详情
    queryBlog(val, key){
      this.$router.push('/blog/' + key)
    },
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
      blogApi.getBlogList(this.currentPage, this.pageSize).then(response => {
        if (!response.code) {
          const resp = response.data;
          // console.log(resp.data.rows)
          this.blogData = resp.data.rows;
          this.total = resp.data.total;
        }
      });
    },
    onSearch() {
      this.$router.push('/blog/edit')
    }
  },
  mounted: function() {
    blogApi.getBlogList(1, 9).then(response => {
      if (!response.code) {
        const resp = response.data;
        // console.log(resp.data.rows)
        this.blogData = resp.data.rows;
        this.total = resp.data.total;
      }
    });
  }
};
</script>

<style>
.blog-row {
  font-size: 20px;
  font-weight: bold;
  margin-top: 10px;
  margin-bottom: 10px;
}
.blog-card {
  margin: 10px;
}
.blog-tag {
  margin: 2px;
}
.blog-title {
  font-size: 15px;
  font-weight: bolder;
  color: dimgray;
}
.blog-text {
  font-size: 12px;
}
.blog-auth {
  text-align: right;
  font-size: 12px;
  color:black
}
.demo-form-inline {
  margin-top: 20px;
}
</style>
