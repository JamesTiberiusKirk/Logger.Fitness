module.exports = {
    devServer: {
        proxy: {
            '^/api': {
                target: 'http://localhost:5000',
                changeOrigin: true,
                secure: false,
                pathRewrite: { '^/api': '/api' },
                logLevel: 'debug'
            },
        }
    }
};