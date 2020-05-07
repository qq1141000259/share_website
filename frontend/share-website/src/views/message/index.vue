<template>
  <div>
    <el-form :model="numberValidateForm" status-icon ref="numberValidateForm" label-width="100px" class="demo-ruleForm">
      <div style="text-align:center">
      <el-input type="textarea" v-model="numberValidateForm.textarea" autocomplete="off"></el-input>
      <el-button type="primary" @click="submitForm('numberValidateForm')">提交</el-button>
      </div>
      <el-form-item
        prop="textarea"
        :rules="[
          { required: true, message: '内容不能为空', trigger: 'blur'}
        ]"
      >
      </el-form-item>
    </el-form>
    <el-table :data="tableData">
      <el-table-column prop="date" label="日期" width="180"></el-table-column>
      <el-table-column prop="user" label="姓名" width="180"></el-table-column>
      <el-table-column prop="content" label="内容" width="280"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
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
      tableData: [],
      numberValidateForm: {
          textarea: ''
      }
    };
  },
  methods: {
    handleDelete(index, row) {
      pictureApi.deleteMsgRow(row.id).then(response => {
        if (!response.code) {
          this.fetchData()
        }
      });
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
      pictureApi.getMsgList(this.currentPage, this.pageSize).then(response => {
        if (!response.code) {
          const resp = response.data;
          // console.log(resp.data.rows)
          this.tableData = resp.data.rows;
          this.total = resp.data.total;
        }
      });
    },
    submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            pictureApi.postMsgRow(this.numberValidateForm.textarea).then(response => {
              this.fetchData()
            });
            // alert(this.numberValidateForm.textarea)
          } else {
            console.log('error submit!!');
            return false;
          }
        });
    },
  },
  mounted: function() {
    pictureApi.getMsgList(1, 9).then(response => {
      if (!response.code) {
        const resp = response.data;
        // console.log(resp.data.rows)
        this.tableData = resp.data.rows;
        this.total = resp.data.total;
      }
    });
  }
};
</script>
