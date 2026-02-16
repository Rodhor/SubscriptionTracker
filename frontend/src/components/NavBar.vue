<script setup lang="ts">
// No extra logic needed for basic routing,
// but you can add current route detection here later.
</script>

<template>
  <v-navigation-drawer expand-on-hover permanent rail>
    <v-list>
      <v-list-item prepend-icon="mdi-view-dashboard" title="Dashboard" to="/" />
    </v-list>

    <v-divider />

    <v-list density="compact" nav>
      <v-list-item
        prepend-icon="mdi-calendar-sync"
        title="Abos"
        to="/subscriptions"
        value="subscriptions"
      />

      <v-list-item
        prepend-icon="mdi-domain"
        title="Kunden"
        to="/companies"
        value="companies"
      />
    </v-list>

    <template v-slot:append>
      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-cog"
          title="Einstellungen"
          to="/settings"
        />
      </v-list>
    </template>
  </v-navigation-drawer>
</template>
