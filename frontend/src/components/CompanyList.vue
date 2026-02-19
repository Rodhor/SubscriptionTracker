<template>
  <v-container fluid>
    <v-card variant="flat" border>
      <v-card-title class="pa-4 d-flex align-center">
        Kundenverwaltung
        <v-spacer></v-spacer>
        <v-btn
          color="primary"
          prepend-icon="mdi-plus"
          @click="showCreateDialog = true"
        >
          Kunde anlegen
        </v-btn>
      </v-card-title>

      <v-data-table
        :headers="headers"
        :items="companies"
        :loading="loading"
        hover
      >
        <template v-slot:item.is_active="{ value }">
          <v-chip
            :color="value ? 'success' : 'error'"
            size="small"
            variant="tonal"
          >
            {{ value ? "Aktiv" : "Inaktiv" }}
          </v-chip>
        </template>

        <template v-slot:item.created_at="{ value }">
          {{ value ? new Date(value).toLocaleDateString("de-DE") : "-" }}
        </template>

        <template v-slot:item.actions="{ item }">
          <v-btn
            icon="mdi-pencil"
            variant="text"
            size="small"
            @click="editCompany(item)"
          ></v-btn>
          <v-btn
            icon="mdi-delete"
            variant="text"
            size="small"
            color="error"
            @click="deleteCompany(item)"
          ></v-btn>
        </template>
      </v-data-table>
    </v-card>

    <CreateCompanyDialog v-model="showCreateDialog" @created="fetchCompanies" />
  </v-container>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { ListCompaniesCommand } from "../../wailsjs/go/CompanyCommand/companyCommand";
import { domain } from "../../wailsjs/go/models";
import CreateCompanyDialog from "../components/dialogs/CreateCompany.vue";

const companies = ref<domain.CompanyResponse[]>([]);
const loading = ref(false);
const showCreateDialog = ref(false);

const headers = [
  { title: "Name", key: "CompanyBase.name" }, // Note: If nested via Go inline, use dot notation or flatten
  { title: "Gruppe", key: "CompanyBase.area" },
  { title: "TANSS ID", key: "CompanyBase.tannssid" },
  { title: "Status", key: "CompanyBase.is_active", align: "center" as const },
  { title: "", key: "actions", align: "end" as const, sortable: false },
];

const fetchCompanies = async () => {
  loading.value = true;
  try {
    const result = await ListCompaniesCommand();
    // result.data is now castable to the domain models Wails generated
    if (result.status === 200 && result.data) {
      companies.value = result.data;
    }
  } catch (err) {
    console.error("Fetch Error:", err);
  } finally {
    loading.value = false;
  }
};

const editCompany = (item: domain.CompanyResponse) => {
  console.log("Edit:", item.id);
};

const deleteCompany = (item: domain.CompanyResponse) => {
  console.log("Delete:", item.id);
};

onMounted(fetchCompanies);
</script>
