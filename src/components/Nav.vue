<template>
  <v-card class="mx-auto" max-width="320">
    <v-img :src="profileInfo.img">
      <v-card-title>
        <v-col align="center">
          <v-avatar size="130" color="grey">
            <img :src="profileInfo.avatar" alt />
          </v-avatar>
          <div class="ma-4 white--text">{{profileInfo.name}}</div>
        </v-col>
      </v-card-title>
      <v-divider></v-divider>
    </v-img>

    <v-card-title>About Me:</v-card-title>
    <v-card-text>{{profileInfo.description}}</v-card-text>

    <v-divider color="indigo"></v-divider>

    <v-list nav dense>
      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="blue darken-2">{{'mdi-qqchat'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profileInfo.qq}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="green darken-2">{{'mdi-wechat'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profileInfo.weChat}}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="indigo">{{'mdi-email'}}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{profileInfo.email}}</v-list-item-content>
      </v-list-item>
    </v-list>
  </v-card>
</template>

<script>
export default {
    data() {
        return {
            profileInfo: {
                id: 1,
                img: '',
                avatar: '',
                name: '',
                description: '',
                qq: '',
                weChat: '',
                email: ''
            }
        }
    },
    created() {
        this.getProfile()
    },
    methods: {
        async getProfile() {
            const { data: res } = await this.$http.get(`admin/profile/get/${this.profileInfo.id}`)
            if (res.status === 200) {
                const id = this.profileInfo.id
                this.profileInfo = res.data
                this.profileInfo.id = id
            }
        }
    }
}
</script>

<style>

</style>
