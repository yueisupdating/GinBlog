<template>
        <v-col>
        <v-card v-for="item in artList" :key="item.id" @click="$router.push(`/detail/${item.ID}`)">
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
                        <v-icon>{{'mdi-calendar-month'}}</v-icon>
                        <span>{{item.CreatedAt | dateformat("YYYY-MM-DD HH:MM")}}</span>
                    </v-card-text>
                </v-col>
            </v-row>
        </v-card>

        <div style="margin: 20px;">
            <v-pagination v-model="queryParam.pagenum" :length="Math.ceil(total/queryParam.pagesize)"  total-visible="5" @input="getArtList()">
            </v-pagination>
        </div>
    </v-col>
</template>

<script>
export default {
  data() {
    return {
      artList: [],
      queryParam: {
        pagesize: 3,
        pagenum: 1
      },
      total: 0,
      id: this.$route.params.id
    }
  },
  created() {
    this.getArtList()
  },
  methods: {
    // 获取文章列表
    async getArtList() {
      const { data: res } = await this.$http.get(`admin/article/category/${this.id}`,
        {
          params: {
            id: this.id,
            pagesize: this.queryParam.pagesize,
            pagenum: this.queryParam.pagenum
          }

        }
      )
      this.artList = res.data
      this.total = res.total
    }
  }
}
</script>

<style>

</style>
