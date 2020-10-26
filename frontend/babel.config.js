module.exports = {
  presets: [
    '@vue/cli-plugin-babel/preset',
  ],
  plugins: [
    ['module-resolver', {
      root: ['./src'],
      alias: {
        '@icons': './src/assets/icons',
        '@style': './src/assets/style',
        '@data': './src/assets/data',
      },
    }],
  ],
};
