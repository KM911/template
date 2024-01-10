import { createApp } from "vue";
import App from "./App.vue";
import NutUI from "@nutui/nutui";
import "@nutui/nutui/dist/style.css";

const app = createApp(App)
import { Calendar } from '@nutui/nutui';

app.use(NutUI)
app.use(Calendar);
app.mount("#app");
