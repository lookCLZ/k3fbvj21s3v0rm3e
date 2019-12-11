<template>
  <div class="page">
      <router-view></router-view>
      <Start/>
  </div>
</template>

<script>
import Start from "./pages/start.vue"
export default {
  components:{
    Start
  }
}
</script>

<style>
.page{
  height:100%;
}
</style>