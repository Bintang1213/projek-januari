<template>
  <nav class="w-full bg-white/90 backdrop-blur-md sticky top-0 z-50 shadow-sm">
    <div class="max-w-7xl mx-auto flex justify-between items-center px-8 md:px-16 py-4">
      <div class="text-2xl font-black tracking-tighter cursor-pointer" @click="navigateTo('/')">
        F<span class="text-red-500">OO</span>D
      </div>

      <ul class="hidden md:flex space-x-10 text-gray-600 font-semibold text-sm">
        <li class="text-red-500 cursor-pointer">Home</li>
        <li class="hover:text-red-500 cursor-pointer transition">Menu</li>
        <li class="hover:text-red-500 cursor-pointer transition">Service</li>
        <li class="hover:text-red-500 cursor-pointer transition">Shop</li>
      </ul>

      <div class="flex items-center space-x-6">
        <button class="hover:opacity-70 transition">
          <img src="/search_icon.png" alt="Search" class="w-6 h-6 object-contain" />
        </button>

        <button class="relative hover:opacity-70 transition">
          <img src="/cart.png" alt="Cart" class="w-6 h-6 object-contain" />
          <span class="absolute -top-2 -right-2 bg-red-500 text-white text-[10px] rounded-full w-4 h-4 flex items-center justify-center font-bold">2</span>
        </button>

        <ClientOnly>
          <div v-if="user" class="flex items-center space-x-4">
            <div class="text-right hidden sm:block">
              <p class="font-bold text-gray-800 text-sm leading-none">{{ user.name }}</p>
              <p class="text-[10px] text-gray-400 uppercase tracking-widest mt-1">{{ user.role }}</p>
            </div>
            <button @click="logout" class="bg-gray-100 text-gray-600 px-5 py-2 rounded-full text-xs font-bold hover:bg-red-50 hover:text-red-500 transition-all">
              Logout
            </button>
          </div>
          
          <button v-else @click="showModal = true" class="bg-red-500 text-white px-8 py-2.5 rounded-full font-bold shadow-md hover:bg-red-600 active:scale-95 transition-all text-sm">
            Sign In
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

const showModal = ref(false)
const user = ref(null)

onMounted(() => {
  const savedUser = localStorage.getItem('user')
  if (savedUser) user.value = JSON.parse(savedUser)
})

const handleLoginSuccess = (userData) => {
  user.value = userData
  showModal.value = false
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  user.value = null
  toast.warn('Anda telah keluar.', { autoClose: 1500 })
  navigateTo('/')
}
</script>