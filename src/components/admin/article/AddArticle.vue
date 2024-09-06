<template>
    <div>
        <span>    {{id? '编辑文章页面':'新增文章页面'}}</span>

        <a-card>
            <a-form-model :model="formData" ref="formData" :rules="articleRules">

                <a-form-model-item label="文章标题" prop="title">
                    <a-input v-model="formData.title" style="width:400px"></a-input>
                </a-form-model-item>

                <a-form-model-item label="文章描述" prop="description">
                    <a-input type="textarea" v-model="formData.description"></a-input>
                </a-form-model-item>

                <a-form-model-item label="文章分类" prop="cid">
                    <a-select @change="categoryChange" placeholder="请选择分类">
                        <a-select-option
                            v-for="item in cateList"
                            :key="item.ID"
                            :value="item.ID"
                        >cateChange
                        {{item.categoryname}}
                        </a-select-option>
                    </a-select>
                </a-form-model-item>

                <a-form-model-item label="文章缩略图" prop="img">
                    <a-upload name="file" listType="picture"
                        :action="upURL"
                        :headers="headers"
                        @change="upChange"
                    >
                        <a-button>
                            <a-icon type="upload"></a-icon>
                            点击上传图片
                        </a-button>
                            <template v-if="id">
                                <img :src="formData.img" style="width: 120px; height: 100px; margin-left: 15px" />
                            </template>
                    </a-upload>
                </a-form-model-item>

                <a-form-model-item label="文章内容" prop="content">
                    <a-input v-model="formData.content" ></a-input>
                </a-form-model-item>

                <a-form-model-item class="addButton">
                    <a-button type="primary" style="margin:20px" @click="editArticleOK(formData.id)">
                        {{
                            id?'更新':'提交'
                        }}
                    </a-button>
                    <a-button type="info" style="margin:20px" @click="resetChange">取消</a-button>
                </a-form-model-item>

            </a-form-model>
        </a-card>
    </div>

</template>

<script>
import { baseURL } from '../../../plugin/http'

export default {
    data() {
        return {
            articleRules: {
                title: [{ required: true, message: '请输入文章标题', trigger: 'blur' }],
                cid: [{ required: true, message: '请选择文章分类', trigger: 'blur' }],
                description: [
                    { required: true, message: '请输入文章描述', trigger: 'blur' },
                    { max: 120, message: '描述最多可写120个字符', trigger: 'blur' }
                ],
                // img: [{ required: true, message: '请选择文章缩略图', trigger: 'change' }],
                content: [{ required: true, message: '请输入文章内容', trigger: 'blur' }]
            },
            upURL: baseURL + '/admin/upload',
            formData: {
                id: 0,
                title: '',
                description: '',
                img: '',
                cid: undefined,
                content: ''
            },
            cateList: [],
            headers: {}
        }
    },
    props: ['id'],
    created() {
        this.getCateList()
        this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
        if (this.id) {
            this.getArticle(this.id)
        }
    },
    methods: {
        categoryChange(value) {
            this.formData.cid = value
        },
        async getCateList() {
            const { data: res } = await this.$http.get('admin/cateList')
            if (res.status !== 200) {
                this.$message.error(res.message)
            }
            this.cateList = res.data
        },
        upChange(info) {
            if (info.file.status === 'done') {
                this.$message.success('图片上传成功')
                this.formData.img = info.file.response.url
                console.log(this.formData.img)
            } else if (info.file.status === 'error') {
                this.$message.error('图片上传失败')
            }
        },
        async getArticle() {
            const { data: res } = await this.$http.get(`/admin/article/get/${this.id}`)
            if (res.status !== 200) {
                this.$message.error(res.message)
            }
            this.formData = res.data
            this.formData.id = this.id
        },
        resetChange() {
            this.$refs.formData.resetFields()
            this.$message.info('已取消编辑')
        },
        async editArticleOK(id) {
            this.$refs.formData.validate(async (valid) => {
                if (!valid) return this.$message.error('参数验证未通过，请按要求录入文章内容')
                if (id === 0) {
                    const { data: res } = await this.$http.post('/admin/article/add', this.formData)
                    if (res.status !== 200) {
                        return this.$message.error(res.message)
                    }
                    this.$router.push('/admin/articleList')
                    this.$message.success('已成功新增文章')
                } else {
                    const { data: res } = await this.$http.put(`/admin/article/update/${id}`, this.formData)
                    if (res.status !== 200) {
                        return this.$message.error(res.message)
                    }
                    this.$router.push('/admin/articleList')
                    this.$message.success('已成功修改文章')
                }
        })
    }
    }
}
</script>

<style scoped>
.addButton{
    display: flex;
    justify-content: flex-end;
    margin:30px;
}
</style>
