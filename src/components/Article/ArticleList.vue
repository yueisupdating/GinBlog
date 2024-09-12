<template>
    <v-col>
        <v-card v-for="item in artList" :key="item.id" @click="$router.push(`detail/${item.ID}`)">
            <v-row>
                <v-col cols="1">
                    <v-img :src="item.img" max-height="120" max-width="120"></v-img>
                </v-col>
                <v-col>
                    <v-card-title>
                    <v-chip>{{item.Category.categoryname}}</v-chip>
                    {{item.title}}
                    </v-card-title>

                    <v-card-subtitle>{{item.description}}</v-card-subtitle>
                    <v-divider></v-divider>
                    <v-card-text>
                        <v-icon >{{'mdi-calendar-month'}}</v-icon>
                        <span style="margin-right: 20px;">{{item.CreatedAt | dateformat("YYYY-MM-DD HH:MM")}}</span>

                        <v-icon>{{'mdi-eye'}}</v-icon>
                        <span style="margin-right: 20px;">{{item.viewCount}}</span>
                    </v-card-text>
                </v-col>
            </v-row>
        </v-card>

        <div style="margin: 20px;">
            <v-pagination v-model="queryParams.pagenum" :length="Math.ceil(total/queryParams.pagesize)"  total-visible="5" @input="getArtList()">
            </v-pagination>
        </div>
    </v-col>
</template>

<script>
export default {
    data() {
        return {
            artList: [],
            total: 0,
            queryParams: {
                pagesize: 3,
                pagenum: 1
            }
        }
    },
    created() {
        this.getArtList()
    },
    methods: {
        async getArtList() {
            const { data: res } = await this.$http.get('admin/get/articleList', {
            params: {
                pagesize: this.queryParams.pagesize,
                pagenum: this.queryParams.pagenum
            }
            })
            this.artList = res.data
            this.total = res.total
        }
    }
}
</script>

<style>

</style>
