import Vue from "vue";

import App from "./app.vue";
import Start from "./pages/start.vue";
import PostComponent from "./pages/post.vue";
import VueRouter from "vue-router";

import "lib-flexible/flexible";

// Vue.use(Start).use(Post);

const routes = [
  { path: "/post", name: "post", component: PostComponent },
  { path: "/w", name: "start", component: Start }
];
const router = new VueRouter({
  routes // (缩写) 相当于 routes: routes
});
Vue.use(VueRouter);
new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
