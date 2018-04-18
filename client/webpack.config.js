
require('dotenv').config({ path: '../.env' });
const htmlPlugin = require('html-webpack-plugin');
const path = require('path');
const webpack = require('webpack');

const directory = path.resolve(__dirname, 'dist');

module.exports = {
  devServer: {
    contentBase: directory,
    hot: true
  },
  devtool: 'inline-source-map',
  entry: ['whatwg-fetch', './src/index.js'],
  mode: 'development',
  module: {
    rules: [
      {
        exclude: /node_modules/,
        loader: "babel-loader",
        query: {
          plugins: [
            'transform-class-properties',
            'transform-object-rest-spread'
          ],
          presets: ['env', 'react']
        },
        test: /\.jsx?$/
      }
    ]
  },
  output: {
    filename: 'bundle.js',
    path: directory,
    publicPath: '/'
  },
  plugins: [
    new webpack.DefinePlugin({
      'API_PORT': process.env.DASH_PORT
    }),
    new webpack.NamedModulesPlugin(),
    new webpack.HotModuleReplacementPlugin(),
    new htmlPlugin({
      filename: 'index.html',
      hash: true,
      inject: 'body',
      template: './src/template.html',
      title: 'Bulwark Dashboard',
      xhtml: true
    })
  ],
  resolve: {
    extensions: ['.js', '.jsx']
  }
};
