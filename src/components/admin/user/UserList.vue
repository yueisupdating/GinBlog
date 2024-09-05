<template>
  <div>
 <a-card>
    <a-row :gutter="16">
        <a-col :span="12" >
            <a-input-search placeholder="输入要查询的用户名" enter-button allowClear v-model="queryParam.username" @search="getUserList">
            </a-input-search>
        </a-col>

        <a-col :span="12" >
            <a-button type="primary" @click="addUserVisible = true">新建用户</a-button>
        </a-col>
    </a-row>

    <a-table style="margin-top: 15px"
        :dataSource="dataSource"
        bordered
        :columns='columns'
        rowKey="ID"
        :pagination="pagination"

    >
        <span slot="role" slot-scope="data">{{ data == 1 ? '管理员' : '普通用户' }}</span>

        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="editUser(data.ID)"

            >编辑</a-button>
            <a-button
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deleteUser(data.ID)"
            >删除</a-button>
            <a-button type="info" icon="info">修改密码</a-button>
          </div>
        </template>

    </a-table>
 </a-card>
              <!--新增用户按钮-->
    <a-modal title="新增用户" closeable="true" :visible="addUserVisible"
      @ok="addUserOk" @cancel="addUserCancel" destroyOnClose>
      <a-form-model :model="addFormData" ref="addFormData" :rules="addUserRules">
        <a-form-model-item prop="username" label="请输入用户名">
          <a-input v-model="addFormData.username"></a-input>
        </a-form-model-item>

        <a-form-model-item prop="password" label="请输入密码">
          <a-input-password type="password" v-model="addFormData.password"></a-input-password>
        </a-form-model-item>

        <a-form-model-item prop="checkpassword" label="请再次输入密码">
          <a-input-password type="password" v-model="addFormData.checkpassword"></a-input-password>
        </a-form-model-item>

      </a-form-model>
    </a-modal>

                  <!--编辑用户按钮-->
    <a-modal title="编辑用户" closeable="true" :visible="editUserVisible" @ok="editUserOk" @cancel="editUserCancel" destroyOnClose>
      <a-form-model :model="editFormData" ref="editFormData" :rules="editUserRules">
        <a-form-model-item prop="username" label="请输入用户名">
          <a-input v-model="editFormData.username"></a-input>
        </a-form-model-item>

        <a-form-model-item label="是否为管理员">
          <a-switch
            :checked="IsAdmin"
            checked-children="是"
            un-checked-children="否"
            @change="adminChange"
          />
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
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    align: 'center',
    scopedSlots: { customRender: 'role' }
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
    computed: {
      IsAdmin: function () {
        if (this.editFormData.role === 1) {
          return true
        } else {
          return false
        }
      }
    },
    data() {
        return {
            columns,
            editFormData: {
              username: '',
              role: 2
            },
            addFormData: {
              username: '',
              password: '',
              checkpassword: '',
              role: 2
            },
            editUserVisible: false,
            addUserVisible: false,
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

                  this.getUserList()
                },
                onShowChange: (current, size) => {
                  this.pagination.defaultCurrent = current
                  this.pagination.defaultPageSize = size

                  this.getUserList()
                }
            },
            editUserRules: {
              username: [
              { required: true, message: '请输入用户名', trigger: 'blur' },
              {
                min: 4,
                max: 20,
                message: '用户名必须在4到20个字符之间',
                trigger: 'blur'
              }
            ]
            },
            queryParam: {
                username: ''
            },
            addUserRules: {
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
            ],
            checkpassword: [
              { required: true, message: '请再一次输入密码', trigger: 'blur' },
              {
                validator: (rule, value, callback) => {
                  /*
                  if (this.addFormData.checkpassword === '') {
                    callback(new Error())
                  }
                  */
                  if (this.addFormData.password !== this.addFormData.checkpassword) {
                    callback(new Error('密码不一致，请重新输入'))
                  } else {
                    callback()
                  }
                },
                trigger: 'blur'
              }
              ]
            }
        }
    },
    created() {
        this.getUserList()
    },
    methods: {
        adminChange(checked) {
          if (checked) {
            this.editFormData.role = 1
          } else {
            this.editFormData.role = 2
          }
        },
        async editUser(id) {
          this.editUserVisible = true
          const { data: res } = await this.$http.get(`admin/user/get/${id}`)
          this.editFormData = res.data
          this.editFormData.id = id
        },
        async getUserList() {
            const { data: res } = await this.$http.get('admin/userList', {
            params: {
              username: this.queryParam.username
            }
            })

            if (res.status !== 200) {
                this.$message.error(res.message)
            }
            this.dataSource = res.data
            this.pagination.total = res.total
        },
        deleteUser(id) {
          this.$confirm({
            title: '确定要删除该用户吗?',
            content: '删除后无法恢复',
            onOk: async() => {
              const { data: res } = await this.$http.delete(`admin/user/delete/${id}`)
              if (res.status !== 200) return this.$message.error(res.message)
              this.$message.success('删除成功')
              this.getUserList()
            },
            onCancel: () => {
              this.$message.info('已取消删除')
            }
        })
      },
      async editUserOk() {
        this.$refs.editFormData.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
          const { data: res } = await this.$http.put(`admin/user/update/${this.editFormData.id}`, {
            username: this.editFormData.username,
            role: this.editFormData.role
          })
          if (res.status !== 200) return this.$message.error(res.message)
          this.editUserVisible = false
          this.$message.success('编辑成功')
          this.getUserList()
        })
        },
      async addUserOk() {
        this.$refs.addFormData.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
          const { data: res } = await this.$http.post('admin/user/add', {
            username: this.addFormData.username,
            password: this.addFormData.password,
            role: this.addFormData.role
          })
          if (res.status !== 200) return this.$message.error(res.message)
          this.addUserVisible = false
          this.$message.success('添加成功')
          this.getUserList()
        })
      },
      addUserCancel() {
        this.$refs.addFormData.resetFields()
        this.addUserVisible = false
        this.$message.info('新增用户已取消')
      },
      editUserCancel() {
        this.$refs.editFormData.resetFields()
        this.editUserVisible = false
        this.$message.info('编辑已取消')
      }
}
}
</script>
