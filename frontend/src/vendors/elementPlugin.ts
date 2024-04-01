import {
  ElDropdown,
  ElDropdownItem,
  ElDropdownMenu,
  ElIcon,
  ElInput,
} from "element-plus";
import "element-plus/theme-chalk/index.css";
import { App } from "vue";

const comps = [ElIcon, ElInput, ElDropdown, ElDropdownMenu, ElDropdownItem];

export default {
  install(app: App) {
    comps.forEach((comp) => {
      app.use(comp);
    });
  },
};
