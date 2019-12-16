import Vue from "vue";

import Admin from "./admin.vue";
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';

import 'lib-flexible/flexible'
Vue.config.productionTip = false;

Vue.use(Antd);


new Vue({
  render: h => h(Admin)
}).$mount("#app");
