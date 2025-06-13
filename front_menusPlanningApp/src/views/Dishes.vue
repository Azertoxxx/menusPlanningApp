<script setup lang="ts">
import { ref } from "vue";
import AddDishForm from "../components/dishes/AddDishForm.vue";
import DishesScroll from "../components/dishes/DishesScroll.vue";

const showViewDialog = ref<boolean>(false);

const onDishSubmitted = () => {
  showViewDialog.value = false;
};
</script>

<template>
  <h3>Plats</h3>
  <v-btn @click="showViewDialog = true">
    <v-icon start>mdi-plus</v-icon>
    Ajouter un plat
  </v-btn>

  <DishesScroll />

  <v-dialog v-model="showViewDialog" max-width="500px">
    <v-card>
      <v-card-title>Ajouter un plat</v-card-title>
      <v-card-text>
        <AddDishForm @submitted="onDishSubmitted" />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary" @click="showViewDialog = false">Fermer</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

</template>
