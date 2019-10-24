const path = require('path')
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = function (env, args) {
    let config = {
        entry: {
            index: './src/index.tsx',
        },
        output: {
            filename: 'bundle.js',
            path: path.resolve(__dirname, 'build')
        },
        plugins: [
            new CopyWebpackPlugin([
                'public'
            ])
        ],
        devServer: {
            contentBase: './dist',
            port: 4000,
        },
        module: {
            rules: [
                {
                    test: /\.js$/,
                    exclude: /node_modules/,
                    loader: "babel-loader"
                },
                {
                    test: /\.tsx?$/,
                    use: 'ts-loader',
                    exclude: /node_modules/
                },
                {
                    test: /\.css$/,
                    exclude: /node_modules/,
                    use: [
                        'style-loader',
                        'css-loader'
                    ]
                }
            ]
        },
        resolve: {
            extensions: ['.js', '.jsx', '.ts', '.tsx'],
        }
    }
    if (env && env.prod) {
        config.mode = 'production'
    } else {
        config.devtool = 'inline-source-map'
        config.mode = 'development'
    }

    return config
}