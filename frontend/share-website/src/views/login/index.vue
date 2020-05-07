<template>
  <div class="login-container">
      <el-form ref="form" :rules="rules" :model="form" label-width="80px" class="login-form">
          <h2 class="login-title">我们的故事</h2>
          <el-form-item label="帐号" prop="username">
              <el-input v-model="form.username"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
              <el-input v-model="form.password" type="password"></el-input>
          </el-form-item>
          <el-form-item>
              <el-button type="primary" @click="submitForm('form')">登录</el-button>
              <el-button type="primary" @click="dialogVisible = true">注册</el-button>
          </el-form-item>
      </el-form>
      <el-dialog
        title="注册"
        :visible.sync="dialogVisible"
        width="30%"
        :before-close="closeHandler">
            <el-form ref="form2" :rules="rules2" :model="form2" label-width="80px">
                <el-form-item label="帐号" prop="username">
                    <el-input v-model="form2.username"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model="form2.password"></el-input>
                </el-form-item>
                <el-form-item label="昵称" prop="name">
                    <el-input v-model="form2.name"></el-input>
                </el-form-item>
                <el-form-item label="email" prop="email">
                    <el-input v-model="form2.email"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm2('form2')">确认</el-button>
                </el-form-item>
            </el-form>
      </el-dialog>
  </div>
</template>

<script>
import {login, register, getUserInfo} from '@/api/login'

  export default {
    data() {
      return {
        form: {
            username: '',
            password: ''
         },
        form2: {
            username: '',
            password: '',
            name: '',
            email: '',
         }, 
         rules: {
             username: [
                  {required: true, message: '帐号不能为空', trigger: 'blur' },
             ],
             password: [
                 {required: true, message: '密码不能为空', trigger: 'blur' },
             ]
         },
         rules2: {
             username: [
                  {required: true, message: '帐号不能为空', trigger: 'blur' },
             ],
             password: [
                 {required: true, message: '密码不能为空', trigger: 'blur' },
             ],
             name: [
                 {required: true, message: '昵称不能为空', trigger: 'blur' },
             ],
             email: [
                 {required: true, message: 'email不能为空', trigger: 'blur' },
             ]   
         },
         dialogVisible: false
      }
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate(valid => {
            // console.log(valid)
            if (valid) {
                // 提交表单给后台进行验证是否正确
                login(this.form.username, this.form.password).then(response => {
                    const resp = response.data 
                    if(!resp.code) {
                        console.log("成功")
                        localStorage.setItem('share-website-user', JSON.stringify(resp.data.name))
                        localStorage.setItem('share-website-token', resp.data.id)
                        // 前往首页 
                        this.$router.push('/')
                    }else {
                        // 未通过，弹出警告
                        // alert(resp.message)
                        this.$message({
                            message: resp.message,
                            type: 'warning'
                        });
                        this.$router.push('/')
                    }
                })
            }else{
                console.log('验证失败')
                return false
            }
        })
      },
      // 注册
      submitForm2(formName) {
        this.$refs[formName].validate(valid => {
            // console.log(valid)
            if (valid) {
                const data = {
                    username: this.form2.username,
                    password: this.form2.password,
                    name: this.form2.name,
                    email: this.form2.email,
                }
                // 提交表单给后台进行验证是否正确
                register(data).then(response => {
                    const resp = response.data 
                    if(!resp.code) {
                        console.log("成功")
                        localStorage.setItem('share-website-user', JSON.stringify(resp.data.name))
                        localStorage.setItem('share-website-token', resp.data.id)
                        // 前往首页 
                        this.$router.push('/home')
                    }else {
                        // 未通过，弹出警告
                        // alert(resp.message)
                        this.$message({
                            message: resp.message,
                            type: 'warning'
                        });
                    }
                })
            }else{
                console.log('验证失败')
                return false
            }
        })
      },
      closeHandler(done) {
        this.$confirm('推出注册')
          .then(_ => {
            done();
          })
          .catch(_ => {});
      }

  }
}
</script>

<style scoped>
.login-form {
    width: 350px;
    /* 上下间隙 160px, 左右自动居中 */
    margin: 160px auto;
    background-color: rgb(255,255,255,0.8);
    padding: 28px;
    border-radius: 20px;
}
.login-container {
    position: absolute;
    width: 100%;
    height: 100%;
    background: url('../../assets/login.jpeg');
    background-repeat: no-repeat;
    background-position:center;
}
.login-title {
    color: #303133;
    text-align: center;
}
</style>