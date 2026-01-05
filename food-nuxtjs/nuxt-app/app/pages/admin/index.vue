<template>
  <div>
    <h1 class="text-3xl font-bold text-gray-800">Dashboard Admin</h1>
    <p class="text-gray-500 mt-2">Selamat datang di area manajemen.</p>
    
    <div class="mt-8 p-6 bg-white rounded-2xl shadow-sm border border-gray-200">
      <p class="font-bold text-gray-700">Ringkasan Hari Ini</p>
      <div class="mt-4 text-4xl font-black text-red-500">5 Pesanan Baru</div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'admin' 
})

onMounted(() => {
  const user = JSON.parse(localStorage.getItem('user'))
  if (!user || user.role !== 'admin') {
    navigateTo('/')
  }
})
</script>