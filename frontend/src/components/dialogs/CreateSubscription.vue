<template>
  <v-dialog v-model="internalModel" max-width="700px" persistent>
    <v-card prepend-icon="mdi-calendar-plus" title="Neues Abonnement anlegen">
      <v-card-text>
        <v-form ref="form" v-model="isFormValid">
          <v-row dense>
            <v-col cols="12">
              <v-autocomplete
                v-model="formData.company_id"
                :items="companies"
                item-title="name"
                item-value="id"
                label="Zugehöriger Kunde*"
                required
                :rules="[(v) => !!v || 'Bitte einen Kunden auswählen']"
                variant="outlined"
                prepend-inner-icon="mdi-domain"
              ></v-autocomplete>
            </v-col>

            <v-col cols="12">
              <v-text-field
                v-model="formData.name"
                label="Abonnement Name*"
                placeholder="z.B. Microsoft 365 Business"
                required
                :rules="[(v) => !!v || 'Name ist erforderlich']"
                variant="outlined"
              ></v-text-field>
            </v-col>

            <v-col cols="12" class="mb-3">
              <v-text-field
                v-model="formData.license"
                label="Lizenzschlüssel*"
                placeholder="Fügen Sie hier den vollständigen Schlüssel ein..."
                required
                variant="outlined"
                persistent-placeholder
                class="font-monospace"
                :append-inner-icon="showLicense ? 'mdi-eye' : 'mdi-eye-off'"
                :type="showLicense ? 'text' : 'password'"
                @click:append-inner="showLicense = !showLicense"
              >
                <template v-slot:details>
                  <div class="text-caption d-flex justify-space-between w-100">
                    <span>Länge: {{ formData.license.length }} Zeichen</span>
                    <span v-if="formData.license.length > 50" class="text-grey"
                      >Langer Schlüssel erkannt</span
                    >
                  </div>
                </template>
              </v-text-field>
            </v-col>

            <v-col cols="12" sm="6">
              <v-select
                v-model="formData.tier"
                :items="tierOptions"
                label="Subscription Tier*"
                variant="outlined"
              ></v-select>
            </v-col>

            <v-col cols="12" sm="6">
              <v-select
                v-model="formData.status"
                :items="statusOptions"
                label="Status*"
                variant="outlined"
              ></v-select>
            </v-col>

            <v-col cols="12" sm="6">
              <v-text-field
                v-model="formData.start_date"
                label="Startdatum"
                type="date"
                variant="outlined"
              ></v-text-field>
            </v-col>

            <v-col cols="12" sm="6">
              <v-text-field
                v-model="formData.end_date"
                label="Enddatum (Optional)"
                type="date"
                variant="outlined"
                clearable
              ></v-text-field>
            </v-col>

            <v-col cols="12">
              <v-text-field
                v-model="tempFeature"
                label="Features hinzufügen"
                placeholder="Feature eingeben und Enter drücken"
                variant="filled"
                density="compact"
                append-inner-icon="mdi-plus"
                @keyup.enter="addFeature"
                @click:append-inner="addFeature"
              ></v-text-field>

              <v-chip-group column>
                <v-chip
                  v-for="(feature, index) in formData.features"
                  :key="index"
                  closable
                  size="small"
                  color="secondary"
                  @click:close="removeFeature(index)"
                >
                  {{ feature }}
                </v-chip>
              </v-chip-group>
            </v-col>

            <v-col cols="12">
              <v-switch
                v-model="formData.auto_renew"
                label="Automatische Verlängerung (Auto-Renew)"
                color="secondary"
                hide-details
              ></v-switch>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>

      <v-divider></v-divider>

      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn variant="text" @click="close">Abbrechen</v-btn>
        <v-btn
          color="secondary"
          variant="flat"
          :disabled="!isFormValid"
          :loading="loading"
          @click="submit"
        >
          Abo Speichern
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";

const props = defineProps<{ modelValue: boolean }>();
const emit = defineEmits(["update:modelValue", "created"]);

const internalModel = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
});

// State
const isFormValid = ref(false);
const showLicense = ref(false);
const loading = ref(false);
const tempFeature = ref("");
const companies = ref<{ name: string; id: string }[]>([]); // To be filled from backend

// Data structure matches CreateSubscriptionRequest
const formData = ref({
  company_id: null,
  name: "",
  license: "",
  tier: "Core",
  features: [] as string[],
  status: "Active",
  start_date: new Date().toISOString().substr(0, 10),
  end_date: null,
  auto_renew: true,
});

// Enums from your Go domain
const tierOptions = ["Core", "Enterprise", "Enterprise API"];
const statusOptions = ["Active", "Inactive", "Cancelled", "Pending", "Expired"];

const addFeature = () => {
  if (tempFeature.value.trim()) {
    formData.value.features.push(tempFeature.value.trim());
    tempFeature.value = "";
  }
};

const removeFeature = (index: number) => {
  formData.value.features.splice(index, 1);
};

const close = () => {
  internalModel.value = false;
  // Reset logic...
};

const submit = async () => {
  loading.value = true;
  console.log("Sending CreateSubscriptionRequest:", formData.value);
  // Call Wails backend here
  setTimeout(() => {
    loading.value = false;
    emit("created");
    close();
  }, 1000);
};

// Lifecycle: Fetch companies so we can populate the dropdown
onMounted(async () => {
  // Mock data - in reality: companies.value = await ListCompaniesCommand()
  companies.value = [
    { name: "Müller IT", id: "uuid-1" },
    { name: "Schmidt KG", id: "uuid-2" },
  ];
});
</script>
