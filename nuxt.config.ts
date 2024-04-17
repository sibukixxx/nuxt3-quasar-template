// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    head: {
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      title: 'Nuxt3 Example App',
      meta: [
        {
          name: 'description',
          content: 'Vue3, Quasar and etc',
        },
        { name: 'keywords', content: 'nuxt,quasar,vue3,nuxt3' },
      ],
    },
  },
  devtools: { enabled: true },
  srcDir: 'src',
  typescript: {
    shim: false,     // shimsファイル生成の無効化（VSCodeでVolarを使う場合はfalseにする）
    strict: true,    // 型チェックを厳格化
    typeCheck: false, // nuxt devまたはnuxt build時に型チェックを実行
    tsConfig: {
      compilerOptions: {
        strict: true,
        types: ['@pinia/nuxt', './type.d.ts'],
      },
    },
  },
  modules: [
    'nuxt-quasar-ui'
  ],
  quasar: {
    plugins: [
      'Notify'
    ],
    extras: {
      fontIcons: ['material-icons']
    }
  },
  vite: {
    logLevel: 'info',
  },
})
