<template>
    <a-menu theme="dark" mode="inline" @click="goToPage">
        <a-menu-item key="index">
            <a-icon type="dashboard" />
            <span>仪表盘</span>
        </a-menu-item>

        <a-menu-item key="userList">
            <a-icon type="user" />
            <span>用户列表</span>
        </a-menu-item>

        <a-menu-item key="categoryList">
            <a-icon type="book" />
            <span>分类列表</span>
        </a-menu-item>
        <a-sub-menu>
            <span slot="title">
                <a-icon type="file" />
                <span>文章管理</span>
            </span>

            <a-menu-item key="articleList">
                <a-icon type="snippets" />
                <span>文章列表</span>
            </a-menu-item>

            <a-menu-item key="article/add">
                <a-icon type="form" />
                <span>写文章</span>
            </a-menu-item>

        </a-sub-menu>
    </a-menu>
</template>

<script>
export default {
    methods: {
        goToPage(e) {
            this.$router.push('/admin/' + e.key).catch((err) => {
                console.log(err)
            })
        }
    }
}
</script>

<style>

</style>
