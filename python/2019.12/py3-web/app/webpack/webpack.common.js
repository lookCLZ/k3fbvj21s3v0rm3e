const path = require('path')
const webpack = require('webpack')

let jsEntries = [
  "main.js"
]
const arrToObj = arr => {
  const obj = {}
  arr.map(x => (obj[`js/${x}`] = `./js/${x}`))
  return obj
}
jsEntries = arrToObj(jsEntries)
const cssEntries = { 'css/style': './css/style.css' }

const plugins = []

module.exports = {
  entries: {
    jsEntries,
    cssEntries
  },
  plugins
}