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
                v-if="formData.comments && formData.comments.length > 0"
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
import { domain } from "../../../wailsjs/go/models";
import { CreateCompanyCommand } from "../../../wailsjs/go/CompanyCommand/companyCommand";
import { ProductArea } from "@/types/domain";

/**
 * Component Props and Emits
 */
const props = defineProps<{ modelValue: boolean }>();
const emit = defineEmits(["update:modelValue", "created"]);

/**
 * Internal Dialog Visibility Control
 */
const internalModel = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
});

/**
 * Form State and Configuration
 */
const isFormValid = ref(false);
const loading = ref(false);
const tempComment = ref("");

// Cast area options to bypass strict Enum-to-String validation mismatch
const areaOptions = Object.values(ProductArea) as any;

/**
 * Company Data Model (Wails domain class)
 * Initialized with default values for CreateCompanyRequest
 */
const formData = ref(
  new domain.CreateCompanyRequest({
    name: "",
    area: ProductArea.Software,
    tannssid: "",
    is_active: true,
    comments: [],
  }),
);

/**
 * Internal Note Management
 */
const addComment = () => {
  const val = tempComment.value.trim();
  if (val) {
    if (!formData.value.comments) {
      formData.value.comments = [];
    }
    formData.value.comments.push(val);
    tempComment.value = "";
  }
};

const removeComment = (index: number) => {
  formData.value.comments?.splice(index, 1);
};

/**
 * Dialog Lifecycle Actions
 */
const resetForm = () => {
  formData.value = new domain.CreateCompanyRequest({
    name: "",
    area: ProductArea.Software,
    tannssid: "",
    is_active: true,
    comments: [],
  });
};

const close = () => {
  internalModel.value = false;
  resetForm();
};

/**
 * Backend Operations
 */
const submit = async () => {
  if (!isFormValid.value) return;

  loading.value = true;
  try {
    // Send CreateCompanyRequest instance to Go backend
    const result = await CreateCompanyCommand(formData.value);

    // Validate backend response status (expecting 201 Created)
    if (result.status !== 201) {
      throw new Error(result.message || "Fehler beim Erstellen des Kunden");
    }

    // Refresh parent state and close dialog
    emit("created");
    close();
  } catch (err) {
    console.error("Company creation error:", err);
  } finally {
    loading.value = false;
  }
};
</script>
