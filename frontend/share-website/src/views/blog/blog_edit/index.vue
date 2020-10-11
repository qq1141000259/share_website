<template>

  <div>
    <app-link></app-link>
    <div style="margin-top:10px">
      <el-row style="position: fixed;right: 30px;z-index: 99999">
        <el-button type="success" icon="el-icon-check" circle @click="submitForm"></el-button>
      </el-row>
      <el-form ref="formData" :model="formData" label-width="50px" size="mini" :rules="rules">
        <el-form-item label="标题">
          <el-input v-model="formData.name"></el-input>
        </el-form-item>
        <el-form-item label="标签">
          <el-tag
            :key="tag"
            v-for="tag in dynamicTags"
            closable
            :disable-transitions="false"
            @close="handleClose(tag)">
            {{tag}}
          </el-tag>
          <el-input
            class="input-new-tag"
            v-if="inputVisible"
            v-model="inputValue"
            ref="saveTagInput"
            size="small"
            @keyup.enter.native="handleInputConfirm"
            @blur="handleInputConfirm"
          >
          </el-input>
          <el-button v-else class="button-new-tag" size="small" @click="showInput">+ New Tag</el-button>
        </el-form-item>
        <el-form-item label="简介">
          <el-input type="textarea" v-model="formData.desc"></el-input>
        </el-form-item>
        <el-form-item label="正文">
          <div id="main">
            <mavon-editor style="z-index: 99" ref=md :ishljs = "true" v-model="formData.value"/>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm">立即创建</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import blogApi from "@/api/blog";
import 'mavon-editor/dist/css/index.css'
import AppLink from './Link.vue'
export default {
  components: {AppLink},
  data() {
    return {
      dynamicTags: ["技术"],
      inputVisible: false,
      inputValue: '', 
      formData: {
        name: "",
        value: "",
        desc: ""
      },
      rules: {
             name: [
                  {required: true, message: '帐号不能为空', trigger: 'blur' },
             ],
             value: [
                 {required: true, message: '密码不能为空', trigger: 'blur' },
             ]  
         },
    };
  },
  methods: {
    handleClose(tag) {
      this.dynamicTags.splice(this.dynamicTags.indexOf(tag), 1);
    },

    showInput() {
      this.inputVisible = true;
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },

    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        this.dynamicTags.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = '';
    },

    handleDelete(index, row) {
      wordApi.deleteWordRow(row.id).then(response => {
        if (!response.code) {
          this.fetchData();
        }
      });
    },
    submitForm(formData){
      this.$refs["formData"].validate(valid => {
        if (valid) {
          const data = {
              title: this.formData.name,
              tag: this.dynamicTags,
              desc: this.formData.desc,
              content: this.formData.value
          };
          console.log(this.dynamicTags);
          blogApi.postBlogRow(data).then(
            response =>{
              this.$router.push('/blog')
            });
        }else{
          console.log('验证失败')
          return false
        }
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
.el-tag + .el-tag {
    margin-left: 10px;
  }
.button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}
.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}
</style>
