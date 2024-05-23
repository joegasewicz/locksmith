const path = require("path");

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
                    "style-loader",
                    "css-loader",
                    "sass-loader",
                ],
            },
        ],
    },
    resolve: {
        extensions: [".tsx", ".ts", ".js", ".scss", ".jpg", ".png"],
    },
    output: {
        filename: "bundle.js",
        path: path.resolve(__dirname, "dist/js"),
    },
};
