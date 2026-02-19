<template>
  <v-container fluid>
    <v-row class="mb-2">
      <v-col cols="12">
        <v-card border flat class="pa-4 bg-grey-lighten-4">
          <div class="d-flex align-center">
            <div class="text-h6 font-weight-bold mr-6">Quick Actions</div>
            <v-btn
              color="primary"
              prepend-icon="mdi-domain-plus"
              class="mr-3"
              @click="showCreateCompanyDialog = true"
            >
              Kunde anlegen
            </v-btn>
            <v-btn
              color="secondary"
              prepend-icon="mdi-calendar-plus"
              @click="showCreateSubscriptionDialog = true"
            >
              Abo hinzufügen
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn variant="text" icon="mdi-refresh"></v-btn>
          </div>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col
        cols="12"
        sm="6"
        md="3"
        v-for="stat in mainStats"
        :key="stat.title"
      >
        <v-card border flat class="pa-4">
          <div class="d-flex align-center">
            <v-avatar :color="stat.bgColor" size="56" rounded>
              <v-icon
                :icon="stat.icon"
                :color="stat.iconColor"
                size="32"
              ></v-icon>
            </v-avatar>
            <div class="ms-4">
              <div class="text-overline font-weight-bold">{{ stat.title }}</div>
              <div class="text-h4 font-weight-bold">{{ stat.value }}</div>
            </div>
          </div>
        </v-card>
      </v-col>
    </v-row>

    <v-row class="mt-4">
      <v-col cols="12">
        <v-card border flat>
          <v-card-title class="text-subtitle-1 pb-0"
            >Verteilung nach Bereich</v-card-title
          >
          <v-card-text>
            <v-row dense>
              <v-col
                v-for="(count, area) in areaCounts"
                :key="area"
                cols="12"
                sm="4"
                md="2.4"
              >
                <v-card
                  border
                  flat
                  class="pa-3 text-center bg-blue-grey-lighten-5"
                >
                  <div
                    class="text-caption text-uppercase font-weight-black text-grey"
                  >
                    {{ area }}
                  </div>
                  <div class="text-h5 font-weight-bold">{{ count }}</div>
                </v-card>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-row class="mt-4">
      <v-col cols="12">
        <v-card border flat>
          <v-card-item class="bg-blue-grey-lighten-5">
            <v-card-title class="d-flex align-center">
              <v-icon
                icon="mdi-clock-fast"
                class="mr-2"
                color="warning"
              ></v-icon>
              Anstehende Verlängerungen
              <v-spacer></v-spacer>
              <v-btn variant="text" size="small" to="/subscriptions"
                >Alle Abos anzeigen</v-btn
              >
            </v-card-title>
          </v-card-item>
          <v-divider></v-divider>
          <v-table hover>
            <thead>
              <tr>
                <th class="text-left">Abonnement</th>
                <th class="text-left">Kunde</th>
                <th class="text-left">Laufzeit bis</th>
                <th class="text-left">Status</th>
                <th class="text-right">Aktionen</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="sub in upcomingRenewals" :key="sub.id">
                <td class="font-weight-bold">{{ sub.name }}</td>
                <td>{{ sub.companyName }}</td>
                <td>
                  <v-chip
                    size="small"
                    prepend-icon="mdi-calendar"
                    variant="text"
                  >
                    {{ sub.end_date }}
                  </v-chip>
                </td>
                <td>
                  <v-chip
                    :color="sub.auto_renew ? 'success' : 'orange'"
                    size="x-small"
                    label
                  >
                    {{ sub.auto_renew ? "AUTO-RENEW" : "MANUELL" }}
                  </v-chip>
                </td>
                <td class="text-right">
                  <v-btn
                    icon="mdi-eye"
                    variant="text"
                    size="small"
                    color="primary"
                  ></v-btn>
                </td>
              </tr>
            </tbody>
          </v-table>
        </v-card>
      </v-col>
    </v-row>
    <CreateCompany v-model="showCreateCompanyDialog" />
    <CreateSubscription v-model="showCreateSubscriptionDialog" />
  </v-container>
</template>

<script setup lang="ts">
import { ref } from "vue";
import CreateCompany from "../components/dialogs/CreateCompany.vue";
import CreateSubscription from "../components/dialogs/CreateSubscription.vue";

// --- State ---
const showCreateCompanyDialog = ref(false);
const showCreateSubscriptionDialog = ref(false);

// Statics for UI Mockup
const mainStats = ref([
  {
    title: "Aktive Abos",
    value: "342",
    icon: "mdi-sync",
    bgColor: "success-lighten-4",
    iconColor: "success",
  },
  {
    title: "Aktive Kunden",
    value: "89",
    icon: "mdi-domain",
    bgColor: "info-lighten-4",
    iconColor: "info",
  },
  {
    title: "Verlängerungen (30T)",
    value: "12",
    icon: "mdi-clock-alert",
    bgColor: "warning-lighten-4",
    iconColor: "warning",
  },
  {
    title: "Gesamtbestand",
    value: "102",
    icon: "mdi-database",
    bgColor: "purple-lighten-4",
    iconColor: "purple",
  },
]);

const areaCounts = ref({
  Microtech: 45,
  HaufeX360: 12,
  Software: 22,
  IT: 10,
  Consulting: 5,
});

const upcomingRenewals = ref([
  {
    id: 1,
    name: "M365 Business Premium",
    companyName: "Müller IT",
    end_date: "01.03.2026",
    auto_renew: true,
  },
  {
    id: 2,
    name: "Haufe API Connect",
    companyName: "Schmidt KG",
    end_date: "05.03.2026",
    auto_renew: false,
  },
  {
    id: 3,
    name: "Virenschutz MGD",
    companyName: "Bäckerei Sonnenschein",
    end_date: "12.03.2026",
    auto_renew: true,
  },
]);
</script>
