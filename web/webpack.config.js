const path = require("path");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const webpack = require("webpack");

module.exports = {
    entry: "./public/ts/index.ts",
    devtool: "inline-source-map",
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: "ts-loader",
                exclude: /node_modules/,
            },
            {
                test: /\.scss$/i,
                use: [
                    MiniCssExtractPlugin.loader,
                    "css-loader",
                    "sass-loader",
                ],
            },
        ],
    },
    resolve: {
        extensions: [".css", ".tsx", ".ts", ".js", ".scss", ".jpg", ".png"],
    },
    output: {
        filename: "[name].bundle.js",
        path: path.resolve(__dirname, "static"),
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: "[name].bundle.css",
        }),
        new webpack.DefinePlugin({
            "process.env.API_URL": JSON.stringify(process.env.API_URL) || JSON.stringify("http://localhost:8001/api/v1"),
        }),
        new webpack.ProvidePlugin({
            $: "jquery",
            bootstrap: "bootstrap",
        }),
    ],
    optimization: {
        splitChunks: {
            chunks: "all",
        }
    }
};