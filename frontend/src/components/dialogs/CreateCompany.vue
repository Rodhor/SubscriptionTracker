<template>
  <v-dialog v-model="internalModel" max-width="600px" persistent>
    <v-card prepend-icon="mdi-domain-plus" title="Neuen Kunden anlegen">
      <v-card-text>
        <v-form ref="form" v-model="isFormValid">
          <v-row dense>
            <v-col cols="12">
              <v-text-field
                v-model="formData.name"
                label="Firmenname*"
                required
                :rules="[(v) => !!v || 'Name ist erforderlich']"
                variant="outlined"
              ></v-text-field>
            </v-col>

            <v-col cols="12" sm="6">
              <v-select
                v-model="formData.area"
                :items="areaOptions"
                label="Bereich*"
                required
                :rules="[(v) => !!v || 'Bereich wählen']"
                variant="outlined"
              ></v-select>
            </v-col>

            <v-col cols="12" sm="6">
              <v-text-field
                v-model="formData.tannssid"
                label="TANSS ID*"
                placeholder="T-XXXX"
                required
                :rules="[(v) => !!v || 'ID ist erforderlich']"
                variant="outlined"
              ></v-text-field>
            </v-col>

            <v-col cols="12">
              <v-switch
                v-model="formData.is_active"
                label="Kunde ist aktiv"
                color="success"
                hide-details
              ></v-switch>
            </v-col>

            <v-divider class="my-4"></v-divider>

            <v-col cols="12">
              <div class="text-subtitle-2 mb-2">Interne Notizen</div>
              <v-text-field
                v-model="tempComment"
                label="Neue Notiz hinzufügen"
                placeholder="Text eingeben und Enter drücken"
                variant="filled"
                density="compact"
                append-inner-icon="mdi-plus"
                @keyup.enter="addComment"
                @click:append-inner="addComment"
              ></v-text-field>

              <v-sheet
                v-if="formData.comments.length > 0"
                border
                rounded
                class="pa-2 mt-2 bg-grey-lighten-5"
                max-height="150"
                style="overflow-y: auto"
              >
                <v-chip
                  v-for="(comment, index) in formData.comments"
                  :key="index"
                  closable
                  size="small"
                  color="primary"
                  variant="flat"
                  class="ma-1"
                  @click:close="removeComment(index)"
                >
                  {{ comment }}
                </v-chip>
              </v-sheet>
              <div v-else class="text-caption text-grey text-center py-2">
                Noch keine Notizen hinzugefügt
              </div>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-divider></v-divider>

      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn variant="text" @click="close">Abbrechen</v-btn>
        <v-btn
          color="primary"
          variant="flat"
          :disabled="!isFormValid"
          :loading="loading"
          @click="submit"
        >
          Speichern
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";

const props = defineProps<{ modelValue: boolean }>();
const emit = defineEmits(["update:modelValue", "created"]);

const internalModel = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
});

const isFormValid = ref(false);
const loading = ref(false);
const tempComment = ref("");

const formData = ref({
  name: "",
  area: null,
  tannssid: "",
  is_active: true,
  comments: [] as string[],
});

const areaOptions = ["Microtech", "HaufeX360", "Software", "IT", "Consulting"];

// The logic to "stack" comments
const addComment = () => {
  const val = tempComment.value.trim();
  if (val) {
    // We push to the array, exactly what the Go slice expects
    formData.value.comments.push(val);
    tempComment.value = "";
  }
};

const removeComment = (index: number) => {
  formData.value.comments.splice(index, 1);
};

const close = () => {
  internalModel.value = false;
  formData.value = {
    name: "",
    area: null,
    tannssid: "",
    is_active: true,
    comments: [],
  };
};

const submit = async () => {
  loading.value = true;
  // This sends the full object including the comments array to your Go backend
  console.log("Sending CreateCompanyRequest to Wails:", formData.value);

  setTimeout(() => {
    loading.value = false;
    emit("created");
    close();
  }, 800);
};
</script>
