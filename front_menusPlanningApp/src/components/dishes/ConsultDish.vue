<script setup lang="ts">
import type { Dish } from '../../models/dish.model';

const props = defineProps<{
  modelValue: boolean;
  dish: Dish | null;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void;
}>();

const closeDialog = () => {
  emit('update:modelValue', false);
};
</script>

<template>
  <v-dialog v-model="props.modelValue" max-width="600">
    <v-card>
      <v-card-title>Consult dish</v-card-title>
      <v-avatar size="40">
        <img :src="props.dish?.imageUrl" alt="Dish image" style="object-fit: cover; width: 100%; height: 100%;" />
      </v-avatar>
      <v-card-text>
        <v-row>
          <v-col cols="12">
            <strong>Nom:</strong> {{ props.dish?.name || '-' }}
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12">
            <strong>Description:</strong> {{ props.dish?.description || '-' }}
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12">
            <strong>Ingrédients:</strong> {{ props.dish?.ingredients || '-' }}
          </v-col>
        </v-row>

        <v-row v-if="props.dish?.imageUrl">
          <v-col cols="12">
            <strong>Image:</strong><br />
            <v-img :src="props.dish.imageUrl" max-width="200" />
          </v-col>
        </v-row>
      </v-card-text>

      <v-card-actions class="justify-end">
        <v-btn text @click="closeDialog">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
