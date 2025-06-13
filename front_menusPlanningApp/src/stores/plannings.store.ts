import { defineStore } from "pinia";
import { ref } from "vue";
import type { Planning } from "../models/planning.model";
import { createPlanning, getPlannings, updatePlanning, deletePlanning } from "../api/planning.api";

export const usePlanningsStore = defineStore("plannings", () => {
        const planningList = ref<Planning[]>([]);
        const loading = ref(false);

    const loadPlanning = async () => {
        try {
            loading.value = true;
            const result = await getPlannings();
            planningList.value = result;
        } catch (error) {
            throw error;
        } finally {
            loading.value = false;
        }
    }

    const addPlanning = async (planning: Planning) => {
        try {
            loading.value = true;
            const createdPlanning = await createPlanning(planning);
            //reactive update of the list
            planningList.value = [...planningList.value, createdPlanning]; 
        } catch (error) {
            throw error;
        } finally {
            loading.value = false;
        }        
    };

    const savePlanning = async (planning: Planning) => {
        try {
            loading.value = true;
            await updatePlanning(planning);
            const index = planningList.value.findIndex((p) => p.id === planning.id);
            if (index !== -1) {
                planningList.value[index] = planning;
            }
        } catch (error) {
            throw error;
        } finally {
            loading.value = false;
        }   
    };

    const deletePlanning = async (id: string) => {
        try {
            loading.value = true;
            await deletePlanning(id);
            const index = planningList.value.findIndex((p) => p.id === id);
            if (index !== -1) {
                planningList.value.splice(index, 1);
            }

        } catch (error) {
            throw error;
        } finally {
            loading.value = false;
        }
    };

    return {
        planningList,
        loading,
        loadPlanning,
        addPlanning,
        savePlanning,
        deletePlanning
    };
});