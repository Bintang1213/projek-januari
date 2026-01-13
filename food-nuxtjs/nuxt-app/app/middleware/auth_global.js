// app/middleware/auth.global.js
export default defineNuxtMiddleware((to, from) => {
  // Hanya jalankan pengecekan di browser (Client Side)
  if (process.client) {
    const userRaw = localStorage.getItem('user')
    const token = localStorage.getItem('token')
    const user = userRaw ? JSON.parse(userRaw) : null

    // 1. Proteksi Halaman Admin
    if (to.path.startsWith('/admin')) {
      // Jika tidak ada token atau role bukan admin
      if (!token || user?.role !== 'admin') {
        return navigateTo('/')
      }
    }

    // 2. Proteksi Halaman Checkout (Minimal harus login)
    if (to.path === '/checkout') {
      if (!token) {
        return navigateTo('/')
      }
    }
  }
})