import { createRouter, createWebHashHistory, createWebHistory } from "vue-router";
import HelloWorldVue from "@/components/HelloWorld.vue";
import Mastered from "@/components/Mastered.vue";
import Vocabulary from "@/components/Vocabulary.vue";

const routes = [
    {
        path: '/',
        redirect: '/home',
    },
    {
      name: "home",
      path: "/home",
      component: HelloWorldVue,
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

export default createRouter({
  history: createWebHashHistory(),
  routes,
});
