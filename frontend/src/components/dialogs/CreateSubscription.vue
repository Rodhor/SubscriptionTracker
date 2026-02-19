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
import { domain } from "../../../wailsjs/go/models";
import { ListCompaniesCommand } from "../../../wailsjs/go/CompanyCommand/companyCommand";
import { CreateSubscriptionCommand } from "../../../wailsjs/go/SubscriptionCommand/subscriptionCommand";
import { SubscriptionTier, SubscriptionStatus } from "@/types/domain";

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
const showLicense = ref(false);
const tempFeature = ref("");
const companies = ref<{ name: string; id: string }[]>([]);

const tierOptions = Object.values(SubscriptionTier) as any;
const statusOptions = Object.values(SubscriptionStatus) as any;

/**
 * Subscription Data Model (Wails domain class)
 */
const formData = ref(
  new domain.CreateSubscriptionRequest({
    company_id: "",
    name: "",
    license: "",
    tier: SubscriptionTier.Core,
    features: [],
    status: SubscriptionStatus.Active,
    start_date: new Date().toISOString().substring(0, 10),
    end_date: null,
    auto_renew: true,
  }),
);

/**
 * Form Field Actions
 */
const addFeature = () => {
  const val = tempFeature.value.trim();
  if (val) {
    if (!formData.value.features) {
      formData.value.features = [];
    }
    formData.value.features.push(val);
    tempFeature.value = "";
  }
};

const removeFeature = (index: number) => {
  formData.value.features?.splice(index, 1);
};

/**
 * Dialog Lifecycle Actions
 */
const resetForm = () => {
  formData.value = new domain.CreateSubscriptionRequest({
    company_id: "",
    name: "",
    license: "",
    tier: SubscriptionTier.Core,
    features: [],
    status: SubscriptionStatus.Active,
    start_date: new Date().toISOString().substring(0, 10),
    end_date: null,
    auto_renew: true,
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
    // Clone the form data so we don't mess up the UI binding
    const payload = { ...formData.value };

    // Append Time and Z (UTC) suffix to satisfy Go's time.Time parser
    if (payload.start_date && !payload.start_date.includes("T")) {
      payload.start_date = `${payload.start_date}T00:00:00Z`;
    }

    if (payload.end_date && !payload.end_date.includes("T")) {
      payload.end_date = `${payload.end_date}T00:00:00Z`;
    } else if (!payload.end_date) {
      // Ensure null is handled if the backend expects a pointer or omitempty
      payload.end_date = undefined;
    }

    const result = await CreateSubscriptionCommand(payload);

    if (result.status !== 201) {
      throw new Error(result.message || "Failed to create subscription");
    }

    emit("created");
    close();
  } catch (err) {
    console.error("Subscription creation error:", err);
  } finally {
    loading.value = false;
  }
};

/**
 * Initialization Logic
 */
onMounted(async () => {
  try {
    const result = await ListCompaniesCommand();

    if (result.status === 200 && Array.isArray(result.data)) {
      // Map the nested Go structure to a flat list for the dropdown
      companies.value = result.data.map((c: any) => ({
        // We look for 'name' inside CompanyBase, fallback to 'name'
        name: c.CompanyBase?.name || c.name || "Unbekannter Kunde",
        id: c.id,
      }));
    }
  } catch (err) {
    console.error("Initialization error (Company List):", err);
  }
});
</script>
