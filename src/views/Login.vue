<template>
  <div class="container">
    <div class="loginBox">

      <a-form-model :model="formData" :rules="rules" ref="formData" class="loginForm">
        <h1 class="loginTitle">登录界面</h1>
        <a-form-model-item prop="username">
          <a-input placeholder="请输入用户名" v-model="formData.username">
            <a-icon slot="prefix" type="user"/>
          </a-input>
        </a-form-model-item>

        <a-form-model-item prop="password">
          <a-input placeholder="请输入密码" v-model="formData.password" type="password">
          <a-icon slot="prefix" type="lock"></a-icon>
          </a-input>
        </a-form-model-item>

        <a-form-model-item >
          <div class="loginButton">
            <a-button type="primary" @click="login" style="margin:10px">登录</a-button>
            <a-button type="info" @click='reset' style="margin:10px">重置</a-button>
          </div>
        </a-form-model-item>
      </a-form-model>

    </div>
  </div>
</template>

<style scoped>
.container {
  height: 100%;
  background-color: #1F1F1F;
}
.loginBox {
  background-color: white;
  height: 340px;
  width:400px;
  position:absolute;
  top:50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 10px;
}
.loginButton{
  margin: 20px;
  display: flex;
  justify-content: flex-end;
}
.loginForm {
  width: 100%;
  position: absolute;
  bottom: 0%;
  padding: 0 10px;
  box-sizing: border-box;
}
.loginTitle{
  display: flex;
  justify-content: center;
}
</style>

<script>
export default {
  data() {
    return {
      formData: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          {
            min: 4,
            max: 20,
            message: '用户名必须在4到20个字符之间',
            trigger: 'blur'
          }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          {
            min: 5,
            max: 50,
            message: '密码必须在5到50个字符之间',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  methods: {
    login() {
      this.$refs.formData.validate(async (valid) => {
        if (!valid) return this.$message.error('请按要求重新输入')
        const { data: res } = await this.$http.post('login', this.formData)
        if (res.status !== 200) return this.$message.error(res.message)
        window.sessionStorage.setItem('token', res.token)
        this.$router.push('admin/index')
      })
    },
    reset() {
      this.$refs.formData.resetFields()
    }
  }
}
</script>
