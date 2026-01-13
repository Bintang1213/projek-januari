<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white py-10 text-[#374151]">
    <div class="max-w-6xl mx-auto px-6">
      
      <!-- Empty Cart State -->
      <div v-if="cartStore.items.length === 0" class="py-20 text-center bg-white rounded-3xl border border-gray-100 shadow-lg">
        <svg class="w-32 h-32 mx-auto mb-6 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/>
        </svg>
        <p class="text-gray-400 font-black mb-6 text-xl tracking-tighter uppercase">Keranjang Kosong</p>
        <NuxtLink to="/" class="bg-red-500 text-white px-10 py-4 rounded-full font-black text-xs uppercase tracking-[0.2em] shadow-lg shadow-red-200 inline-block hover:bg-red-600 transition">
          Mulai Belanja
        </NuxtLink>
      </div>

      <!-- Checkout Content -->
      <div v-else class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-start">
        
        <!-- Left Section: Order Info & Cart Items -->
        <section class="lg:col-span-7 space-y-6">
          
          <!-- Order Information Card -->
          <div class="bg-white p-8 rounded-3xl shadow-lg border border-gray-100">
            <div class="flex items-center gap-3 mb-8">
              <svg class="w-5 h-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/>
              </svg>
              <h2 class="text-sm font-black uppercase tracking-widest text-gray-700">Informasi Pemesanan</h2>
            </div>
            
            <!-- Delivery Method -->
            <div class="grid grid-cols-2 gap-4 mb-8">
              <div @click="deliveryMethod = 'bungkus'" 
                   class="flex items-center justify-center gap-3 p-6 rounded-2xl border-2 cursor-pointer transition-all hover:scale-105"
                   :class="deliveryMethod === 'bungkus' ? 'border-red-500 bg-red-50 text-red-600 shadow-lg shadow-red-100' : 'border-gray-200 text-gray-500 bg-white'">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"/>
                </svg>
                <span class="uppercase text-xs tracking-widest font-black">Bungkus</span>
              </div>
              <div @click="deliveryMethod = 'antar'" 
                   class="flex items-center justify-center gap-3 p-6 rounded-2xl border-2 cursor-pointer transition-all hover:scale-105"
                   :class="deliveryMethod === 'antar' ? 'border-red-500 bg-red-50 text-red-600 shadow-lg shadow-red-100' : 'border-gray-200 text-gray-500 bg-white'">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16V6a1 1 0 00-1-1H4a1 1 0 00-1 1v10a1 1 0 001 1h1m8-1a1 1 0 01-1 1H9m4-1V8a1 1 0 011-1h2.586a1 1 0 01.707.293l3.414 3.414a1 1 0 01.293.707V16a1 1 0 01-1 1h-1m-6-1a1 1 0 001 1h1M5 17a2 2 0 104 0m-4 0a2 2 0 114 0m6 0a2 2 0 104 0m-4 0a2 2 0 114 0"/>
                </svg>
                <span class="uppercase text-xs tracking-widest font-black">Diantar</span>
              </div>
            </div>

            <!-- Form Fields -->
            <div class="space-y-6">
              <div v-if="deliveryMethod === 'antar'" class="space-y-3">
                <label class="text-xs font-black uppercase text-gray-500 tracking-widest flex items-center gap-2">
                  <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"/>
                  </svg>
                  Alamat Lengkap
                </label>
                <textarea v-model="form.address" placeholder="Contoh: Jl. Sudirman No. 1, Jakarta" 
                  class="w-full p-5 rounded-2xl border-2 border-gray-200 bg-gray-50 focus:bg-white focus:border-red-500 focus:outline-none text-sm min-h-[120px] font-semibold transition-all"></textarea>
              </div>
              
              <div class="space-y-3">
                <label class="text-xs font-black uppercase text-gray-500 tracking-widest flex items-center gap-2">
                  <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
                  </svg>
                  Catatan (Opsional)
                </label>
                <input v-model="form.note" type="text" placeholder="Extra sambal pedas ya..." 
                  class="w-full p-5 rounded-2xl border-2 border-gray-200 bg-gray-50 focus:bg-white focus:border-red-500 focus:outline-none text-sm font-semibold transition-all" />
              </div>
            </div>
          </div>

          <!-- Cart Items -->
          <div class="space-y-4">
            <div class="flex items-center gap-3 mb-4">
              <svg class="w-5 h-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/>
              </svg>
              <h2 class="text-sm font-black uppercase tracking-widest text-gray-700">Pesanan Anda</h2>
            </div>
            
            <div v-for="item in cartStore.items" :key="item.id"
                 class="bg-white p-6 rounded-3xl flex gap-6 items-center border-2 border-gray-100 shadow-sm hover:shadow-md transition-all">
              <img :src="item.menu.image ? `http://localhost:3001/uploads/${item.menu.image}` : '/no-image.png'"
                   class="w-24 h-24 rounded-2xl object-cover bg-gray-50 shadow-lg" />
              <div class="flex-1">
                <div class="flex justify-between items-start mb-3">
                  <h3 class="text-gray-800 font-black text-base">{{ item.menu.name }}</h3>
                  <p class="text-red-500 text-xl font-black">{{ formatPrice(item.menu.price * item.quantity) }}</p>
                </div>
                
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-400 font-semibold">{{ formatPrice(item.menu.price) }} / item</span>
                  
                  <!-- Quantity Controls -->
                  <div class="flex items-center gap-3 bg-gray-50 rounded-full px-2 py-1">
                    <button 
                      @click="cartStore.decrementQuantity(item.menu_id)" 
                      class="w-8 h-8 flex items-center justify-center bg-white rounded-full hover:bg-red-50 hover:text-red-500 transition font-black shadow-sm"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4"/>
                      </svg>
                    </button>
                    
                    <span class="font-black text-base w-8 text-center">{{ item.quantity }}</span>
                    
                    <button 
                      @click="cartStore.addToCart(item.menu_id)" 
                      class="w-8 h-8 flex items-center justify-center bg-red-500 text-white rounded-full hover:bg-red-600 transition font-black shadow-sm"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
                      </svg>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- Right Section: Summary & Payment -->
        <aside class="lg:col-span-5">
          <div class="bg-white p-8 rounded-3xl border-2 border-gray-100 shadow-2xl sticky top-10">
            <h2 class="text-xl font-black mb-8 border-b-2 border-gray-100 pb-6 uppercase tracking-tight flex items-center gap-3">
              <svg class="w-6 h-6 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"/>
              </svg>
              Ringkasan Pesanan
            </h2>

            <!-- Payment Method -->
            <div class="mb-8 space-y-4">
              <p class="text-xs font-black text-gray-500 uppercase tracking-widest mb-4">Metode Pembayaran</p>
              <div class="space-y-3">
                <div @click="paymentMethod = 'non-tunai'" 
                     class="flex items-center gap-3 p-5 rounded-2xl border-2 cursor-pointer transition-all hover:scale-105"
                     :class="paymentMethod === 'non-tunai' ? 'border-red-500 bg-red-50 text-red-600 shadow-lg shadow-red-100' : 'border-gray-200 bg-white text-gray-500'">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"/>
                  </svg>
                  <span class="text-xs uppercase tracking-widest font-black">Non-Tunai</span>
                </div>
                <div @click="paymentMethod = 'tunai'" 
                     class="flex items-center gap-3 p-5 rounded-2xl border-2 cursor-pointer transition-all hover:scale-105"
                     :class="paymentMethod === 'tunai' ? 'border-green-500 bg-green-50 text-green-600 shadow-lg shadow-green-100' : 'border-gray-200 bg-white text-gray-500'">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"/>
                  </svg>
                  <span class="text-xs uppercase tracking-widest font-black">Tunai / COD</span>
                </div>
              </div>
            </div>

            <!-- Price Breakdown -->
            <div class="space-y-4 mb-8 border-t-2 border-dashed border-gray-200 pt-8 text-sm font-bold">
              <div class="flex justify-between text-gray-500 uppercase text-xs">
                <span>Subtotal</span>
                <span class="font-black">{{ formatPrice(subtotal) }}</span>
              </div>
              <div v-if="deliveryMethod === 'antar'" class="flex justify-between text-red-500 uppercase text-xs">
                <span>Ongkir</span>
                <span class="font-black">{{ formatPrice(10000) }}</span>
              </div>
              <div class="flex justify-between text-gray-500 uppercase text-xs">
                <span>Pajak (11%)</span>
                <span class="font-black">{{ formatPrice(tax) }}</span>
              </div>
              <div class="flex justify-between items-center pt-6 border-t-2 border-gray-100">
                <span class="text-gray-500 uppercase font-black text-xs">Total Bayar</span>
                <span class="text-3xl font-black text-red-600 tracking-tighter">{{ formatPrice(finalTotal) }}</span>
              </div>
            </div>

            <!-- Confirm Button -->
            <button @click="handleOrder" :disabled="isLoading"
                    class="w-full bg-gradient-to-r from-red-500 to-red-600 text-white py-5 rounded-2xl font-black text-base hover:from-red-600 hover:to-red-700 active:scale-95 transition-all uppercase tracking-widest disabled:from-gray-300 disabled:to-gray-400 shadow-lg shadow-red-200 flex items-center justify-center gap-3">
              <svg v-if="!isLoading" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
              </svg>
              <span>{{ isLoading ? 'Memproses...' : 'Konfirmasi Pesanan' }}</span>
            </button>
          </div>
        </aside>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, reactive, onMounted } from 'vue'
