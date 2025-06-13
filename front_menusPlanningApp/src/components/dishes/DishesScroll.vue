<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useDishesStore } from '../../stores/dishes.store';
import type { Dish } from '../../models/dish.model';
import ConsultDish from './ConsultDish.vue';
import EditDishForm from './EditDishForm.vue';
import ConfirmDeleteDish from './ConfirmDeleteDish.vue';

const searchQuery = ref('');
const dishesStore = useDishesStore();
const editingDish = ref<Dish | null>(null);
const deletingDishId = ref<string | null>(null);

onMounted(() => {
  dishesStore.loadDishes({ page: 1, itemsPerPage: 10 });
});

const filteredDishes = computed(() =>
  (dishesStore.dishList || []).filter(dish =>
    dish?.name ? dish.name.toLowerCase().includes(searchQuery.value.toLowerCase()) : false
  )
);

const dialogConsultVisible = ref(false);
const dialogEditVisible = ref(false);
const dialogDeleteVisible = ref(false);

function consult(dish: Dish) {
  editingDish.value = { ...dish };
  dialogConsultVisible.value = true;
}

function edit(dish: Dish) {
  editingDish.value = { ...dish };
  dialogEditVisible.value = true;
}

function askDelete(dishId: string) {
  deletingDishId.value = dishId;
  dialogDeleteVisible.value = true;
}

function handleEditSubmitted() {
  dialogEditVisible.value = false;
 // dishesStore.loadDishes({ page: 1, itemsPerPage: 10 }); // Optional refresh
}

async function handleDeleteConfirmed(id: string | null) {
  if (!id) return;
  try {
    await dishesStore.removeDish(id);
    dialogDeleteVisible.value = false;
  //  dishesStore.loadDishes({ page: 1, itemsPerPage: 10 }); // Optional refresh
  } catch (e) {
    console.error('Failed to delete dish', e);
  }
}
</script>

<template>
  <v-card class="mx-auto" max-width="500">
    <v-card-title>List of recipes</v-card-title>

    <v-text-field
      v-model="searchQuery"
      label="Search a dish"
      variant="outlined"
      density="compact"
      class="mx-4 my-2"
    ></v-text-field>

    <v-divider></v-divider>

    <v-progress-linear v-if="dishesStore.loading" indeterminate></v-progress-linear>

    <v-virtual-scroll :items="filteredDishes" height="320" item-height="48">
      <template v-slot:default="{ item }">
        <v-list-item :title="item.name">
          <template v-slot:prepend>
           <v-avatar size="40">
            <img :src="item.imageUrl" alt="Dish image" style="object-fit: cover; width: 100%; height: 100%;" />
          </v-avatar>

          </template>
          <template v-slot:append>
            <v-btn icon="mdi-eye" size="x-small" variant="tonal" @click="consult(item)"></v-btn>
            <v-btn icon="mdi-pencil" size="x-small" variant="tonal" @click="edit(item)"></v-btn>
            <v-btn icon="mdi-delete" size="x-small" variant="tonal" @click="askDelete(item.id!)"></v-btn>
          </template>
        </v-list-item>
      </template>
    </v-virtual-scroll>
  </v-card>

  <ConsultDish
  v-model="dialogConsultVisible"
  :dish="editingDish"
/>

  <EditDishForm
    v-model="dialogEditVisible"
    :dish="editingDish"
    @submitted="handleEditSubmitted"
  />

  <ConfirmDeleteDish
    v-model:show="dialogDeleteVisible"
    :dishId="deletingDishId"
    @confirmDelete="handleDeleteConfirmed"
  />
</template>