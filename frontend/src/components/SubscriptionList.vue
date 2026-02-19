<template>
  <v-container fluid>
    <v-card variant="flat" border>
      <v-card-title class="pa-4 d-flex align-center text-secondary">
        Abonnements
        <v-spacer></v-spacer>
        <v-btn
          color="secondary"
          prepend-icon="mdi-plus"
          @click="showCreateDialog = true"
        >
          Abo hinzufügen
        </v-btn>
      </v-card-title>

      <v-data-table
        :headers="headers"
        :items="subscriptions"
        :loading="loading"
        hover
      >
        <template v-slot:item.SubscriptionBase.name="{ value }">
          <span class="font-weight-medium">{{ value }}</span>
        </template>

        <template v-slot:item.SubscriptionBase.tier="{ item }">
          <div class="py-2">
            <div class="text-subtitle-2">
              {{ item.SubscriptionBase.tier || "Standard" }}
            </div>
            <div
              v-if="item.SubscriptionBase.features?.length"
              class="text-caption text-grey"
            >
              {{ item.SubscriptionBase.features.join(", ") }}
            </div>
          </div>
        </template>

        <template v-slot:item.dates="{ item }">
          <div class="text-caption font-mono">
            {{ formatDate(item.SubscriptionBase.start_date) }} —
            <span
              :class="!item.SubscriptionBase.end_date ? 'text-success' : ''"
            >
              {{
                item.SubscriptionBase.end_date
                  ? formatDate(item.SubscriptionBase.end_date)
                  : "Unbefristet"
              }}
            </span>
          </div>
        </template>

        <template v-slot:item.SubscriptionBase.auto_renew="{ value }">
          <v-tooltip
            location="top"
            :text="
              value
                ? 'Automatische Verlängerung aktiv'
                : 'Manuelle Verlängerung'
            "
          >
            <template v-slot:activator="{ props }">
              <v-avatar
                :color="value ? 'success-lighten-5' : 'grey-lighten-4'"
                size="32"
                v-bind="props"
              >
                <v-icon :color="value ? 'success' : 'grey'" size="18">
                  {{ value ? "mdi-autorenew" : "mdi-calendar-remove" }}
                </v-icon>
              </v-avatar>
            </template>
          </v-tooltip>
        </template>

        <template v-slot:item.SubscriptionBase.status="{ value }">
          <v-chip
            :color="getStatusColor(value)"
            size="small"
            class="text-uppercase font-weight-bold"
            variant="flat"
          >
            {{ value }}
          </v-chip>
        </template>
      </v-data-table>
    </v-card>
    <CreateSubscriptionDialog
      v-model="showCreateDialog"
      @created="fetchSubscriptions"
    />
  </v-container>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { ListSubscriptionsCommand } from "../../wailsjs/go/SubscriptionCommand/subscriptionCommand";
import { domain } from "../../wailsjs/go/models";
import CreateSubscriptionDialog from "../components/dialogs/CreateSubscription.vue";

const subscriptions = ref<domain.SubscriptionResponse[]>([]);
const loading = ref(false);
const showCreateDialog = ref(false);

const headers = [
  { title: "Abonnement", key: "SubscriptionBase.name" },
  { title: "Plan & Features", key: "SubscriptionBase.tier" },
  { title: "Laufzeit", key: "dates", sortable: false },
  {
    title: "Auto-Renew",
    key: "SubscriptionBase.auto_renew",
    align: "center" as const,
  },
  { title: "Status", key: "SubscriptionBase.status", align: "center" as const },
] as const;

const fetchSubscriptions = async () => {
  loading.value = true;
  try {
    const result = await ListSubscriptionsCommand();
    if (result.status === 200 && result.data) {
      subscriptions.value = result.data;
    }
  } catch (err) {
    console.error("Failed to fetch subscriptions:", err);
  } finally {
    loading.value = false;
  }
};

// Helper to handle the string dates from Go
const formatDate = (dateStr: string) => {
  if (!dateStr) return "";
  return new Date(dateStr).toLocaleDateString("de-DE", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  });
};

const getStatusColor = (status: string) => {
  const s = status?.toLowerCase(); // normalize input
  if (s === "active") return "success";
  if (s === "expired") return "error";
  if (s === "pending") return "warning";
  if (s === "cancelled") return "grey-darken-1";
  return "grey";
};

onMounted(fetchSubscriptions);
</script>
