const path = require('path');

let entry = {
  server: {
    import: './src/entry.js',
    filename: '../yjs/dist/bundle.js'
  }
}

let outputPath = path.resolve(__dirname, '.')
let target = 'web'

module.exports = {
  entry: entry,
  module: {
    rules: [
      {
        test: /entry\.js?$/,
        use: [
          {
            loader: "expose-loader",
            options: {
              exposes: ["entry"],
            },
          },
        ],
        exclude: /node_modules/,
      },
    ],
  },
  mode: "production",
  resolve: {
    extensions: ['.js' ],
    fallback: { 
      "assert": require.resolve("assert") 
    }
  },
  target: target,
  output: {
    path: outputPath,
  }
};
