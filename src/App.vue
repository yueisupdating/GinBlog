<template>
  <a-config-provider :locale="locale">
  <div id="app">
    <router-view/>
  </div>
  </a-config-provider>
</template>

<script>
  import zhCN from 'ant-design-vue/es/locale/zh_CN'
  import dayjs from 'dayjs'
  import 'dayjs/locale/zh-cn'
  dayjs.locale('zh-cn')

  export default {
    data() {
      return {
        locale: zhCN
      }
    }
  }
</script>
