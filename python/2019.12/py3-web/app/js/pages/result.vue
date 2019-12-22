<template>
  <div>
    <div class="res">
      <img v-if="img20" src="../../img/20.jpg" alt="" />
      <img v-if="img30" src="../../img/30.jpg" alt="" />
      <img v-if="img50" src="../../img/50.jpg" alt="" />
      <img v-if="imgnull" src="../../img/null.jpg" alt="" />
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      img20: false,
      img30: false,
      img50: false,
      imgnull: false
    };
  },
  mounted() {
    let helpAmount = document.querySelector("#data").dataset.helpAmount + "";
    if (helpAmount == "20") {
      this.img20 = true;
    } else if (helpAmount == "30") {
      this.img30 = true;
    } else if (helpAmount == "50") {
      this.img50 = true;
    } else if (helpAmount == "null") {
      this.imgnull = true;
    }
  }
};
</script>
<style lang="scss" scoped>
.res {
  background: #ffbf00;
  img {
    width: 100%;
  }
}
</style>
