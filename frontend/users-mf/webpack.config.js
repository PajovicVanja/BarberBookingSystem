const HtmlWebpackPlugin = require('html-webpack-plugin');
const { ModuleFederationPlugin } = require('webpack').container;
const path = require('path');

module.exports = (_, argv) => ({
  entry: './src/bootstrap.js',
  mode: argv.mode || 'development',
  devServer: { port: 3002, historyApiFallback: true },
  output: { publicPath: 'auto', clean: true },
  resolve: { extensions: ['.js', '.jsx'] },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: { presets: ['@babel/preset-react'] }
        }
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({ template: './public/index.html' }),
    new ModuleFederationPlugin({
      name: 'usersMf',
      filename: 'remoteEntry.js',
      exposes: { './App': './src/App' },
      shared: {
        react:      { singleton: true, eager: true, requiredVersion: false },
        'react-dom':{ singleton: true, eager: true, requiredVersion: false },
        'react-router-dom': { singleton: true, eager: true, requiredVersion: false }
      }
    })
  ]
});
