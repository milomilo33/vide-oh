const { defineConfig } = require('@vue/cli-service')
const fs = require('fs')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '^/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        ws: true,
        onProxyReq: function(request) {
          request.setHeader("origin", "http://localhost:8080");
        },
      },
    },
  }
})
