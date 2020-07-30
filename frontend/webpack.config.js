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
                    test: /\.(js|jsx)$/,
                    exclude: /node_modules/,
                    loader: "babel-loader"
                },
                {
                    test: /\.tsx?$/,
                    use: 'ts-loader',
                    exclude: /node_modules/
                },
                {
                    test: /\.(png|jpg|gif|ttf|eot|woff|woff2)$/i,
                    use: [
                      {
                        loader: 'url-loader',
                        options: { limit: 8192 }
                      }
                    ]
                },
                {
                    test: /\.css$/,
                    //exclude: /node_modules/,
                    use: [
                        'style-loader',
                        'css-loader'
                    ]
                },
                {
                    test: /\.(png|jpg|gif|svg)$/,
                    exclude: /node_modules/,
                    use: [
                        'script-loader'
                    ]
                },
                {
                    test: /\.less$/,
                    use: [
                        "less-loader" // compiles Less to CSS
                    ]
                },
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