<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-800">Kelola Menu</h1>
      <button @click="bukaModalTambah" class="bg-red-500 text-white px-6 py-2 rounded-xl font-bold hover:bg-red-600 transition shadow-lg shadow-red-200">
        + Tambah Menu
      </button>
    </div>

    <div class="bg-white rounded-2xl shadow-sm border border-gray-200 overflow-hidden">
      <table class="w-full text-left">
        <thead class="bg-gray-50 border-b">
          <tr>
            <th class="p-4">Produk</th>
            <th class="p-4">Kategori</th>
            <th class="p-4">Harga</th>
            <th class="p-4">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in menus" :key="item.ID" class="border-b hover:bg-gray-50 transition">
            <td class="p-4 flex items-center gap-3">
              <img 
                :src="item.image ? `http://localhost:3001/uploads/${item.image}` : 'https://placehold.co/100x100?text=No+Image'" 
                class="w-12 h-12 rounded-lg object-cover border bg-gray-100"
              />
              <span class="font-semibold text-gray-700">{{ item.name }}</span>
            </td>
            <td class="p-4">
              <span 
                class="px-3 py-1 rounded-full text-[10px] font-black uppercase tracking-wider"
                :class="getCategoryStyle(item.category)"
              >
                {{ item.category }}
              </span>
            </td>
            <td class="p-4 text-red-600 font-bold">Rp {{ item.price.toLocaleString() }}</td>
            <td class="p-4">
              <div class="flex gap-3">
                <button @click="bukaModalEdit(item)" class="text-blue-500 hover:text-blue-700 font-medium text-sm">Edit</button>
                <button @click="hapusMenu(item.ID)" class="text-red-500 hover:text-red-700 font-medium text-sm">Hapus</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="isModalOpen" class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center p-4 z-50">
      <div class="bg-white rounded-2xl p-8 w-full max-w-md shadow-2xl">
        <h2 class="text-2xl font-bold mb-4 text-gray-800">{{ isEditMode ? 'Edit Menu' : 'Tambah Menu Baru' }}</h2>
        
        <form @submit.prevent="simpanMenu" class="space-y-4">
          <div>
            <label class="block text-sm font-bold mb-1 text-gray-600">Nama Makanan</label>
            <input v-model="form.name" type="text" class="w-full border p-2 rounded-lg focus:ring-2 focus:ring-red-500 outline-none" placeholder="Contoh: Ayam Bakar" required>
          </div>
          <div>
            <label class="block text-sm font-bold mb-1 text-gray-600">Harga</label>
            <input v-model="form.price" type="number" class="w-full border p-2 rounded-lg focus:ring-2 focus:ring-red-500 outline-none" placeholder="0" required>
          </div>
          
          <div>
            <label class="block text-sm font-bold mb-1 text-gray-600">Kategori</label>
            <select v-model="form.category" class="w-full border p-2 rounded-lg focus:ring-2 focus:ring-red-500 outline-none bg-white" required>
              <option value="" disabled>Pilih Kategori</option>
              <option v-for="cat in foodCategories" :key="cat" :value="cat">{{ cat }}</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-bold mb-1 text-gray-600">Foto Makanan</label>
            <input type="file" @change="handleFileUpload" class="w-full border p-2 rounded-lg text-sm file:mr-4 file:py-1 file:px-4 file:rounded-full file:border-0 file:bg-red-50 file:text-red-700">
          </div>

          <div class="flex justify-end gap-2 mt-6 border-t pt-4">
            <button type="button" @click="isModalOpen = false" class="px-4 py-2 text-gray-400 hover:text-gray-600 font-medium">Batal</button>
            <button type="submit" class="px-6 py-2 bg-red-500 text-white rounded-lg font-bold hover:bg-red-600 transition">
              {{ isEditMode ? 'Simpan Perubahan' : 'Tambah Menu' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { apiFetch } from '~/utils/api'

definePageMeta({ layout: 'admin' })

const menus = ref([])
const isModalOpen = ref(false)
const isEditMode = ref(false)
const selectedFile = ref(null)

// Daftar Kategori Web Makanan
const foodCategories = ['Makanan Utama', 'Minuman', 'Cemilan', 'Dessert', 'Promo']

const form = ref({
  ID: null,
  name: '',
  price: '',
  category: '',
  description: 'Default Description'
})

// Fungsi warna badge kategori
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

const handleFileUpload = (event) => {
  selectedFile.value = event.target.files[0]
}

const loadMenus = async () => {
  try {
    menus.value = await apiFetch('/api/menu')
  } catch (e) {
    console.error("Gagal load data")
  }
}

const bukaModalTambah = () => {
  isEditMode.value = false
  form.value = { ID: null, name: '', price: '', category: '', description: 'Default Description' }
  selectedFile.value = null
  isModalOpen.value = true
}

const bukaModalEdit = (item) => {
  isEditMode.value = true
  form.value = { ...item }
  selectedFile.value = null
  isModalOpen.value = true
}

const simpanMenu = async () => {
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

    alert('Berhasil disimpan!')
    isModalOpen.value = false
    loadMenus()
  } catch (err) {
    alert('Gagal menyimpan!')
  }
}

const hapusMenu = async (id) => {
  if(confirm("Yakin ingin menghapus menu ini?")) {
    try {
      await apiFetch(`/api/admin/menu/${id}`, { method: 'DELETE' })
      loadMenus()
    } catch (e) {
      alert("Gagal hapus")
    }
  }
}

onMounted(loadMenus)
</script>