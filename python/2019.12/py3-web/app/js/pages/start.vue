<template>
  <div class="start">
    <img
      class="banner"
      src="../../img/start-img.jpeg"
      alt=""
    >
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
    <h3>助力排行榜</h3>
    <div class="rank">
      <ul class="list">
        <li>
          <span>
            <img
              src="../../img/avatar.jpg"
              alt=""
            >
          </span>
          <span>
            浪倾城
          </span>
          <span>
            已砍50元
          </span>
        </li>
        <li>
          <span>
            <img
              src="../../img/avatar.jpg"
              alt=""
            >
          </span>
          <span>
            浪倾城
          </span>
          <span>
            已砍50元
          </span>
        </li>
        <li>
          <span>
            <img
              src="../../img/avatar.jpg"
              alt=""
            >
          </span>
          <span>
            浪倾城
          </span>
          <span>
            已砍50元
          </span>
        </li>
        <li>
          <span>
            <img
              src="../../img/avatar.jpg"
              alt=""
            >
          </span>
          <span>
            浪倾城
          </span>
          <span>
            已砍50元
          </span>
        </li>
        <li>
          <span>
            <img
              src="../../img/avatar.jpg"
              alt=""
            >
          </span>
          <span>
            浪倾城
          </span>
          <span>
            已砍50元
          </span>
        </li>
        <li>
          <span>
            <img
              src="../../img/avatar.jpg"
              alt=""
            >
          </span>
          <span>
            浪倾城
          </span>
          <span>
            已砍50元
          </span>
        </li>
        <li>
          <span>
            <img
              src="../../img/avatar.jpg"
              alt=""
            >
          </span>
          <span>
            浪倾城
          </span>
          <span>
            已砍50元
          </span>
        </li>
      </ul>
      <div class="zuzhi">
        @成都热橙文化传播公司
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
.start {
  position: relative;
  // background: url("../../img/start-bg.jpg");
  height: 100vh;
  background-size: contain;
  background-repeat: no-repeat;
  background-color: #fec300;
  overflow: scroll;

  .banner {
    position: fixed;
    width: 100%;
    z-index: 9;
  }
}
.kanjia-ban {
  z-index: 9;
  margin: 0 auto;
  width: pxToRem(860);
  display: flex;
  background: #c8dce6;
  border-radius: pxToRem(50);
  overflow: hidden;
  position: fixed;
  top: pxToRem(620);
  left: 50%;
  transform: translateX(-50%);
  .logo {
    width: pxToRem(250);
    height: pxToRem(190);
    margin-left: pxToRem(28);
    margin-top: pxToRem(46);
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
      right: pxToRem(30);
      bottom: pxToRem(50);
      background: red;
      border-radius: pxToRem(40);
      color: #fff;
      width: pxToRem(210);
      line-height: pxToRem(80);
      font-size: pxToRem(40);
    }
  }
}
h3 {
  position: fixed;
  text-align: center;
  width: 100%;
  z-index: 8;
  top: pxToRem(900);
  display: block;
  background: #fec300;
  margin: 0;
  padding: pxToRem(45) 0 pxToRem(30);
}
.rank {
  text-align: center;
  position: absolute;
  top: pxToRem(1010);
  height: calc(100% - pxToRem(910));
  left: 50%;
  transform: translateX(-50%);
  .list {
    width: pxToRem(860);
    margin: 0 auto;
    padding: 0;
    justify-content: space-around;
    li {
      list-style: none;
      height: pxToRem(160);
      border-bottom: 1px dotted red;
      &:last-child {
        border: 0;
      }
      img {
        width: pxToRem(120);
        border-radius: 50%;
        vertical-align: middle;
      }
      span {
        display: inline-block;
        vertical-align: middle;
        margin: 0 pxToRem(80);
        font-size: pxToRem(38);
        line-height: pxToRem(160);
        height: 100%;
        &:nth-child(1) {
          margin-left: 0;
        }
        &:nth-child(3) {
          margin-right: 0;
        }
      }
    }
  }
  .zuzhi {
    margin: pxToRem(40);
  }
}
</style>