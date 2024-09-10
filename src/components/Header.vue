<template>
    <div class="headerButton">
        <a-button type="danger" @click="logout">退出</a-button>
    </div>
</template>

<script>
export default {
    methods: {
        logout() {
            window.sessionStorage.clear('token')
            this.$router.push('/login')
        }
    }
}
</script>

<style scoped>
.headerButton {
        display: flex;
        justify-content: flex-end;
        align-items: center;
        margin:10px;
    }
</style>
