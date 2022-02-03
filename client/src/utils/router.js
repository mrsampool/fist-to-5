import StudentView from "../views/StudentView";
import {createRouter, createWebHistory} from "vue-router";
import SelectView from "../views/SelectView";

export const router = [
  { path: '/student', component: StudentView },
  { path: '/', component: SelectView }
]

export default createRouter({
  history: createWebHistory(),
  routes: router,
})