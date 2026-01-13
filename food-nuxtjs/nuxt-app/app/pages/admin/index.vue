<template>
  <div>
    <h1 class="text-3xl font-bold text-gray-800">Dashboard Admin</h1>
    <p class="text-gray-500 mt-2">Selamat datang di area manajemen.</p>
    
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mt-8">
      <div class="p-6 bg-white rounded-2xl shadow-sm border border-gray-200">
        <p class="text-xs font-black uppercase text-gray-400 tracking-widest">Pesanan Baru</p>
        <div class="mt-2 text-4xl font-black text-red-500">5</div>
      </div>
      
      <div class="p-6 bg-white rounded-2xl shadow-sm border border-gray-200">
        <p class="text-xs font-black uppercase text-gray-400 tracking-widest">Total Menu</p>
        <div class="mt-2 text-4xl font-black text-blue-500">12</div>
      </div>

      <div class="p-6 bg-white rounded-2xl shadow-sm border border-gray-200">
        <p class="text-xs font-black uppercase text-gray-400 tracking-widest">Status Server</p>
        <div class="mt-2 text-xl font-black text-green-500 uppercase">Online</div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'admin' 
})
</script>