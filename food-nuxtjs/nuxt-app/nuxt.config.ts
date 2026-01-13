export default defineNuxtConfig({
  srcDir: 'app/',

  css: ['~/assets/css/main.css'],

  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt'
  ],

  runtimeConfig: {
    public: {
      midtransClientKey: ''
    }
  },

  app: {
    pageTransition: { name: 'page', mode: 'out-in' },
    head: {
      title: 'nuxt dalit',
      script: [
        {
          src: 'https://app.sandbox.midtrans.com/snap/snap.js',
          defer: true
        }
      ]
    }
  },

  compatibilityDate: '2025-07-15',
  devtools: { enabled: true }
})
