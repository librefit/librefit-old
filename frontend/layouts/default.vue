<template>
  <v-app>
    <Navbar />
    <Snackbar></Snackbar>
    <v-main class="ma-7">
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: 'Default',
  data () {
    return {
      title: 'Vuetify.js'
    }
  }
}
</script>
