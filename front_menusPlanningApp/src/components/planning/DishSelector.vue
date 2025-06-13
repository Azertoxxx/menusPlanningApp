<script setup lang="ts">
import { ref, watch } from 'vue';
import type { PlanningSlot } from '../../models/planning.model';
import type { Dish } from '../../models/dish.model';

const props = defineProps<{
  visible: boolean;
  slot: PlanningSlot | null;
  selectedDishId: string;
  dishes: Dish[];
}>();

const emit = defineEmits<{
  (e: 'update:selectedDishId', value: string): void;
  (e: 'save'): void;
  (e: 'close'): void;
}>();

const localSelectedDishId = ref(props.selectedDishId);

watch(
  () => props.selectedDishId,
  (newVal) => {
    localSelectedDishId.value = newVal;
  }
);

const onSave = () => {
  emit('update:selectedDishId', localSelectedDishId.value);
  emit('save');
};

const onClose = () => {
  emit('close');
};


const selectDish = (dishId: string | undefined) => {
  if (!dishId) return;
  localSelectedDishId.value = dishId;
};
</script>

<template>
  <div v-if="visible" class="dialog-overlay" @click.self="onClose">
    <div class="dialog">
      <h3>Select a dish for {{ slot }}</h3>

      <div class="dish-grid">
        <div
          v-for="dish in dishes"
          :key="dish.id"
          class="dish-card"
          :class="{ selected: localSelectedDishId === dish.id }"
          @click="selectDish(dish.id)"
        >
          <img v-if="dish.imageUrl" :src="dish.imageUrl" alt="dish image" class="dish-image" />
          <p class="dish-name">{{ dish.name }}</p>
        </div>
      </div>
      
      <div class="dialog-actions">
        <button @click="onSave">Save</button>
        <button @click="onClose">Cancel</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  justify-content: center;
  align-items: center;
}

.dialog {
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  min-width: 300px;
}

.dish-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: 1rem;
  margin-top: 1rem;
}

.dish-card {
  border: 2px solid transparent;
  border-radius: 8px;
  padding: 0.5rem;
  text-align: center;
  cursor: pointer;
  transition: border-color 0.2s ease, background-color 0.2s ease;
  background: #f9f9f9;
}

.dish-card.selected {
  border-color: #007bff;
  background-color: #e6f0ff;
}

.dish-image {
  width: 100%;
  height: 80px;
  object-fit: cover;
  border-radius: 4px;
}

.dish-name {
  margin-top: 0.5rem;
  font-size: 0.9rem;
}

.dialog-actions {
  margin-top: 1rem;
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
}
</style>
