<template>
  <div>
    <div class="kanjia-ban">
      <img
        class="logo"
        src="/app_static/img/logo.jpg"
        alt=""
      >
      <div class="text">
        <span class="title">场地布置场地布置</span>
        <div>
          <div>原价：<span class="old-price">1700元</span></div>
          <div>已砍价：<span class="sub-price">500元</span></div>
          <div>现价：<span class="new-price">1200元</span></div>
        </div>
        <button>开始砍价</button>
      </div>
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
.kanjia-ban {
  margin: 0 auto;
  width: pxToRem(920);
  display: flex;
  background: #8be944;
  border-radius: pxToRem(50);
  position: relative;
  overflow: hidden;
  .logo {
    width: pxToRem(300);
    height: pxToRem(230);
    margin: 0.25rem;
  }
  .text {
    margin: pxToRem(25);
    font-size: pxToRem(40);
    color: #000000;
    .title {
      font-size: pxToRem(50);
      color: #298423;
    }
    .old-price {
      text-decoration: line-through;
    }
    .sub-price {
    }
    .new-price {
      color: red;
    }
    button {
      position: absolute;
      right: pxToRem(50);
      bottom: pxToRem(50);
    }
  }
}
</style>