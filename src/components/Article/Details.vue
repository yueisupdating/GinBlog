<template>
<div>
    <div class="justify-center text-h5 d-flex font-weight-bold">{{formData.title}}
    </div>
    <v-divider></v-divider>
    <v-alert color="indigo" outlined max-width="100%" class="pa-3 ma-5" elevation="2">{{formData.description}}</v-alert>
    <div class="pa-3 ma-5 text-justify">{{formData.content}}</div>
</div>
</template>

<script>
export default {
    data() {
        return {
            formData: {},
            id: this.$route.params.id
        }
    },
    methods: {
        async getArticleData() {
            const { data: res } = await this.$http.get(`admin/article/get/${this.id}`)
            if (res.status === 200) {
                this.formData = res.data
            }
        }
    },
    created() {
        this.getArticleData()
        this.$http.get(`viewCount/incr/${this.id}`)
    }
}
</script>

<style>

</style>
