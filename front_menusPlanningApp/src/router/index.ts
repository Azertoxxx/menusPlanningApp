import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Dishes from '../views/Dishes.vue';

const routes = [
  { path: '/dishes', name: 'Dishes', component: Dishes },
  { path: '/planning', name: 'Planning', component: Home },
  { path: '/', component: Home }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;