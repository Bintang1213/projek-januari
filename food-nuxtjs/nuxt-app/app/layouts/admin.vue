<template>
  <div v-if="isAuthorized" class="min-h-screen bg-gray-100 flex">
    <aside class="w-64 bg-gray-900 text-white min-h-screen p-6 flex flex-col fixed">
      <div class="text-2xl font-black text-red-500 mb-8">FOOD ADMIN</div>
      
      <nav class="space-y-4 font-semibold flex-1">
        <NuxtLink to="/admin" class="block hover:text-red-500 cursor-pointer transition" active-class="text-red-500">
          Dashboard
        </NuxtLink>

        <NuxtLink to="/admin/menu" class="block hover:text-red-500 cursor-pointer transition" active-class="text-red-500">
          Kelola Menu
        </NuxtLink>

        <div class="text-gray-600 cursor-not-allowed">Pesanan</div>
      </nav>
      
      <button @click="logout" class="mt-auto text-sm text-gray-400 hover:text-white flex items-center gap-2">
        <span>Logout</span>
      </button>
    </aside>

    <main class="flex-1 p-10 ml-64">
      <slot />
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const isAuthorized = ref(false)

onMounted(() => {
  if (process.client) {
    const user = JSON.parse(localStorage.getItem('user'))
    const token = localStorage.getItem('token')
    
    if (token && user?.role === 'admin') {
      isAuthorized.value = true
    } else {
      navigateTo('/')
    }
  }
})

const logout = () => {
  localStorage.clear() 
  navigateTo('/') 
}
</script>