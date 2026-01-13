<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-3xl font-bold text-gray-800">Kelola Menu</h1>
        <p class="text-gray-500 text-sm">Tambah, edit, atau hapus menu makanan anda.</p>
      </div>
      <button 
        @click="bukaModalTambah" 
        class="bg-red-500 text-white px-6 py-3 rounded-xl font-bold hover:bg-red-600 transition shadow-lg shadow-red-200 flex items-center gap-2"
      >
        <span class="text-xl">+</span> Tambah Menu
      </button>
    </div>

    <div class="bg-white rounded-2xl shadow-sm border border-gray-200 overflow-hidden">
      <table class="w-full text-left">
        <thead class="bg-gray-50 border-b">
          <tr>
            <th class="p-4 text-[10px] font-black uppercase text-gray-400">Produk</th>
            <th class="p-4 text-[10px] font-black uppercase text-gray-400">Kategori</th>
            <th class="p-4 text-[10px] font-black uppercase text-gray-400">Harga</th>
            <th class="p-4 text-[10px] font-black uppercase text-gray-400 text-center">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in menus" :key="item.ID" class="border-b hover:bg-gray-50 transition">
            <td class="p-4 flex items-center gap-4">
              <img 
                :src="item.image ? `${IMAGE_URL}/${item.image}` : 'https://placehold.co/100x100?text=No+Image'" 
                class="w-14 h-14 rounded-xl object-cover border bg-gray-100 shadow-sm"
              />
              <div>
                <span class="font-bold text-gray-800 block uppercase text-sm">{{ item.name }}</span>
              </div>
            </td>
            <td class="p-4">
              <span 
                class="px-3 py-1 rounded-full text-[10px] font-black uppercase tracking-wider"
                :class="getCategoryStyle(item.category)"
              >
                {{ item.category }}
              </span>
            </td>
            <td class="p-4 text-red-600 font-black text-sm">
              {{ formatPrice(item.price) }}
            </td>
            <td class="p-4">
              <div class="flex gap-3 justify-center">
                <button @click="bukaModalEdit(item)" class="bg-blue-50 text-blue-600 px-4 py-1.5 rounded-lg font-bold text-xs hover:bg-blue-600 hover:text-white transition">
                  Edit
                </button>
                <button @click="hapusMenu(item.ID)" class="bg-red-50 text-red-600 px-4 py-1.5 rounded-lg font-bold text-xs hover:bg-red-600 hover:text-white transition">
                  Hapus
                </button>
              </div>
            </td>
          </tr>
          <tr v-if="menus.length === 0">
            <td colspan="4" class="p-10 text-center text-gray-400 font-medium">
              Belum ada data menu. Silahkan tambah menu baru.
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="isModalOpen" class="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center p-4 z-[100]">
      <div class="bg-white rounded-[2rem] p-8 w-full max-w-md shadow-2xl animate-in fade-in zoom-in duration-300">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-black text-gray-800 uppercase tracking-tighter">
            {{ isEditMode ? 'Edit Menu' : 'Menu Baru' }}
          </h2>
          <button @click="isModalOpen = false" class="text-gray-300 hover:text-gray-600">✕</button>
        </div>
        
        <form @submit.prevent="simpanMenu" class="space-y-5">
          <div>
            <label class="block text-[10px] font-black uppercase text-gray-400 mb-1 ml-1">Nama Makanan</label>
            <input v-model="form.name" type="text" class="w-full border-2 border-gray-50 bg-gray-50 p-3 rounded-xl focus:bg-white focus:border-red-500 outline-none transition font-bold text-sm" placeholder="Contoh: Burger King" required>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-[10px] font-black uppercase text-gray-400 mb-1 ml-1">Harga (Rp)</label>
              <input v-model="form.price" type="number" class="w-full border-2 border-gray-50 bg-gray-50 p-3 rounded-xl focus:bg-white focus:border-red-500 outline-none transition font-bold text-sm" placeholder="0" required>
            </div>
            <div>
              <label class="block text-[10px] font-black uppercase text-gray-400 mb-1 ml-1">Kategori</label>
              <select v-model="form.category" class="w-full border-2 border-gray-50 bg-gray-50 p-3 rounded-xl focus:bg-white focus:border-red-500 outline-none transition font-bold text-sm bg-white" required>
                <option value="" disabled>Pilih</option>
                <option v-for="cat in foodCategories" :key="cat" :value="cat">{{ cat }}</option>
              </select>
            </div>
          </div>

          <div>
            <label class="block text-[10px] font-black uppercase text-gray-400 mb-1 ml-1">Deskripsi Ringkas</label>
            <textarea v-model="form.description" rows="2" class="w-full border-2 border-gray-50 bg-gray-50 p-3 rounded-xl focus:bg-white focus:border-red-500 outline-none transition font-bold text-sm" placeholder="Ceritakan keunggulan menu ini..."></textarea>
          </div>

          <div>
            <label class="block text-[10px] font-black uppercase text-gray-400 mb-1 ml-1">Foto Produk</label>
            <div class="relative group">
              <input type="file" @change="handleFileUpload" class="w-full border-2 border-dashed border-gray-200 p-4 rounded-xl text-xs file:hidden cursor-pointer hover:border-red-400 transition">
              <p class="absolute inset-0 flex items-center justify-center text-[10px] text-gray-400 pointer-events-none uppercase font-black">
                {{ selectedFile ? selectedFile.name : 'Klik untuk upload gambar' }}
              </p>
            </div>
          </div>

          <div class="flex gap-3 mt-8">
            <button type="button" @click="isModalOpen = false" class="flex-1 px-4 py-4 text-gray-400 hover:text-gray-600 font-black text-xs uppercase tracking-widest">Batal</button>
            <button type="submit" :disabled="isLoading" class="flex-[2] px-6 py-4 bg-red-500 text-white rounded-2xl font-black text-xs uppercase tracking-widest hover:bg-red-600 transition shadow-lg shadow-red-100 disabled:bg-gray-300">
              {{ isLoading ? 'Memproses...' : (isEditMode ? 'Simpan Perubahan' : 'Terbitkan Menu') }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { apiFetch, IMAGE_URL } from '~/utils/api'

definePageMeta({ layout: 'admin' })

const menus = ref([])
const isModalOpen = ref(false)
const isEditMode = ref(false)
const isLoading = ref(false)
const selectedFile = ref(null)

const foodCategories = ['Makanan Utama', 'Minuman', 'Cemilan', 'Dessert', 'Promo']

const form = ref({
  ID: null,
  name: '',
  price: '',
  category: '',
  description: ''
})

const formatPrice = (p) => new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(p || 0)

const getCategoryStyle = (cat) => {
  const styles = {
    'Makanan Utama': 'bg-green-100 text-green-700',
    'Minuman': 'bg-blue-100 text-blue-700',
    'Cemilan': 'bg-yellow-100 text-yellow-700',
    'Dessert': 'bg-pink-100 text-pink-700',
    'Promo': 'bg-red-100 text-red-700'
  }
  return styles[cat] || 'bg-gray-100 text-gray-700'
}

const handleFileUpload = (e) => {
  selectedFile.value = e.target.files[0]
}

const loadMenus = async () => {
  try {
    const data = await apiFetch('/api/menu')
    menus.value = data
  } catch (e) {
    console.error("Gagal memuat menu")
  }
}

const bukaModalTambah = () => {
  isEditMode.value = false
  form.value = { ID: null, name: '', price: '', category: '', description: '' }
  selectedFile.value = null
  isModalOpen.value = true
}

const bukaModalEdit = (item) => {
  isEditMode.value = true
  // Mapping data dari GORM ke form
  form.value = { 
    ID: item.ID, 
    name: item.name, 
    price: item.price, 
    category: item.category, 
    description: item.description 
  }
  selectedFile.value = null
  isModalOpen.value = true
}

const simpanMenu = async () => {
  isLoading.value = true
  const formData = new FormData()
  formData.append('name', form.value.name)
  formData.append('price', form.value.price)
  formData.append('category', form.value.category)
  formData.append('description', form.value.description)
  
  if (selectedFile.value) {
    formData.append('image', selectedFile.value)
  }

  try {
    const method = isEditMode.value ? 'PUT' : 'POST'
    const url = isEditMode.value ? `/api/admin/menu/${form.value.ID}` : '/api/admin/menu'
    
    await apiFetch(url, {
      method: method,
      body: formData
    })

    isModalOpen.value = false
    await loadMenus()
  } catch (err) {
    alert('Gagal memproses menu. Cek koneksi backend!')
  } finally {
    isLoading.value = false
  }
}

const hapusMenu = async (id) => {
  if (confirm("Hapus menu ini secara permanen?")) {
    try {
      await apiFetch(`/api/admin/menu/${id}`, { method: 'DELETE' })
      loadMenus()
    } catch (e) {
      alert("Gagal menghapus menu")
    }
  }
}

onMounted(loadMenus)
</script>