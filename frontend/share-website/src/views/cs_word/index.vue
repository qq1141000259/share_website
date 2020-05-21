<template>
  <div>
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="单词">
        <el-input v-model="formInline.word" placeholder="单词"></el-input>
      </el-form-item>
      <el-form-item label="释义">
        <el-input v-model="formInline.annotation" placeholder="释义"></el-input>
      </el-form-item>
      <el-form-item label="备注">
        <el-input v-model="formInline.remark" placeholder="备注"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSearch">查询</el-button>
        <el-button type="primary" @click="onSubmit">新增</el-button>
      </el-form-item>
    </el-form>
    <el-table :data="tableData">
      <el-table-column prop="id" width="70"  align="center" label="编号"></el-table-column>
      <el-table-column prop="word" width="220" align="center" label="单词" class-name="word"></el-table-column>
      <el-table-column prop="annotation" label="释义" width="240"></el-table-column>
      <el-table-column prop="remark" label="备注" width="180" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column prop="date" label="上传日期" width="110"></el-table-column>
      <el-table-column prop="user" label="上传者" width="100"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="mini" type="info" @click="onUpdate(scope.$index, scope.row)">修改</el-button>
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
    <div>
      <el-dialog
        :center=true
        :modal=false
        title="修改"
        :visible.sync="dialogVisible"
        :before-close="closeHandler" width="30%">
            <el-form ref="form" :rules="rules" :model="form" label-width="80px">
                <el-form-item label="单词" prop="word">
                    <el-input v-model="form.word"></el-input>
                </el-form-item>
                <el-form-item label="释义" prop="annotation">
                    <el-input v-model="form.annotation"></el-input>
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input v-model="form.remark"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm('form')">确认</el-button>
                </el-form-item>
            </el-form>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import wordApi from "@/api/word";
export default {
  data() {
    return {
      total: 20,
      currentPage: 1,
      pageSize: 10,
      tableData: [],
      dialogVisible: false,
      formInline: {
        word: "",
        annotation: "",
        remark: ""
      },
      form: {
        id: 0,
        word: "",
        annotation: "",
        remark: "",
      },
      rules: {
             word: [
                  {required: true, message: '帐号不能为空', trigger: 'blur' },
             ],
             annotation: [
                 {required: true, message: '密码不能为空', trigger: 'blur' },
             ]  
         },
    };
  },
  methods: {
    handleDelete(index, row) {
      wordApi.deleteWordRow(row.id).then(response => {
        if (!response.code) {
          this.fetchData();
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
      wordApi.getWordList(this.currentPage, this.pageSize, this.formInline.word,this.formInline.annotation, this.formInline.remark).then(response => {
        if (!response.code) {
          const resp = response.data;
          // console.log(resp.data.rows)
          this.tableData = resp.data.rows;
          this.total = resp.data.total;
        }
      });
    },
    onSearch() {
      console.log(this.formInline.word)
      wordApi
        .getWordList(1,9,this.formInline.word,this.formInline.annotation, this.formInline.remark)
        .then(response => {
          const resp = response.data;
          // console.log(resp.data.rows)
          this.tableData = resp.data.rows;
          this.total = resp.data.total;
        });
    },
    onSubmit() {
      const req = {
        word: this.formInline.word,
        annotation: this.formInline.annotation,
        remark: this.formInline.remark
      };
      wordApi.postWordRow(req).then(response => {
        this.fetchData();
      });
    },
    submitForm(form){
      this.$refs[form].validate(valid => {
        if (valid) {
          const data = {
              id: this.form.id,
              word: this.form.word,
              annotation: this.form.annotation,
              remark: this.form.remark
          };
          wordApi.updateWordRow(data).then(
            response =>{
              this.fetchData();
            });
        }else{
          console.log('验证失败')
          return false
        }
        this.dialogVisible = false
      })
    },
    onUpdate(index, row) {
      this.dialogVisible = true
      this.form.word = row.word
      this.form.annotation = row.annotation
      this.form.remark = row.remark
      this.form.id = row.id
    },
    closeHandler(done) {
          this.$confirm('是否取消修改')
            .then(_ => {
              done();
            })
            .catch(_ => {});
        }
    },
  mounted: function() {
    wordApi.getWordList(1, 9, this.formInline.word,this.formInline.annotation, this.formInline.remark).then(response => {
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

<style>
.word {
  font-size: 20px;
  font-weight: bold;
}
.demo-form-inline {
  margin-top: 20px;
}
</style>
