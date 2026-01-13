<template>
  <nav class="w-full bg-white/90 backdrop-blur-md sticky top-0 z-50 shadow-sm">
    <div class="max-w-7xl mx-auto flex justify-between items-center px-8 md:px-16 py-4">
      
      <div class="text-2xl font-black tracking-tighter cursor-pointer" @click="navigateTo('/')">
        T<span class="text-red-500">ENG</span>KO
      </div>

      <ul class="hidden md:flex space-x-10 text-gray-600 font-semibold text-sm">
        <li class="text-red-500 cursor-pointer" @click="navigateTo('/')">Beranda</li>
        <li class="hover:text-red-500 cursor-pointer transition">Menu</li>
        <li class="hover:text-red-500 cursor-pointer transition">Layanan</li>
        <li class="hover:text-red-500 cursor-pointer transition">Toko</li>
      </ul>

      <div class="flex items-center space-x-6">
        <button class="hover:opacity-70 transition">
          <img src="/search_icon.png" alt="Cari" class="w-6 h-6 object-contain" />
        </button>

        <button class="relative hover:opacity-70 transition" @click="navigateTo('/checkout')">
          <img src="/cart.png" alt="Keranjang" class="w-6 h-6 object-contain" />
          <span 
            v-if="cartStore.totalItems > 0"
            class="absolute -top-2 -right-2 bg-red-500 text-white text-[10px] rounded-full min-w-[18px] h-[18px] flex items-center justify-center font-bold px-1 border-2 border-white"
          >
            {{ cartStore.totalItems }}
          </span>
        </button>

        <ClientOnly>
          <div v-if="user" class="flex items-center space-x-4">
            <div class="text-right hidden sm:block">
              <p class="font-bold text-gray-800 text-sm leading-none">{{ user.name }}</p>
              <p class="text-[10px] text-gray-400 uppercase tracking-widest mt-1">{{ user.role }}</p>
            </div>
            <button @click="logout" class="bg-gray-100 text-gray-600 px-5 py-2 rounded-full text-xs font-bold hover:bg-red-50 hover:text-red-500 transition-all">
              Keluar
            </button>
          </div>
          
          <button v-else @click="showModal = true" class="bg-red-500 text-white px-8 py-2.5 rounded-full font-bold text-sm shadow-sm hover:bg-red-600 active:scale-95 transition-all">
            Masuk
          </button>
        </ClientOnly>
      </div>
    </div>

    <AuthModal 
      :is-open="showModal" 
      @close="showModal = false" 
      @login-success="handleLoginSuccess" 
    />
  </nav>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { toast } from 'vue3-toastify'
import { useCartStore } from '~/stores/cart'

const cartStore = useCartStore()
const showModal = ref(false)
const user = ref(null)

// Ambil data user dari localstorage saat pertama kali load halaman
onMounted(async () => {
  const savedUser = localStorage.getItem('user')
  const token = localStorage.getItem('token')
  
  if (savedUser) {
    user.value = JSON.parse(savedUser)
  }
  
  if (token) {
    try {
      await cartStore.fetchCart() 
    } catch (err) {
      console.error("Gagal mengambil data keranjang")
    }
  }
})

// Fungsi ini dipanggil saat AuthModal emit 'login-success'
const handleLoginSuccess = (userData) => {
  user.value = userData 
  showModal.value = false 
  cartStore.fetchCart()
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  
  cartStore.clearCart() 
  user.value = null
  
  toast.warn('Anda telah keluar.', { autoClose: 1500 })
}
</script>