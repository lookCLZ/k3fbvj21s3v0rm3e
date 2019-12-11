<template>
  <div>
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

</style>