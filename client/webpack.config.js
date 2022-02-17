const path = require('path');
const { VueLoaderPlugin } = require('vue-loader');

module.exports = (env) => ({

    entry: path.resolve(__dirname, 'src', 'main.js'),
    output: {
        path: path.resolve(__dirname, 'public'),
        filename: 'bundle.js',
    },
    resolve: {
        extensions: ['.vue', '.js'],
    },
    mode: 'development',
    devServer: {
        client: {
            logging: 'log',
            overlay: {
                errors: true,
                warnings: false,
            },
            progress: true,
        },
        historyApiFallback: { index: 'index.html' },
        open: true,
        proxy: {
            '/api/*': {
                target: 'http://localhost:8080',
                changeOrigin: true,
                secure: false,
            }
        },
        static: {
            directory: path.join(__dirname, 'public'),
        },
    },

    module: {
        rules: [
            {
                test: /\.vue$/,
                include: path.resolve(__dirname, 'src'),
                exclude: /node_modules/,
                loader: 'vue-loader'
            },
            {
                test: /\.js$/,
                loader: 'babel-loader'
            },
            {
                test: /\.css$/,
                use: [
                    'vue-style-loader',
                    'css-loader'
                ]
            }
        ],
    },
    plugins: [
      new VueLoaderPlugin(),
    ]
});