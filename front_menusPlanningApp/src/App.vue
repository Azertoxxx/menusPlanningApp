<script setup lang="ts">
import Navbar from './components/Navbar.vue';
import Upbar from './components/Upbar.vue';
</script>

<template>
    <v-app>
      <Upbar/>
      <Navbar/>
      <v-main>
        <v-container>
          <router-view />
        </v-container>
      </v-main>
    </v-app>
</template>

<style scoped>
</style>