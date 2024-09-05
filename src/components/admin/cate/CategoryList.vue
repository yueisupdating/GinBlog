<template>
  <div>
 <a-card>
    <a-row :gutter="16">
        <a-col :span="12" >
            <a-button type="primary" @click="addcateVisible = true">新建分类</a-button>
        </a-col>
    </a-row>

    <a-table style="margin-top: 15px"
        :dataSource="dataSource"
        bordered
        :columns='columns'
        rowKey="ID"
        :pagination="pagination"
    >

        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="editcate(data.ID)"

            >编辑</a-button>
            <a-button
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deletecate(data.ID)"
            >删除</a-button>
            <a-button type="info" icon="info">修改密码</a-button>
          </div>
        </template>

    </a-table>
 </a-card>
              <!--新增分类按钮-->
    <a-modal title="新增分类" closeable="true" :visible="addcateVisible"
      @ok="addcateOk" @cancel="addcateCancel" destroyOnClose>
      <a-form-model :model="addFormData" ref="addFormData" :rules="addcateRules">
        <a-form-model-item prop="categoryname" label="请输入分类名">
          <a-input v-model="addFormData.categoryname"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

                  <!--编辑分类按钮-->
    <a-modal title="编辑分类" closeable="true" :visible="editcateVisible" @ok="editcateOk" @cancel="editcateCancel" destroyOnClose>
      <a-form-model :model="editFormData" ref="editFormData" :rules="editcateRules">
        <a-form-model-item prop="categoryname" label="请输入分类名">
          <a-input v-model="editFormData.categoryname"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

</div>
</template>

<style scoped>

</style>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类名',
    dataIndex: 'categoryname',
    width: '20%',
    key: 'categoryname',
    align: 'center'
  },
  {
    title: '操作',
    width: '30%',
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
              categoryname: ''
            },
            addFormData: {
              categoryname: ''
            },
            editcateVisible: false,
            addcateVisible: false,
            dataSource: [],
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

                  this.getcateList()
                },
                onShowChange: (current, size) => {
                  this.pagination.defaultCurrent = current
                  this.pagination.defaultPageSize = size

                  this.getcateList()
                }
            },
            editcateRules: {
              categoryname: [
              { required: true, message: '请输入分类名', trigger: 'blur' },
              {
                min: 1,
                max: 20,
                message: '分类名必须在1到20个字符之间',
                trigger: 'blur'
              }
            ]
            },
            queryParam: {
                categoryname: ''
            },
            addcateRules: {
              categoryname: [
              { required: true, message: '请输入分类名', trigger: 'blur' },
              {
                min: 1,
                max: 20,
                message: '分类名必须在1到20个字符之间',
                trigger: 'blur'
              }
            ]
            }
        }
    },
    created() {
        this.getcateList()
    },
    methods: {
        async editcate(id) {
          this.editcateVisible = true
          const { data: res } = await this.$http.get(`admin/cate/get/${id}`)
          this.editFormData = res.data
          this.editFormData.id = id
        },
        async getcateList() {
            const { data: res } = await this.$http.get('admin/cateList', {
            params: {
              categoryname: this.queryParam.categoryname
            }
            })
            if (res.status !== 200) {
                this.$message.error(res.message)
            }
            this.dataSource = res.data
            this.pagination.total = res.total
        },
        deletecate(id) {
          this.$confirm({
            title: '确定要删除该分类吗?',
            content: '删除后无法恢复',
            onOk: async() => {
              const { data: res } = await this.$http.delete(`admin/cate/delete/${id}`)
              if (res.status !== 200) return this.$message.error(res.message)
              this.$message.success('删除成功')
              this.getcateList()
            },
            onCancel: () => {
              this.$message.info('已取消删除')
            }
        })
      },
      async editcateOk() {
        this.$refs.editFormData.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
          const { data: res } = await this.$http.put(`admin/cate/update/${this.editFormData.id}`, {
            categoryname: this.editFormData.categoryname
          })
          if (res.status !== 200) return this.$message.error(res.message)
          this.editcateVisible = false
          this.$message.success('编辑成功')
          this.getcateList()
        })
        },
      async addcateOk() {
        this.$refs.addFormData.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
          const { data: res } = await this.$http.post('admin/cate/add', {
            categoryname: this.addFormData.categoryname
          })
          if (res.status !== 200) return this.$message.error(res.message)
          this.addcateVisible = false
          this.$message.success('添加成功')
          this.getcateList()
        })
      },
      addcateCancel() {
        this.$refs.addFormData.resetFields()
        this.addcateVisible = false
        this.$message.info('新增分类已取消')
      },
      editcateCancel() {
        this.$refs.editFormData.resetFields()
        this.editcateVisible = false
        this.$message.info('编辑已取消')
      }
}
}
</script>
