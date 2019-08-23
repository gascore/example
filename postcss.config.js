module.exports = ctx => ({
    map: ctx.env !== 'production',
    plugins: {
        'postcss-easy-import': { extensions: ['.css', '.pcss', '.scss'] },
        'postcss-mixins': true,
        'postcss-simple-vars': true,
        'postcss-nested': true,
        'lost': true,
        'postcss-preset-env': { browsers: ['last 2 versions', '> 2%'] },
        'cssnano': ctx.env === 'production' ? {} : false,
    },
})