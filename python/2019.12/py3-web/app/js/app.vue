<template>
  <div class="page">
    <router-view></router-view>
    <Start />
  </div>
</template>

<script>
import axios from "axios";
import sha1 from "js-sha1";
import Start from "./pages/start.vue";
export default {
  components: {
    Start
  },
  data() {
    return {
      wxInfo: {}
    };
  },
  mounted() {
    if (!window.localStorage.getItem("rechengparty_wx_db")) {
      this.saveUserInfo();
    } else {
      this.loadData();
    }
  },
  methods: {
    getQueryString(name) {
      let reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
      let r = window.location.search.substr(1).match(reg);
      if (r != null) {
        return unescape(r[2]);
      }
      return null;
    },
    saveUserInfo() {
      let url = "/wx/wechart_user?code=" + this.getQueryString("code");
      axios.get(url).then(res => {
        res = res.data;
        window.localStorage.setItem("rechengparty_wx_db", JSON.stringify(res));
        this.loadData();
        this.setJS_SDK();
      });
    },
    loadData() {
      let wxDb = window.localStorage.getItem("rechengparty_wx_db");
      if (wxDb == "") {
        alert("连接微信服务器失败，请退出当前页面，稍后再试");
        return;
      }
      this.wxInfo = JSON.parse(wxDb);
      console.log("wxDb", wxDb);
      console.log("wxInfo", this.wxInfo);
    },
    setJS_SDK() {
      let jsapi_ticket = this.wxInfo.r_for_js_sdk.ticket;
      let noncestr = +new Date() + "";
      let timestamp = +new Date() + "";
      let url = window.location.href.split("#")[0];
      let str = `jsapi_ticket=${jsapi_ticket}&noncestr=${noncestr}&timestamp=${timestamp}&url=${url}`;
      console.log("str", str);
      let sha1Str = sha1(str);
      console.log("sha1Str", sha1Str);
      wx.config({
        debug: false, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
        appId: "wx65b975e308c72245", // 必填，公众号的唯一标识
        timestamp: timestamp, // 必填，生成签名的时间戳
        nonceStr: noncestr, // 必填，生成签名的随机串
        signature: sha1Str, // 必填，签名
        jsApiList: ["playVoice"] // 必填，需要使用的JS接口列表
      });

      wx.ready(function() {
        var audio = document.getElementById("audioPlay");
        audio.play();
      });
    }
  }
};
</script>

<style>
.page {
  height: 100%;
}
</style>