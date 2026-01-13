// utils/api.js
const BASE_URL = 'http://localhost:3001'

// Tambahkan ini agar gampang dipanggil di komponen mana pun
export const IMAGE_URL = `${BASE_URL}/uploads`

export const apiFetch = async (endpoint, options = {}) => {
  const isFormData = options.body instanceof FormData

  const headers = {
    // Jika FormData, browser akan otomatis set boundary, jadi jangan paksa JSON
    ...(!isFormData && { 'Content-Type': 'application/json' }),
    ...(process.client && localStorage.getItem('token') 
        ? { 'Authorization': `Bearer ${localStorage.getItem('token')}` } 
        : {}),
    ...options.headers
  }

  try {
    return await $fetch(endpoint, {
      baseURL: BASE_URL,
      ...options,
      headers
    })
  } catch (err) {
    const isAuthRoute = endpoint.includes('/login') || endpoint.includes('/register')

    if (process.client) {
      // 401: Token basi / tidak ada token
      if (err.response?.status === 401 && !isAuthRoute) {
        localStorage.clear()
        window.location.href = '/' // Lebih paksa daripada reload
      }

      // 403: Token ada, tapi Role-nya bukan Admin (Akses ditolak IsAdmin Go Fiber)
      if (err.response?.status === 403) {
        alert("Akses Ditolak: Anda tidak memiliki izin untuk fitur ini.")
        window.location.href = '/'
      }
    }
    
    throw err
  }
}