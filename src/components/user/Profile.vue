<template>
    <a-card>
        <a-form-model :model="formData" ref="formData">
            <a-form-model-item label="请输入用户名" prop="name">
                <a-input style="width: 200px" v-model="formData.name"></a-input>
            </a-form-model-item>

            <a-form-model-item label="请输入个人介绍" prop="description">
                <a-input style="width: 400px" v-model="formData.description"></a-input>
            </a-form-model-item>

            <a-form-model-item label="请输入qq账号" prop="qq">
                <a-input style="width: 200px" v-model="formData.qq"></a-input>
            </a-form-model-item>

            <a-form-model-item label="请输入微信号" prop="weChat">
                <a-input style="width: 200px" v-model="formData.weChat"></a-input>
            </a-form-model-item>

            <a-form-model-item label="请输入电子邮箱账号" prop="email">
                <a-input style="width: 200px" v-model="formData.email"></a-input>
            </a-form-model-item>

            <a-form-model-item label="用户头像" prop="avatar">
                    <a-upload name="file" listType="picture"
                        :action="upURL"
                        :headers="headers"
                        @change="avatarChange"
                    >
                        <a-button>
                            <a-icon type="upload"></a-icon>
                            点击上传图片
                        </a-button>

                        <template v-if="formData.avatar">
                            <img :src="formData.avatar" style="width: 120px; height: 100px" />
                        </template>
                    </a-upload>
                </a-form-model-item>

                <a-form-model-item label="头像背景图" prop="img">
                    <a-upload name="file" listType="picture"
                        :action="upURL"
                        :headers="headers"
                        @change="imgChange"
                    >
                        <a-button>
                            <a-icon type="upload"></a-icon>
                            点击上传图片
                        </a-button>
                        <template v-if="formData.img">
                            <img :src="formData.img" style="width: 120px; height: 100px" />
                        </template>
                    </a-upload>
                </a-form-model-item>

                <a-form-model-item class="upButton">
                    <a-button type="primary" @click="updateProfile" style="margin: 10px">提交</a-button>
                    <a-button type="primary" @click="resetProfile" style="margin: 10px">重置</a-button>
                </a-form-model-item>
        </a-form-model>
    </a-card>

</template>
<style scoped>
    .upButton {
        display: flex;
        justify-content: flex-end;
    }
</style>
<script>
import { baseURL } from '../../plugin/http'
export default {
    data() {
        return {
            formData: {
                id: 1,
                img: '',
                avatar: '',
                name: '',
                description: '',
                qq: '',
                weChat: '',
                email: ''
            },
            upURL: baseURL + '/admin/upload',
            headers: {}
        }
    },
    created() {
        this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
        this.getProfile()
    },
    methods: {
        async getProfile() {
            const id = this.formData.id
             const { data: res } = await this.$http.get(`admin/profile/get/${this.formData.id}`)
             if (res.status !== 200) {
                this.$message.error(res)
             }
             this.formData = res.data
             this.formData.id = id
        },
        avatarChange(info) {
            if (info.file.status === 'done') {
                this.$message.success('图片上传成功')
                const imgUrl = info.file.response.url
                this.formData.avatar = imgUrl
            } else if (info.file.status === 'error') {
                this.$message.error('图片上传失败')
            }
        },

        imgChange(info) {
            if (info.file.status === 'done') {
                this.$message.success('图片上传成功')
                const imgUrl = info.file.response.url
                this.formData.img = imgUrl
            } else if (info.file.status === 'error') {
                this.$message.error('图片上传失败')
            }
        },

        async updateProfile() {
            console.log(this.formData.id)

            const { data: res } = await this.$http.put(`admin/profile/update/${this.formData.id}`, this.formData)
            if (res.status !== 200) {
                this.$message.error(res)
            }
            this.$message.success('个人信息成功更新')
            this.$router.push('/admin/index')
        },
        async resetProfile() {
            this.$refs.formData.resetFields()
            this.$message.info('已取消编辑')
        }
    }
}
</script>
