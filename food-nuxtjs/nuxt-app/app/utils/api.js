// utils/api.js
const BASE_URL = 'http://localhost:3001'

export const apiFetch = async (endpoint, options = {}) => {
  const token = localStorage.getItem('token')
  const isFormData = options.body instanceof FormData

  const defaultOptions = {
    baseURL: BASE_URL,
    ...options, // Pindahkan ke atas agar headers di bawah tidak tertimpa
    headers: {
      // 1. Set default Content-Type jika bukan FormData
      ...(!isFormData && { 'Content-Type': 'application/json' }),
      // 2. Masukkan Token jika ada
      ...(token && { Authorization: `Bearer ${token}` }),
      // 3. Masukkan header tambahan jika user mengirimnya di parameter options
      ...options.headers 
    }
  }

  try {
    return await $fetch(endpoint, defaultOptions)
  } catch (err) {
    // Log ini sangat penting untuk debug di GoFiber
    console.error("API Error Detail:", err.response?._data || err.message)
    throw err
  }
}