<template>
  <div class="post">
    <div class="top">
      <img class="title" src="../../img/post-title.jpg" alt="" />
      <div class="qr-container">
        <div id="qrcode" class="qr"></div>
        <img class="code-wrap" src="../../img/code-wrap.jpg" alt="" />
      </div>

      <img class="btn" src="../../img/save-btn.jpg" alt="" />
    </div>
    <div class="bottom">
      <img class="arrow" src="../../img/arrow.jpg" alt="" />
      <img class="intro" src="../../img/intro.jpg" alt="" />
    </div>
  </div>
</template>
<script>
import QRCode from "qrcodejs2";
export default {
  data() {
    return {};
  },
  mounted() {
    this.qrcode();
    this.setCSS();
    // setTimeout(() => {
    //   window.GetImg(".post");
    // }, 5000);
  },
  methods: {
    qrcode() {
      let pwdCode = document.querySelector("#data").dataset.pwdCode;
      // pwdCode = pwdCode.slice(0, 6);
      let qrcode = new QRCode("qrcode", {
        width: 120,
        height: 120,
        text: `https://www.rechengparty.com/wx/scanning/${pwdCode}`, // 二维码地址
        colorDark: "#ffbf00",
        colorLight: "#fff"
      });
    },
    setCSS() {
      document.querySelector("#qrcode img").setAttribute(
        "style",
        `position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);`
      );
    }
  }
};
</script>
<style lang="scss" scoped>
@function pxToRem($num) {
  @return ($num/100) * 1rem;
}
.post {
  height: 100vh;
  background: #ffbf00;
}
.top {
  background: #ffbf00;
  text-align: center;
  .title {
    width: pxToRem(860);
  }
  .qr-container {
    position: relative;
  }
  .qr {
    position: absolute;
    width: pxToRem(470);
    height: pxToRem(440);
    text-align: center;
    left: 50%;
    top: pxToRem(40);
    transform: translateX(-50%);
    background: #fff;
  }
  .qr-img {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
  .code-wrap {
    width: pxToRem(860);
  }
  .btn {
    width: pxToRem(360);
  }
}
.bottom {
  position: absolute;
  bottom: pxToRem(-18);
  text-align: center;
  .arrow {
    width: pxToRem(100);
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
  }
  .intro {
    width: 100%;
  }
}
</style>
