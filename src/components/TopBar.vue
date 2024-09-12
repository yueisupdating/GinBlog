<template>
    <div>
        <v-app-bar app color="indigo">
            <v-avatar color="grey">
                <img :src=topUrl alt />
            </v-avatar>
            <v-container>
                <v-btn text color="white" @click="$router.push('/')">首页</v-btn>
                <v-btn v-for="item in cateList" :key="item.id" text color="white" @click="$router.push(`/category/${item.ID}`)">
                    {{ item.categoryname }}</v-btn>
            </v-container>

            <v-spacer></v-spacer>

            <v-btn icon color="white" text @click="$router.push(`/rank`)">排行榜</v-btn>

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
                append-icon="mdi-text-search"
                v-model="searchName"
                @change="searchTitle(searchName)"
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
            topUrl: 'http://sir7pml1w.hb-bkt.clouddn.com/Fme4x8WOm2CimAdS5Phk3bUKzAl4',
            cateList: [],
            searchName: ''
        }
    },
    methods: {
        async getCateList() {
            const { data: res } = await this.$http.get('/admin/cateList')
            if (res.status === 200) {
                this.cateList = res.data
            }
        },
        searchTitle(title) {
            if (title.length === 0) {
                this.$router.push('/')
                return
            }
            this.$router.push(`/search/${title}`)
        }
    }
}
</script>

<style>

</style>
