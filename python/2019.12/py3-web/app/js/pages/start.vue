<template>
  <div class="start">
    <img @click="ctrlMusic" class="music" src="../../img/music.jpg" alt="" />
    <audio
      id="audioPlay"
      src="https://rechengparty-dist.oss-cn-chengdu.aliyuncs.com/cache/1.mp3"
      autoplay="autoplay"
    ></audio>
    <img class="banner" src="../../img/start-img.jpeg" alt="" />
    <div class="kanjia-ban">
      <img class="logo" src="/app_static/img/logo.jpg" alt="" />
      <div class="text">
        <span class="title">{{ store_name }}</span>
        <div>
          <div>
            原价：<span class="old-price">{{ old_price }}元</span>
          </div>
          <div>
            已砍价：<span class="sub-price">{{ sub_amount }}元</span>
          </div>
          <div>
            现价：<span class="new-price">{{ old_price - sub_amount }}元</span>
          </div>
        </div>

        <router-link to="/post">
          <img class="btn" src="../../img/kan-button.jpg" alt="" />
        </router-link>
      </div>
    </div>
    <h3 v-if="joiners.length > 0">助力排行榜</h3>
    <div class="rank">
      <ul class="list" v-for="item in joiners" :key="item">
        <li>
          <span>
            <img src="../../img/avatar.jpg" alt="" />
          </span>
          <span>
            {{ item.wx_user_name }}
          </span>
          <span> 已砍{{ item.help_amount }}元 </span>
        </li>
      </ul>
      <ul v-if="joiners.length == 0" class="note">
        <li></li>
        <li>1.点击分享砍价,获取海报图片</li>
        <li>2.将海报分享至朋友圈</li>
        <li>3.好友使用微信扫描海报上的二维码，参与砍价</li>
        <li>4.根据好友参与数量，您将获取50元到200元不等的优惠额度</li>
      </ul>
      <br />
      <div class="zuzhi">
        ©热橙派对 2014~2020
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      state: true,
      name: "",
      store_name: "",
      sub_amount: "",
      old_price: "",
      joiners: []
    };
  },
  mounted() {
    if (window.location.href.includes("scanning")) {
      this.$router.push({ path: "result" });
      return;
    }
    this.info();
  },
  methods: {
    ctrlMusic() {
      let audio = document.getElementById("audioPlay");
      if (this.state) {
        audio.pause();
      } else {
        audio.play();
      }
      this.state = !this.state;
    },
    info() {
      this.store_name = document.querySelector("#data").dataset.store_name;
      this.sub_amount = document.querySelector("#data").dataset.sub_amount;
      this.old_price = document.querySelector("#data").dataset.old_price;
      let temp = document
        .querySelector("#data")
        .dataset.joiners.replace(/\'/g, '"');
      temp = temp.replace(/None/g, "null");
      console.log(temp);
      this.joiners = JSON.parse(temp);
    }
  }
};
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
  .music {
    position: fixed;
    z-index: 10;
    top: pxToRem(30);
    right: pxToRem(30);
    width: pxToRem(200);
    cursor: pointer;
  }
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
  background: #fff;
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
      color: #fec300;
    }
    .old-price {
      text-decoration: line-through;
    }
    .sub-price {
    }
    .new-price {
      color: red;
    }
    .btn {
      position: absolute;
      right: pxToRem(30);
      bottom: pxToRem(50);
      width: pxToRem(210);
      cursor: pointer;
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
  .note {
    margin: pxToRem(-50) 0 0;
    padding: 0;
    width: pxToRem(720);
    li {
      list-style: none;
      color: #fff;
      text-align: left;
      &:first-child {
        position: relative;
        text-align: center;
        display: inline-block;
        width: 100%;
        border-top: 1px dotted #fff;
        &::before {
          position: absolute;
          content: "活动说明";
          top: pxToRem(-30);
          transform: translateX(-50%);
          background: #fec300;
          padding: 0 pxToRem(30) 0;
        }
      }
      &:nth-child(2) {
        margin-top: pxToRem(30);
      }
    }
  }
  .zuzhi {
    margin: pxToRem(40);
    color: #fff;
  }
}
</style>
