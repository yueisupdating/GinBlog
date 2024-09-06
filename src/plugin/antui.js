import Vue from 'vue'
import { Button, FormModel, Input, Icon, message, Layout, Menu, Card, Row, Col, Table, Pagination, ConfigProvider, Modal, Select, Switch, Upload } from 'ant-design-vue'

message.config({
  top: '60px',
  duration: 2,
  maxCount: 3
})

Vue.prototype.$message = message
Vue.prototype.$confirm = Modal.confirm

Vue.use(Button)
Vue.use(Input)
Vue.use(FormModel)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Menu)
Vue.use(Card)
Vue.use(Col)
Vue.use(Row)
Vue.use(Table)
Vue.use(Pagination)
Vue.use(ConfigProvider)
Vue.use(Modal)
Vue.use(Select)
Vue.use(Switch)
Vue.use(Upload)
