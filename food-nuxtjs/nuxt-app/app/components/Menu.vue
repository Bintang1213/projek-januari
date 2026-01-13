<script setup>
import { ref, computed, onMounted } from 'vue'
import { apiFetch } from '~/utils/api'
import { useCartStore } from '~/stores/cart'

const cartStore = useCartStore()

// State
const menus = ref([])
const loading = ref(true)
const activeIndex = ref(null)
const selectedCategory = ref('Semua')
const categories = ['Semua', 'Makanan Utama', 'Minuman', 'Cemilan', 'Dessert', 'Promo']

// Ambil Data Menu dari Backend
const fetchPublicMenus = async () => {
  loading.value = true
  try {
    const data = await apiFetch('/api/menu')
    menus.value = data || []
  } catch (err) {
    console.error("Gagal mengambil menu:", err)
  } finally {
    loading.value = false
  }
}

// Logic Filter Kategori
const filteredMenus = computed(() => {
  if (selectedCategory.value === 'Semua') {
    return menus.value
  }
  return menus.value.filter(m => m.category === selectedCategory.value)
})

// Ambil jumlah quantity dari Store berdasarkan menu_id
const getItemQuantity = (menuId) => {
  const item = cartStore.items.find(i => i.menu_id === menuId)
  return item ? item.quantity : 0
}

// Toggle Detail/Tombol Beli
const selectMenu = (id) => {
  activeIndex.value = activeIndex.value === id ? null : id
}

onMounted(async () => {
  await fetchPublicMenus()
  // Ambil data keranjang dari database saat halaman dimuat
  if (localStorage.getItem('token')) {
    await cartStore.fetchCart()
  }
})
</script>

<template>
  <section class="py-20 px-4 md:px-8 bg-white min-h-screen">
    <div class="max-w-7xl mx-auto">
      
      <div class="text-center mb-12">
        <p class="text-red-500 font-bold text-sm mb-2 uppercase tracking-widest">Menu Lezat Kami</p>
        <h2 class="text-4xl md:text-5xl font-black text-[#2D2D2D] leading-tight">
          Menu Yang Selalu Membuat Anda <br class="hidden md:block"> Jatuh Cinta
        </h2>
      </div>

      <div class="flex flex-wrap justify-center gap-3 md:gap-4 mb-16">
        <button 
          v-for="cat in categories" :key="cat"
          @click="selectedCategory = cat"
          class="px-6 py-3 rounded-2xl font-bold transition-all duration-300 shadow-sm border text-sm md:text-base"
          :class="selectedCategory === cat 
            ? 'bg-[#F54748] text-white border-[#F54748] shadow-lg shadow-red-200 scale-105' 
            : 'bg-white border-gray-100 text-gray-500 hover:border-red-200'"
        >
          <span>{{ cat }}</span>
        </button>
      </div>

      <div v-if="loading" class="flex flex-col items-center justify-center py-20">
        <div class="animate-spin rounded-full h-12 w-12 border-t-4 border-red-500 border-solid mb-4"></div>
        <p class="font-bold text-gray-400">Menyiapkan hidangan lezat...</p>
      </div>

      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-8">
        <div 
          v-for="item in filteredMenus" :key="item.ID" 
          @click="selectMenu(item.ID)"
          class="group p-6 rounded-[2.5rem] transition-all duration-500 text-center cursor-pointer relative transform"
          :class="activeIndex === item.ID 
            ? 'bg-[#F54748] text-white shadow-2xl shadow-red-200 scale-105 z-10' 
            : 'bg-white border border-gray-100 text-gray-800 hover:shadow-2xl hover:-translate-y-2'"
        >
          <div v-if="item.category === 'Promo'" 
            class="absolute top-6 left-6 bg-yellow-400 text-white text-[10px] font-black px-3 py-1 rounded-full rotate-[-10deg] shadow-md z-20">
            PROMO HOT!
          </div>

          <div class="flex justify-center mb-6 pt-4">
            <img 
              :src="item.image ? `http://localhost:3001/uploads/${item.image}` : '/placeholder-food.png'" 
              :alt="item.name" 
              class="w-32 h-32 object-cover rounded-full drop-shadow-2xl transition-all duration-500 group-hover:rotate-6"
              :class="activeIndex === item.ID ? 'scale-110 rotate-12' : ''" 
            />
          </div>

          <div class="space-y-2">
            <h3 class="font-black text-lg capitalize leading-tight h-12 flex items-center justify-center">
              {{ item.name }}
            </h3>
            <p class="text-[11px] font-bold uppercase tracking-widest" :class="activeIndex === item.ID ? 'text-white/80' : 'text-gray-400'">
              {{ item.category }}
            </p>
            <p class="text-2xl font-black mb-6">
              <span class="text-xs mr-1">Rp</span>{{ item.price.toLocaleString('id-ID') }}
            </p>
          </div>
          
          <div 
            v-if="activeIndex === item.ID" 
            class="mt-6 flex items-center justify-between bg-white/20 backdrop-blur-md rounded-2xl p-2 animate-in fade-in zoom-in duration-300"
            @click.stop
          >
            <div class="flex items-center space-x-4 px-2">
              <button 
                @click="cartStore.decrementQuantity(item.ID)" 
                class="w-8 h-8 flex items-center justify-center bg-white/20 rounded-lg hover:bg-white/40 transition font-black"
              >−</button>
              
              <span class="font-black text-lg w-4">{{ getItemQuantity(item.ID) }}</span>
              
              <button 
                @click="cartStore.addToCart(item.ID)" 
                class="w-8 h-8 flex items-center justify-center bg-white/20 rounded-lg hover:bg-white/40 transition font-black"
              >+</button>
            </div>
            
            <button 
              @click="cartStore.addToCart(item.ID)" 
              class="bg-white text-red-500 w-10 h-10 rounded-xl shadow-lg flex items-center justify-center hover:bg-gray-100 active:scale-90 transition-all"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/>
              </svg>
            </button>
          </div>
        </div>
      </div>

      <div v-if="!loading && filteredMenus.length === 0" class="text-center py-20 border-2 border-dashed border-gray-100 rounded-[3rem]">
        <svg class="w-24 h-24 mx-auto mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        <p class="text-gray-400 font-bold">Maaf, menu kategori "{{ selectedCategory }}" sedang tidak tersedia.</p>
      </div>

      <div class="mt-20 text-center">
        <NuxtLink to="/all-menu" class="px-12 py-4 border-2 border-gray-100 text-gray-800 rounded-full font-black text-sm uppercase tracking-widest hover:bg-red-500 hover:text-white hover:border-red-500 transition-all duration-300 inline-block shadow-sm">
          Jelajahi Semua Menu
        </NuxtLink>
      </div>

    </div>
  </section>
</template>

<style scoped>
.animate-in {
  animation: scaleIn 0.3s ease-out;
}

@keyframes scaleIn {
  from { opacity: 0; transform: scale(0.9); }
  to { opacity: 1; transform: scale(1); }
}

/* Line clamp untuk nama menu agar seragam */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;  
  overflow: hidden;
}
</style>