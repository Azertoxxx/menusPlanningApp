import { defineStore } from "pinia";
import { ref } from "vue";
import type { Dish } from "../models/dish.model";
import { getDishes, createDish, updateDish, deleteDish } from "../api/dish.api";

export const useDishesStore = defineStore("dishes", () => {
    const dishList = ref<Dish[]>([]);
    const requestItemsPerPage = ref(10);
    const requestPage = ref(1);
    const requestSortBy = ref("name");
    const requestSortOrder = ref("asc");
    const requestName = ref("");
    const loading = ref(false);

    const loadDishes = async ({page, itemsPerPage}: {page: number, itemsPerPage: number | undefined}) => {
        try {
            if (page) requestPage.value = page;
            if (itemsPerPage) requestItemsPerPage.value = itemsPerPage;
            loading.value = true;

            const result = await getDishes({
                limit: requestItemsPerPage.value,
                offset: (requestPage.value - 1) * requestItemsPerPage.value
            });
            dishList.value = result;

        } catch (error) {
            throw error;
        } finally {
            loading.value = false;
        }
    }

    const addDish = async (dish: Dish) => {
        try {
            loading.value = true;
            const createdDish = await createDish(dish);
            dishList.value.push(createdDish); 
        } catch (error) {
            throw error;
        } finally {
            loading.value = false;
        }
    };

    const saveDish = async (dish: Dish) => {
        try {
            loading.value = true;
            await updateDish(dish);
            const index = dishList.value.findIndex((d) => d.id === dish.id);
            if (index !== -1) {
                dishList.value[index] = dish;
            }
        } catch (error) {
            throw error;
        } finally {
            loading.value = false;
        }
    };

    const removeDish = async (id: string) => {
        try {
            loading.value = true;
            await deleteDish(id);
            const index = dishList.value.findIndex((d) => d.id === id);
            if (index !== -1) {
                dishList.value.splice(index, 1);
            }
        } catch (error) {
            throw error;
        } finally {
            loading.value = false;
        }
    };

    return {
        dishList,
        loadDishes,
        addDish,
        saveDish,
        removeDish,
        loading
    }
});