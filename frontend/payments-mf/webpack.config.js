// payments-mf/webpack.config.js
const HtmlWebpackPlugin        = require('html-webpack-plugin');
const { ModuleFederationPlugin } = require('webpack').container;

module.exports = (_, argv) => ({
  entry : './src/bootstrap.js',
  mode  : argv.mode || 'development',

  devServer : {
    port              : 3004,
    historyApiFallback: true,
    headers           : {
      'Access-Control-Allow-Origin' : '*',
      'Cross-Origin-Resource-Policy': 'cross-origin',
    },
  },

  optimization : { splitChunks: false, runtimeChunk: false },

  output : {
    publicPath         : 'auto',
    clean              : true,
    crossOriginLoading : 'anonymous',
  },

  resolve : { extensions: ['.js', '.jsx'] },

  module  : {
    rules : [
      { test: /\.css$/i, use: ['style-loader', 'css-loader'] },
      {
        test   : /\.jsx?$/,
        exclude: /node_modules/,
        use    : { loader: 'babel-loader', options: { presets: ['@babel/preset-react'] } },
      },
    ],
  },

  plugins : [
    new HtmlWebpackPlugin({ template: './public/index.html' }),
    new ModuleFederationPlugin({
      name    : 'paymentsMf',
      filename: 'remoteEntry.js',
      exposes : { './App': './src/App' },
      shared  : {
        react             : { singleton: true, eager: true, requiredVersion: false },
        'react-dom'       : { singleton: true, eager: true, requiredVersion: false },
        'react-router-dom': { singleton: true, eager: true, requiredVersion: false },
      },
    }),
  ],
});
