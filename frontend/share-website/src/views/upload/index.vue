<template>
  <div class="upload-div">
    <div class="op">
      <el-input v-model="title" class="input" placeholder="上传标题"></el-input>
      <el-select v-model="value" placeholder="请选择" style="display:inline-block;margin:20px;">   
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value">
        </el-option>
      </el-select>
      </div>
    <el-upload
    class="upload-demo"
    ref="upload"
    action="fakeaction"
    :on-preview="handlePreview"
    :on-remove="handleRemove"
    :file-list="fileList"
    :http-request="uploadSectionFile"
    :auto-upload="false"> 
      <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
      <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">上传到服务器</el-button>
    </el-upload>
    
  </div>
  
</template>

<script>
import uploadApi from "@/api/upload";
  export default {
    data() {
      return {
        options: [
          {
          value: 'picture',
          label: '图片'
          },
          {
          value: 'vedio',
          label: '视频'
          }
        ],
        value: 'picture',
        fileList: [],
        title: ''
      };
    },
    methods: {
      submitUpload() {
        this.$refs.upload.submit();
      },
      handleRemove(file, fileList) {
        console.log(file, fileList);
      },
      handlePreview(file) {
        console.log(file);
      },
      uploadSectionFile(param) {
        const file = param.file;
        uploadApi.upload(file, this.value, this.title)
        this.fileList = []
      }
    }
  }
</script>

<style scoped>
.upload-div {
    text-align: center;
    margin-top: 10px;
    padding: 10px;
}
.upload-demo {
    padding: 10px;
}
.input {
  display: inline-block;

}
.op {
  display: inline-block;
}
</style>

