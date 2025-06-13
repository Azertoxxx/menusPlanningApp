<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{
  show: boolean;
  dishId: string | null;
}>();

const emit = defineEmits<{
  (e: "update:show", value: boolean): void;
  (e: "confirmDelete", value: string | null): void;
}>();

const localShow = ref(props.show);

watch(() => props.show, (val) => {
  localShow.value = val;
});

const closeDialog = () => {
  emit('update:show', false);
};

const submitDelete = () => {
  emit('confirmDelete', props.dishId);
};
</script>

<template>
  <v-dialog v-model="localShow" max-width="600" @update:modelValue="emit('update:show', $event)">
    <v-card>
      <v-card-title>Delete confirmation </v-card-title>
      <v-card-text>Are you sure to delete this dish ?</v-card-text>
      <v-card-actions>
        <v-btn text @click="closeDialog">Cancel</v-btn>
        <v-btn color="red" @click="submitDelete">Delete</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
