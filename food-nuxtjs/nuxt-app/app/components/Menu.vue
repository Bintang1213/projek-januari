<script setup>
import { apiFetch } from '~/utils/api'

const activeIndex = ref(null)
const menus = ref([])
const loading = ref(true)

// Daftar Kategori (Sudah termasuk Promo)
const categories = ['All', 'Makanan Utama', 'Minuman', 'Cemilan', 'Dessert', 'Promo']
const selectedCategory = ref('All')

const fetchPublicMenus = async () => {
  loading.value = true
  try {
    // Memanggil endpoint public GoFiber
    const data = await apiFetch('/api/menu')
    menus.value = data
  } catch (err) {
    console.error("Gagal mengambil menu:", err)
  } finally {
    loading.value = false
  }
}

// Logic Filter: Memfilter kategori DAN membatasi hanya 5 menu
const filteredMenus = computed(() => {
  let list = []
  if (selectedCategory.value === 'All') {
    list = menus.value
  } else {
    list = menus.value.filter(m => m.category === selectedCategory.value)
  }
  
  // HANYA MENAMPILKAN 5 MENU
  return list.slice(0, 5)
})

const selectMenu = (id) => {
  activeIndex.value = activeIndex.value === id ? null : id
}

// Ikon untuk tombol kategori
const getIcon = (cat) => {
  if (cat === 'Minuman') return '🥤'
  if (cat === 'Dessert') return '🍰'
  if (cat === 'Cemilan') return '🍟'
  if (cat === 'Promo') return '🔥'
  return '🍔'
}

onMounted(() => {
  fetchPublicMenus()
})
</script>

<template>
  <section class="py-20 px-8 bg-white">
    <div class="max-w-7xl mx-auto">
      <div class="text-center mb-12">
        <p class="text-red-500 font-medium text-sm mb-2">Our Menu</p>
        <h2 class="text-3xl md:text-5xl font-black text-[#2D2D2D]">
          Menu That Always Make You <br class="hidden md:block"> To Fall In Love
        </h2>
      </div>

      <div class="flex flex-wrap justify-center gap-4 mb-16">
        <button 
          v-for="cat in categories" :key="cat"
          @click="selectedCategory = cat"
          class="flex items-center space-x-3 px-8 py-3 rounded-xl font-bold transition shadow-sm border"
          :class="selectedCategory === cat 
            ? 'bg-[#F54748] text-white border-[#F54748]' 
            : 'bg-white border-gray-100 text-gray-500 hover:border-red-200'"
        >
          <span class="text-xl">{{ getIcon(cat) }}</span>
          <span>{{ cat }}</span>
        </button>
      </div>

      <div v-if="loading" class="text-center py-10 font-bold text-gray-400">
        Memuat menu lezat...
      </div>

      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-8">
        <div v-for="item in filteredMenus" :key="item.ID" 
          @click="selectMenu(item.ID)"
          class="p-6 rounded-[2.5rem] transition duration-500 text-center cursor-pointer relative transform"
          :class="activeIndex === item.ID 
            ? 'bg-[#F54748] text-white shadow-2xl shadow-red-200 scale-105 z-10' 
            : 'bg-white border border-gray-100 text-gray-800 hover:shadow-xl hover:-translate-y-2'"
        >
          <div v-if="item.category === 'Promo'" 
            class="absolute top-2 left-2 bg-yellow-400 text-white text-[10px] font-black px-2 py-1 rounded-lg rotate-[-10deg] shadow-sm z-20">
            PROMO!
          </div>

          <div class="flex justify-center mb-6">
            <img 
              :src="item.image ? `http://localhost:3001/uploads/${item.image}` : '/kol.png'" 
              :alt="item.name" 
              class="w-24 h-24 object-cover rounded-full drop-shadow-lg transition duration-500"
              :class="activeIndex === item.ID ? 'scale-110' : ''" 
            />
          </div>

          <h3 class="font-bold text-lg mb-1 capitalize leading-tight h-12 flex items-center justify-center line-clamp-2">
            {{ item.name }}
          </h3>
          <p class="text-[10px] mb-4" :class="activeIndex === item.ID ? 'text-white/80' : 'text-gray-400'">
            {{ item.category }}
          </p>
          <p class="text-xl font-black mb-6">Rp {{ item.price.toLocaleString('id-ID') }}</p>
          
          <div v-if="activeIndex === item.ID" class="flex items-center justify-between bg-white/10 rounded-xl p-2 animate-in fade-in zoom-in duration-300">
            <div class="flex items-center space-x-3">
              <button @click.stop class="text-lg font-bold hover:scale-125 transition">−</button>
              <span class="font-bold">1</span>
              <button @click.stop class="text-lg font-bold hover:scale-125 transition">+</button>
            </div>
            <button @click.stop class="bg-white text-red-500 p-2 rounded-lg shadow-md hover:bg-gray-100">
               🛒
            </button>
          </div>
        </div>
      </div>

      <div v-if="!loading && filteredMenus.length === 0" class="text-center py-10 text-gray-400">
        Belum ada menu di kategori ini.
      </div>

      <div class="mt-20 text-center">
        <button class="px-10 py-3 border border-gray-200 text-red-500 rounded-full font-bold hover:bg-red-500 hover:text-white transition duration-300">
          Show More
        </button>
      </div>
    </div>
  </section>
</template>