<template>
    <div>
        <v-app-bar app color="indigo">
            <v-avatar color="grey"></v-avatar>
            <v-container>
                <v-btn text color="white" @click="$router.push('/')">首页</v-btn>
                <v-btn v-for="item in cateList" :key="item.id" text color="white">
                    {{ item.categoryname }}</v-btn>
            </v-container>

            <v-spacer></v-spacer>

            <v-responsive  color="white">
                <v-text-field
                dense
                flat
                hide-details
                solo-inverted
                rounded
                placeholder="请输入文章标题查找"
                dark
                ></v-text-field>
            </v-responsive>
        </v-app-bar>
    </div>
</template>

<script>
export default {
    created() {
        this.getCateList()
    },
    data() {
        return {
            cateList: []
        }
    },
    methods: {
        async getCateList() {
            const { data: res } = await this.$http.get('/admin/cateList')
            if (res.status === 200) {
                this.cateList = res.data
            }
            console.log(res)
        }
    }
}
</script>

<style>

</style>
