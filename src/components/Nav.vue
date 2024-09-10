<template>
    <v-card class="mx-auto" max-width="300">
        <v-img :src="formData.img">
            <v-card-title>
                <v-col align="center">
                    <v-avatar size="120" color="grey">
                        <img :src="formData.avatar"/>
                    </v-avatar>
                    <div class="text--white">{{formData.name}}</div>
                </v-col>
            </v-card-title>

        <v-divider></v-divider>

        </v-img>
        <v-card-title>About Me:</v-card-title>
        <v-card-text>{{formData.description}}</v-card-text>
        <v-divider></v-divider>

        <v-list>
            <v-list-item>
                <v-list-item-icon>
                <v-icon color="blue darken-2">{{'mdi-qqchat'}}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>{{formData.qq}}</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon>
                <v-icon color="green darken-2">{{'mdi-wechat'}}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>{{formData.weChat}}</v-list-item-content>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon>
                <v-icon color="green darken-2">{{'mdi-email'}}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>{{formData.email}}</v-list-item-content>
            </v-list-item>
        </v-list>
    </v-card>
</template>

<script>
export default {
    data() {
        return {
            formData: {
                id: 0,
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
    methods: {
        async getProfile() {
            const { data: res } = this.$router.get(`profiles/${this.formData.id}`)
            if (res.status === 200) {
                this.formData = res.data
            }
        }
    }
}
</script>

<style>

</style>
