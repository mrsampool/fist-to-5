import {createRouter, createWebHistory} from "vue-router";

import StudentView from "../views/StudentView";
import TeacherView from "../views/TeacherView";
import SelectView from "../views/SelectView";
import Ask from "../components/Ask";
import Responses from "../components/Responses";

export const router = [
  { path: '/student', component: StudentView },
  {
    path: '/teacher',
    component: TeacherView,
    children: [
      { path: 'ask', component: Ask },
      { path: 'responses', component: Responses }
    ]
  },
  { path: '/', component: SelectView }
]

export default createRouter({
  history: createWebHistory(),
  routes: router,
})