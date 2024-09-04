<template>
 <a-card>
    <a-row :gutter="16">
        <a-col :span="12" >
            <a-input-search placeholder="输入要查询的用户名" enter-button allowClear v-model="queryParam.username" @search="getUserList">
            </a-input-search>
        </a-col>

        <a-col :span="12" >
            <a-button type="primary">新建用户</a-button>
        </a-col>
    </a-row>

    <a-table style="margin-top: 15px"
        :dataSource="dataSource"
        bordered
        :columns='columns'
        rowKey="ID"
        :pagination="pagination"

    >
        <span slot="role" slot-scope="text,data">{{ data === 1 ? '管理员' : '普通用户' }}</span>

        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
              type="primary"
              icon="edit"
              style="margin-right: 15px"

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
    data() {
        return {
            columns,
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

            queryParam: {
                username: ''
            }
        }
    },
    created() {
        this.getUserList()
    },
    methods: {
        async getUserList() {
            const { data: res } = await this.$http.get('admin/users', {
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
    }
}
}
</script>
