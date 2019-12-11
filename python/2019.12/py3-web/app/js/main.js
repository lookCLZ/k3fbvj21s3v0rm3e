import Vue from "vue";

import App from "./app.vue";
import Start from "./pages/start.vue"
import VueRouter from "vue-router";

import 'lib-flexible/flexible'

const routes = [
  { path: "/start", component: Start },
];
const router = new VueRouter({
  routes // (缩写) 相当于 routes: routes
});

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
Vue.use(VueRouter);
