import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    redirect: "/doc/page",
  },
  {
    path: "/doc/page",
    component: () => import("@/views/doc/DocPage.vue"),
    name: "DocPage",
    meta: {
      title: "文档列表",
      keepAlive: true,
    },
  },
];
const router = createRouter({
  history: createWebHistory("/"),
  routes,
});

export default router;
