module.exports = {
  devServer: {
    port: 9999,
    proxy: {
      '/api': {
          target: 'http://localhost:3000/',
          changeOrigin: true,
      },
  },
  },
  pwa: {
    workboxPluginMode: "InjectManifest",
    workboxOptions: {
      // swSrc is required in InjectManifest mode.
      swSrc: "src/service-worker.js"
    }
  }
};
