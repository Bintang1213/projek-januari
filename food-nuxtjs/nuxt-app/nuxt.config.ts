export default defineNuxtConfig({
  // Memberitahu Nuxt folder utama ada di 'app'
  srcDir: 'app/',
  
  // Daftarkan file CSS global kamu di sini
  css: ['~/assets/css/main.css'],

  // Daftarkan modul tailwind
  modules: ['@nuxtjs/tailwindcss'],

  compatibilityDate: '2025-07-15',
  devtools: { enabled: true }
})