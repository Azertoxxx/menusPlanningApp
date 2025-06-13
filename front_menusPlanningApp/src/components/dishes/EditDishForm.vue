  <script setup lang="ts">
  import { reactive, watch, ref } from "vue";
  import type { Dish } from "../../models/dish.model";
  import { useDishesStore } from "../../stores/dishes.store";

  const props = defineProps<{
    modelValue: boolean;
    dish: Dish | null;
  }>();

  const emit = defineEmits<{
    (e: "update:modelValue", value: boolean): void;
    (e: "submitted"): void;
  }>();

  const dishesStore = useDishesStore();

  const form = reactive<Dish>({
    id: "",
    name: "",
    description: "",
    ingredients: "",
    imageUrl: "",
  });

  watch(() => props.modelValue ,(val) => {
      if (val && props.dish) {
        Object.assign(form, props.dish); 
      }
    },
    { immediate: true }
  );

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
      await dishesStore.saveDish(form);
      emit("submitted");
      emit("update:modelValue", false); 
    } catch (error) {
      console.error("Failed to update dish", error);
    }
  };

  const closeDialog = () => {
    emit("update:modelValue", false);
  };
  </script>

  <template>
    <v-dialog v-model="props.modelValue" max-width="600">
      <v-card>
        <v-card-title>Modifier le plat</v-card-title>
        <v-card-text>
          <v-form @submit.prevent="submitForm">
            <v-row>
              <v-col cols="12">
                <v-text-field v-model="form.name" label="Nom du plat" required />
              </v-col>
            </v-row>

            <v-row>
              <v-col cols="12">
                <v-textarea v-model="form.description" label="Description" />
              </v-col>
            </v-row>

            <v-row>
              <v-col cols="12">
                <v-textarea v-model="form.ingredients" label="Ingrédients" required />
              </v-col>
            </v-row>

            <v-row>
              <v-col cols="12">
                <v-file-input
                  label="Image du plat"
                  accept="image/*"
                  :multiple="false"
                  v-model="imageFile"
                />
              </v-col>
           </v-row>

            <v-card-actions class="justify-end">
              <v-btn text @click="closeDialog">Annuler</v-btn>
              <v-btn color="primary" type="submit">Enregistrer</v-btn>
            </v-card-actions>
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
  </template>
