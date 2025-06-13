<script setup lang="ts">
import { reactive, ref, watch } from "vue";
import type { Dish } from "../../models/dish.model";
import { useDishesStore } from "../../stores/dishes.store";

const emit = defineEmits<{
  (e: "submitted"): void;
}>();
const form = reactive<Dish>({
    name: "",
    description: "",
    ingredients: "",
    imageUrl: ""
});

const dishesStore = useDishesStore();

const imageFile = ref<File | null>(null);

watch(imageFile, (file) => {
  if (!file) return;

  const reader = new FileReader();
  reader.onload = () => {
    form.imageUrl = reader.result as string;
  };
  reader.readAsDataURL(file);
});

const submitForm = async () => {
  try {
    await dishesStore.addDish(form);
    emit("submitted");
  } catch (error) {
    console.error("Failed to create dish", error);
  }
};
</script>

<template>
  <v-form @submit.prevent="submitForm">
    <v-row>
      <v-col cols="12">
        <v-text-field
          v-model="form.name"
          label="Nom du plat"
          required
        />
      </v-col>
    </v-row>
        <v-row>
      <v-col cols="12">
        <v-textarea
          v-model="form.description"
          label="Description"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-textarea
          v-model="form.ingredients"
          label="Ingredients"
          required
        />
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="12">
        <v-file-input
          label="Dish Image"
          accept="image/*"
          :multiple="false"
          v-model="imageFile"
        />
      </v-col>
    </v-row>

    <v-btn color="success" type="submit">Créer</v-btn>
  </v-form>
</template>
