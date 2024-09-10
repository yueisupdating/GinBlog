<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
import HomeView from './views/HomeView.vue'

export default {
  name: 'App',
  comments: { HomeView }

}
</script>
