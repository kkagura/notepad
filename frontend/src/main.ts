import { createApp } from "vue";
import router from "./router";
import App from "./App.vue";
import elementPlugin from "./vendors/elementPlugin";
import "@/assets/style/style.less";

createApp(App).use(router).use(elementPlugin).mount("#app");
