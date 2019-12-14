<template>
  <div class="page">
    <router-view></router-view>
    <Start />
  </div>
</template>

<script>
import axios from "axios";
import Start from "./pages/start.vue";
export default {
  components: {
    Start
  },
  data() {
    return {};
  },
  mounted() {
    console.log(this.getQueryString("code"));
    let url = "/wechart_user?code=" + this.getQueryString("code");
    axios.get(url).then(res => {
      res = res.data;
      window.localStorage.setItem("rechengparty_user", JSON.stringify(res));
    });
  },
  methods: {
    getQueryString(name) {
      let reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
      let r = window.location.search.substr(1).match(reg);
      if (r != null) {
        return unescape(r[2]);
      }
      return null;
    }
  }
};
</script>

<style>
.page {
  height: 100%;
}
</style>