import Vue from "vue"

import App from "./app.vue"

const spaApp = new Vue({
    el:"#app",
    render:c=c(App)
})

Vue.use({
    spaApp
})