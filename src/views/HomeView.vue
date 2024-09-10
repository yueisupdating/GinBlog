<template>
  <v-app app>

      <TopBar></TopBar>

      <v-main class="grey lighten-3">
        <v-container>
          <v-row>
            <v-col cols="3">
              <Nav></Nav>
            </v-col >

            <v-col cols="3">
              <v-sheet min-height="80vh" rounded="lg" min-width="120vh">
                <router-view :key="$route.path"></router-view>
              </v-sheet>
            </v-col>

          </v-row>
        </v-container>
      </v-main>

      <Footer></Footer>

  </v-app>
</template>

<script>
  import Footer from '../components/Footer.vue'
  import TopBar from '../components/TopBar.vue'
  import Nav from '../components/Nav.vue'

  export default {
    components: { Footer, TopBar, Nav }
  }
</script>
