import { createApp } from "vue"
import { createPinia } from "pinia"
import "@picocss/pico/css/pico.min.css"

import App from "./App.vue"
import router from "./router"
import { HTTPError } from "ky"

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount("#app")
