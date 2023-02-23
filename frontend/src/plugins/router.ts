import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from "vue-router";
import Home from "@/components/Home.vue";
import Mastered from "@/components/Mastered.vue";
import Vocabulary from "@/components/Vocabulary.vue";
import { Ready } from "@/../wailsjs/go/relingo/Client";

const routes = [
  {
    path: "/",
    redirect: "/home",
  },
  {
    name: "home",
    path: "/home",
    component: Home,
  },
  {
    name: "words",
    path: "/words",
    component: Vocabulary,
  },
  {
    name: "mastered",
    path: "/mastered",
    component: Mastered,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

// router.beforeEach(async (to, from) => {
//   const isReady = await Ready();
//   console.log(1111, isReady, to.name);
  
//   if (!isReady && to.name !== "login") {
//     // 将用户重定向到登录页面
//     return { name: "login" };
//   }
// });

export default router;
