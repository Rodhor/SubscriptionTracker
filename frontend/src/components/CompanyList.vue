<template>
  <v-container fluid>
    <v-card variant="flat" border>
      <v-card-title class="pa-4 d-flex align-center">
        Kundenverwaltung
        <v-spacer></v-spacer>
        <v-btn color="primary" prepend-icon="mdi-plus">Kunde anlegen</v-btn>
      </v-card-title>

      <v-data-table :headers="headers" :items="companies" hover>
        <template v-slot:item.area="{ value }">
          <v-chip size="small" variant="tonal">{{ value }}</v-chip>
        </template>

        <template v-slot:item.is_active="{ value }">
          <v-chip
            :color="value ? 'success' : 'grey-lighten-1'"
            size="x-small"
            :variant="value ? 'flat' : 'tonal'"
          >
            {{ value ? "AKTIV" : "INAKTIV" }}
          </v-chip>
        </template>

        <template v-slot:item.actions>
          <v-btn icon="mdi-pencil-outline" variant="text" size="small"></v-btn>
        </template>
      </v-data-table>
    </v-card>
  </v-container>
</template>

<script setup lang="ts">
const headers = [
  { title: "Name", key: "name" },
  { title: "Gruppe", key: "area" },
  { title: "TANSS ID", key: "tannssid" },
  { title: "Status", key: "is_active", align: "center" },
  { title: "", key: "actions", align: "end", sortable: false },
] as const;

const companies = [
  {
    name: "Müller IT",
    area: "Managed Services",
    tannssid: "T-1002",
    is_active: true,
  },
  {
    name: "Bäckerei Schmidt",
    area: "Cloud Only",
    tannssid: "T-5501",
    is_active: false,
  },
];
</script>
