const path = require('path')
const ExtractTextPlugin = require('extract-text-webpack-plugin')
const OptimizeCssAssetsPlugin = require('optimize-css-assets-webpack-plugin')
const Uglify = require('uglifyjs-webpack-plugin')
const CopyWebpackPlugin = require('copy-webpack-plugin')
const commonConfig = require('./webpack.common.js')

module.exports = [
  {
    devtool: 'cheap-module-source-map',
    entry: commonConfig.entries.jsEntries,
    output: {
      path: path.resolve(__dirname, '../../www/app_static'),
      filename: '[name].js'
    },
    module: {
      rules: [
        {
          test: /\.js$/,
          loader: 'babel-loader'
        },
        {
          test: /\.vue$/,
          loader: 'vue-loader'
        },
        {
          test: /\.css$/,
          use: ['style-loader', 'css-loader']
        },
        {
          test: /\.scss$/,
          use: [
            'style-loader',
            {
              loader: 'css-loader',
              options: {
                modules: true,
                importLoaders: 1,
                localIdentName: '[name]__[local]___[hash:base64:5]'
              }
            },
            'sass-loader'
          ]
        },
        {
          test: /\.(gif|png|woff|woff2|eot|ttf|svg|jpg|jpeg)$/,
          loader: 'url-loader',
          options: {
            limit: 8192,
            name: 'img/[name].[hash:8].[ext]',
            publicPath: '/app_static'
          }
        }
      ]
    },
    plugins: []
  },
  {
    devtool: 'cheap-module-source-map',
    entry: commonConfig.entries.cssEntries,
    output: {
      path: path.resolve(__dirname, '../../www/app_static'),
      filename: '[name].css'
    },
    module: {
      rules: [
        {
          test: /\.(css|scss)$/,
          use: ExtractTextPlugin.extract({
            fallback: 'style-loader',
            use: ['css-loader', 'sass-loader']
          })
        },
        {
          test: /\.(gif|png|woff|woff2|eot|ttf|svg|jpg|jpeg)$/,
          loader: 'url-loader',
          options: {
            limit: 10000,
            name: 'img/[name].[hash:8].[ext]',
            publicPath: '/static'
          }
        },
        {
          test: /\.(woff2?|eot|ttf|otf)$/,
          use: [
            {
              loader: 'url-loader',
              options: {
                limit: 5000
              }
            }
          ]
        }
      ]
    },
    plugins: [
      new ExtractTextPlugin('[name].css'),
      new CopyWebpackPlugin([
        {
          from: './img',
          to: '../www/app_static/img'
        }
      ])
    ]
  }
]

if (process.env.NODE_ENV === 'production') {
  // module.exports[0].plugins.push(new Uglify())
  module.exports[1].plugins.push(
    new OptimizeCssAssetsPlugin({
      assetNameRegExp: /\.css$/g,
      cssProcessor: require('cssnano'),
      cssProcessorPluginOptions: {
        preset: ['default', { discardComments: { removeAll: true } }]
      },
      canPrint: true
    })
  )
}
