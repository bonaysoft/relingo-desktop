/**
 * plugins/index.ts
 *
 * Automatically included in `./src/main.ts`
 */

// Plugins
import { loadFonts } from './webfontloader'
import router  from "./router";
import vuetify from './vuetify'

// Types
import type { App } from 'vue'
import dayjs from './dayjs'

export function registerPlugins (app: App) {
  loadFonts()
  app.use(vuetify)
  app.use(router)
  app.use(dayjs)
}
