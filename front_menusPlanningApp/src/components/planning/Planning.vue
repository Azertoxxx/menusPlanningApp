<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useDishesStore } from '../../stores/dishes.store';
import { usePlanningsStore } from '../../stores/plannings.store';
import type { Planning } from '../../models/planning.model';
import { PlanningSlot } from '../../models/planning.model';
import DishSelector  from './DishSelector.vue';

const planningsStore = usePlanningsStore();
const dishesStore = useDishesStore();

onMounted(async () => {
  await dishesStore.loadDishes({ page: 1, itemsPerPage: 100 });
  await planningsStore.loadPlanning();
});

const daysOfWeek = [
  "MONDAY",
  "TUESDAY",
  "WEDNESDAY",
  "THURSDAY",
  "FRIDAY",
  "SATURDAY",
  "SUNDAY",
];

const getSlotEnum = (day: string, period: string): PlanningSlot =>
  `${day}_${period}` as PlanningSlot;

const getDishNameForSlot = (slot: PlanningSlot) => {
  const match = planningsStore.planningList.find((p) => p.slot === slot);
  return match?.dish?.name || "";
};

const dialogVisible = ref(false);
const selectedSlot = ref<PlanningSlot | null>(null);
const selectedDishId = ref<string>("");

const openDialog = (slot: PlanningSlot) => {
  selectedSlot.value = slot;
  selectedDishId.value =
    planningsStore.planningList.find((p) => p.slot === slot)?.dish?.id || "";
  dialogVisible.value = true;
};

const closeDialog = () => {
  dialogVisible.value = false;
  selectedSlot.value = null;
  selectedDishId.value = "";
};

const assignDishToSlot = async () => {
  if (!selectedSlot.value || !selectedDishId.value) return;

  const dish = dishesStore.dishList.find((d) => d.id === selectedDishId.value);
  if (!dish) return;

  const existingPlanning = planningsStore.planningList.find(
    (p) => p.slot === selectedSlot.value
  );

  const planning: Planning = {
    slot: selectedSlot.value,
    dish: dish,
    id: existingPlanning?.id,
  };

  try {
    if (existingPlanning) {
      await planningsStore.savePlanning(planning);
    } else {
      await planningsStore.addPlanning(planning);
    }
  } catch (error) {
    console.error("Failed to assign dish", error);
  } finally {
    closeDialog();
  }
};

const shoppingListVisible = ref(false);
const shoppingList = ref<string[]>([]);

const generateShoppingList = () => {
  const dishIdsInPlanning = new Set(
    planningsStore.planningList.map((p) => p.dish?.id).filter(Boolean)
  );

  const selectedDishes = dishesStore.dishList.filter((d) =>
    dishIdsInPlanning.has(d.id)
  );

  const list: string[] = [];

  selectedDishes.forEach((dish) => {
    list.push(`🍽️ ${dish.name}`);
    list.push(` ${dish.ingredients}`);
    list.push("");
  });

  shoppingList.value = list;
  shoppingListVisible.value = true;
};
</script>

<template>
  <div class="weekly-grid">
    <div class="day-column" v-for="day in daysOfWeek" :key="day">
      <div class="day-header">{{ day }}</div>

      <div
        class="slot-button"
        v-for="period in ['LUNCH', 'DINNER']"
        :key="period"
        @click="openDialog(getSlotEnum(day, period))"
      >
        {{ getDishNameForSlot(getSlotEnum(day, period)) || period }}
      </div>
    </div>


    <DishSelector
      :visible="dialogVisible"
      :slot="selectedSlot"
      :selectedDishId="selectedDishId"
      :dishes="dishesStore.dishList"
      @update:selectedDishId="(val: string) => (selectedDishId = val)"
      @save="assignDishToSlot"
      @close="closeDialog"
    />
  </div>
  <div>
    <v-btn color="success" type="button" @click="generateShoppingList">
      Generate Shopping List
    </v-btn>
  </div>

  <div v-if="shoppingListVisible" class="shopping-list">
    <h3>🛒 Shopping List</h3>
    <pre>{{ shoppingList.join('\n') }}</pre>
  </div>
  <div>
    <v-btn color="success" type="submit">Send shopping list by email or SMS</v-btn>
  </div>
</template>


<style scoped>
.weekly-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 1rem;
  margin-top: 2rem;
}

.day-column {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.day-header {
  font-weight: bold;
  margin-bottom: 1rem;
}

.slot-button {
  width: 100%;
  padding: 0.75rem;
  margin-bottom: 0.5rem;
  background-color: #f0f0f0;
  border: 1px solid #ccc;
  text-align: center;
  cursor: pointer;
  border-radius: 5px;
}

.slot-button:hover {
  background-color: #e0e0e0;
}

.shopping-list {
  margin-top: 2rem;
  padding: 1rem;
  background: #f9f9f9;
  border: 1px solid #ccc;
  border-radius: 8px;
  white-space: pre-wrap;
}
</style>