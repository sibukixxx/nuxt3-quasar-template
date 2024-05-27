
import { defineConfig } from 'histoire';
import { HstVue } from '@histoire/plugin-vue';
import vue from '@vitejs/plugin-vue'
import { transformAssetUrls } from '@quasar/vite-plugin'

export default defineConfig({
  setupFile: '/src/histoire.setup.ts',
  plugins: [
    HstVue(),
  ],
  vite: {
    plugins: [
      vue({
        template: { transformAssetUrls }
      }),
    ]
  },
});