<template>
  <v-container fluid>
    <v-card variant="flat" border>
      <v-card-title class="pa-4 d-flex align-center text-secondary">
        Abonnements
        <v-spacer></v-spacer>
        <v-btn color="secondary" prepend-icon="mdi-plus">Abo hinzufügen</v-btn>
      </v-card-title>

      <v-data-table :headers="headers" :items="subscriptions" hover>
        <template v-slot:item.tier="{ item }">
          <div class="py-2">
            <div class="font-weight-bold">{{ item.tier }}</div>
            <div class="text-caption text-grey-darken-1">
              {{ item.features.join(", ") }}
            </div>
          </div>
        </template>

        <template v-slot:item.dates="{ item }">
          <div class="text-body-2">
            {{ item.start_date }} — {{ item.end_date || "Unbefristet" }}
          </div>
        </template>

        <template v-slot:item.auto_renew="{ value }">
          <v-icon :color="value ? 'success' : 'grey-lighten-1'" size="small">
            {{ value ? "mdi-autorenew" : "mdi-arrow-right-bottom" }}
          </v-icon>
        </template>

        <template v-slot:item.status="{ value }">
          <v-chip
            :color="value === 'Active' ? 'success' : 'warning'"
            size="small"
            label
          >
            {{ value }}
          </v-chip>
        </template>
      </v-data-table>
    </v-card>
  </v-container>
</template>

<script setup lang="ts">
const headers = [
  { title: "Abonnement", key: "name" },
  { title: "Plan & Features", key: "tier" },
  { title: "Laufzeit", key: "dates" },
  { title: "Auto-Renew", key: "auto_renew", align: "center" },
  { title: "Status", key: "status", align: "center" },
] as const;

const subscriptions = [
  {
    name: "Microsoft 365 Business",
    tier: "Premium",
    features: ["Azure AD P1", "Intune", "Defender"],
    start_date: "01.01.2026",
    end_date: "31.12.2026",
    auto_renew: true,
    status: "Active",
  },
  {
    name: "ESET Endpoint Security",
    tier: "Core",
    features: ["Antivirus", "Firewall"],
    start_date: "15.02.2026",
    end_date: null,
    auto_renew: false,
    status: "Pending",
  },
];
</script>
