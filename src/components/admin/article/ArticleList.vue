<template>
  <div>
 <a-card>
    <a-row :gutter="16">
        <a-col :span="12" >
            <a-input-search placeholder="输入要查询的文章题目" enter-button allowClear v-model="queryParam.title " @search="getArticleList">
            </a-input-search>
        </a-col>

        <a-col :span="8" >
            <a-select placeholder="请选择分类" style="width: 200px" @change="cateChange">
                <a-select-option
                v-for="item in cateList"
                :key="item.ID"
                :value="item.ID"
                >{{ item.categoryname }}</a-select-option>
            </a-select>
        </a-col>
    </a-row>

    <a-table style="margin-top: 15px"
        :dataSource="dataSource"
        bordered
        :columns='columns'
        rowKey="ID"
        :pagination="pagination"
    >
        <span class="ArtImg" slot="img" slot-scope="img">
          <img :src="img" />
        </span>

        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="editArticle(data.ID)"
            >编辑</a-button>
            <a-button
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deleteArticle(data.ID)"
            >删除</a-button>
          </div>
        </template>
    </a-table>
 </a-card>
                  <!--编辑文章按钮-->
    <a-modal title="编辑文章" closeable="true" :visible="editArticleVisible" @ok="editArticleOk" @cancel="editArticleCancel" destroyOnClose>
      <a-form-model :model="editFormData" ref="editFormData">
        <a-form-model-item prop="title " label="请修改文章名">
          <a-input v-model="editFormData.title "></a-input>
        </a-form-model-item>

        <a-form-model-item prop="description" label="请修改文章描述">
          <a-input v-model="editFormData.description "></a-input>
        </a-form-model-item>

      </a-form-model>
    </a-modal>

</div>
</template>

<style scoped>
.ArtImg img{
   height:80px;
   width:100px;
}
</style>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'id',
    align: 'center'
  },
  {
    title: '文章名',
    dataIndex: 'title',
    width: '10%',
    key: 'title',
    align: 'center'
  },
  {
    title: '描述',
    dataIndex: 'description',
    width: '10%',
    key: 'description',
    align: 'center'
  },
    {
    title: '缩略图',
    dataIndex: 'img',
    width: '15%',
    key: 'img',
    align: 'center',
    scopedSlots: { customRender: 'img' }
  },
  {
    title: '操作',
    width: '20%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
    data() {
        return {
            columns,
            editFormData: {
                title: '',
                description: '',
                content: '',
                img: ''
            },
            editArticleVisible: false,
            dataSource: [],
            cateList: [],
            pagination: {
                pageSizeOptions: ['1', '2', '5', '10', '20'],
                defaultCurrent: 1,
                defaultPageSize: 5,
                total: 0,
                showSizeChanger: true,
                showTotal: (total) => `共 ${total} 条数据`,
                onChange: (page, pageSize) => {
                  this.pagination.defaultCurrent = page
                  this.pagination.defaultPageSize = pageSize

                  this.getArticleList()
                },
                onShowChange: (current, size) => {
                  this.pagination.defaultCurrent = current
                  this.pagination.defaultPageSize = size

                  this.getArticleList()
                }
            },
            queryParam: {
                title: ''
            }
        }
    },
    created() {
        this.getArticleList()
        this.getCateList()
    },
    methods: {
        async cateChange(id) {
            const { data: res } = await this.$http.get(`admin/article/category/${id}`)

            if (res.status !== 200) {
                this.$message.error(res.message)
            }
            this.dataSource = res.data
            this.pagination.total = res.total
        },
        async getCateList() { //, { params: { categoryname: '' } }
            const { data: res } = await this.$http.get('admin/cateList')

            if (res.status !== 200) {
                this.$message.error(res.message)
            }
            this.cateList = res.data
        },
        async editArticle(id) {
          this.editArticleVisible = true
          const { data: res } = await this.$http.get(`admin/article/get/${id}`)

          this.editFormData = res.data
          this.editFormData.id = id
        },
        async getArticleList() {
            const { data: res } = await this.$http.get('admin/articleList', {
            params: {
                title: this.queryParam.title
            }
            })

            if (res.status !== 200) {
                this.$message.error(res.message)
            }
            this.dataSource = res.data
            this.pagination.total = res.total
        },
        deleteArticle(id) {
          this.$confirm({
            title: '确定要删除该文章吗?',
            content: '删除后无法恢复',
            onOk: async() => {
              const { data: res } = await this.$http.delete(`admin/article/delete/${id}`)
              if (res.status !== 200) return this.$message.error(res.message)
              this.$message.success('删除成功')
              this.getArticleList()
            },
            onCancel: () => {
              this.$message.info('已取消删除')
            }
        })
      },
      async editArticleOk() {
          const { data: res } = await this.$http.put(`admin/article/update/${this.editFormData.id}`, {
            title: this.editFormData.title,
            description: this.editFormData.description,
            img: this.editFormData.img
          })
          if (res.status !== 200) return this.$message.error(res.message)
          this.editArticleVisible = false
          this.$message.success('编辑成功')
          this.getArticleList()
        },
      editArticleCancel() {
        this.$refs.editFormData.resetFields()
        this.editArticleVisible = false
        this.$message.info('编辑已取消')
      }
}
}
</script>
