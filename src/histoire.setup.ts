import {createApp} from 'vue'
import type {App} from 'vue'
import {Quasar, Platform} from 'quasar'
import '@quasar/extras/material-icons/material-icons.css'
import 'quasar/src/css/index.sass'

import {defineSetupVue3} from '@histoire/plugin-vue'

export const setupVue3 = defineSetupVue3(({app, story, variant}) => {
  app.use(Quasar, {})
  app.config.globalProperties.$q.platform = Platform
})

// export default function setupVue({ app }: { app: ReturnType<typeof createApp> }) {
//   app.use(Quasar, {
//     plugins: {}, // Import Quasar plugins if needed
//   })
//   app.config.globalProperties.$q.platform = Platform
// }

// export default function setupVue({ app }: { app: App }) {
//   app.use(Quasar, {
//     plugins: {}, // Import Quasar plugins if needed
//   })
//
//   // Make sure platform detection is available
//   app.config.globalProperties.$q.platform = Platform
// }