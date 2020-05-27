module.exports = {
  publicPath: '/',
  outputDir: 'dist',
  lintOnSave: true,
  devServer: {
    port: 1234,
    host: '0.0.0.0',
    disableHostCheck: true,
    proxy: {
      '/': {
        target: 'http://localhost:8083',
        changeOrigin: true,
        pathRewrite: {
          '^/': '/'
        }
      }
    }
  },
  transpileDependencies: [
    'vuetify'
  ]
}
