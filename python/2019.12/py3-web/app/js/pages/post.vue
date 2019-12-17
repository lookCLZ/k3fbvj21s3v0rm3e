<template>
  <div class="post">
    <div class="top">
      <img
        class="title"
        src="../../img/post-title.jpg"
        alt=""
      >
      <div class="qr-container">
        <img
          class="qr"
          src="../../img/qr.png"
          alt=""
        >
        <img
          class="code-wrap"
          src="../../img/code-wrap.jpg"
          alt=""
        >
      </div>

      <img
        class="btn"
        src="../../img/save-btn.jpg"
        alt=""
      >
    </div>
    <div class="bottom">
      <img
        class="arrow"
        src="../../img/arrow.jpg"
        alt=""
      >
      <img
        class="intro"
        src="../../img/intro.jpg"
        alt=""
      >
    </div>

  </div>
</template>
<script>
export default {};
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
    top: 16px;
    width: 168px;

    left: 50%;
    transform: translateX(-50%);
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
  bottom: pxToRem(-10);
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