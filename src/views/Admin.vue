<template>
  <a-layout class="container">
      <a-layout-sider breakpoint="lg">
        <div class="logo">
            <span>GinBlog</span>
        </div>
        <Sider></Sider>
      </a-layout-sider>

      <a-layout>
        <a-layout-header><Header></Header></a-layout-header>

        <a-layout-content>
            <router-view :key="$route.fullPath"></router-view>
        </a-layout-content>

        <a-layout-footer><Footer></Footer></a-layout-footer>
      </a-layout>
    </a-layout>
</template>

<script>
import Footer from '../components/Footer.vue'
import Header from '../components/Header.vue'
import Sider from '../components/Sider.vue'

export default {
    components: {
        Footer, Header, Sider
    }
}
</script>

<style>
.container {
    height: 100%;
}
.logo{
        height: 32px;
        margin: 16px;
        background-color: #fff;
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 17px;
    }
</style>
