import Vue from "vue";

import App from "./app.vue";
import Start from "./pages/start.vue";
import PostComponent from "./pages/post.vue";
import VueRouter from "vue-router";

import "lib-flexible/flexible";

// Vue.use(Start).use(Post);

const routes = [
  { path: "/post", name: "post", component: PostComponent },
  { path: "/", name: "start", component: Start }
];
const router = new VueRouter({
  mode: "history",
  routes 
});
Vue.use(VueRouter);
new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
