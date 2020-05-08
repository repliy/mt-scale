module.exports = {
  publicPath: '/',
  outputDir: 'dist',
  lintOnSave: true,
  devServer: {
    port: 1234,
    host: '0.0.0.0',
    disableHostCheck: true,
    proxy: {
      '/api': {
        target: 'http://192.168.199.183:8020',
        changeOrigin: true,
        pathRewrite: {
          '^/api': '/'
        }
      }
    }
  },
  transpileDependencies: [
    'vuetify'
  ]
}
