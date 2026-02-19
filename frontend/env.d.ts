/// <reference types="vite/client" />
/// <reference types="unplugin-vue-router/client" />
/// <reference types="vite-plugin-vue-layouts-next/client" />

declare module "*.vue" {
  import type { DefineComponent } from "vue";
  // The {} represents the props, the second {} represents the raw data, and 'any' is the rest.
  const component: DefineComponent<{}, {}, any>;
  export default component;
}