import { useCartStore } from '~/stores/cart'
import { apiFetch } from '~/utils/api'

const cartStore = useCartStore()
const deliveryMethod = ref('bungkus') 
const paymentMethod = ref('non-tunai')
const isLoading = ref(false)
const form = reactive({ address: '', note: '' })

const subtotal = computed(() => cartStore.totalPrice || 0)
const shippingFee = computed(() => deliveryMethod.value === 'antar' ? 10000 : 0)
const tax = computed(() => (subtotal.value + shippingFee.value) * 0.11)
const finalTotal = computed(() => subtotal.value + shippingFee.value + tax.value)

const formatPrice = (p) => new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(p || 0)

async function handleOrder() {
  if (deliveryMethod.value === 'antar' && !form.address) {
    alert('Alamat pengiriman wajib diisi!')
    return
  }

  isLoading.value = true

  const payload = {
    total: finalTotal.value,
    payment: paymentMethod.value,
    method: deliveryMethod.value,
    address: deliveryMethod.value === 'antar' ? form.address : 'Ambil di Toko',
    note: form.note,
    items: cartStore.items.map(item => ({
      menu_id: item.menu_id,
      name: item.menu.name,
      price: item.menu.price,
      qty: item.quantity
    }))
  }

  try {
    const response = await apiFetch('/api/order', {
      method: 'POST',
      body: payload
    })

    if (paymentMethod.value === 'non-tunai' && response.token) {
      if (window.snap) {
        window.snap.pay(response.token, {
          onSuccess: () => {
            cartStore.clearCart() 
            navigateTo(`/struk?id=${response.order_id}`)
          },
          onPending: () => {
            cartStore.clearCart()
            navigateTo(`/struk?id=${response.order_id}`)
          },
          onClose: () => {
            isLoading.value = false
            alert("Selesaikan pembayaran untuk memproses pesanan.")
          }
        })
      }
    } else {
      cartStore.clearCart()
      navigateTo(`/struk?id=${response.order_id}`)
    }
  } catch (err) {
    alert("Gagal memproses pesanan.")
    isLoading.value = false
  }
}

onMounted(() => {
  if (process.client && localStorage.getItem('token')) {
    cartStore.fetchCart()
  }
})
</script>