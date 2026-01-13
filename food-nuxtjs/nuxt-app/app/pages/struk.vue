<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white py-6 flex justify-center items-start px-4">
    <div v-if="order" class="bg-white w-full max-w-md shadow-xl rounded-2xl overflow-hidden border border-gray-100">
      
      <!-- Header Stripe -->
      <div class="h-2 bg-gradient-to-r from-red-500 to-red-600"></div>
      
      <!-- Header Section -->
      <div class="p-5">
        <div class="text-center border-b border-dashed border-gray-200 pb-4 mb-4">
          <h1 class="text-2xl font-black text-gray-900 tracking-tight mb-1">
            TE<span class="text-red-500">NG</span>KO
          </h1>
          <p class="text-[9px] font-bold text-gray-400 uppercase tracking-widest mb-2">Struk Pembayaran Digital</p>
          <div class="inline-flex items-center gap-1.5 bg-green-50 text-green-600 px-2.5 py-1 rounded-full">
            <svg class="w-2.5 h-2.5" fill="currentColor" viewBox="0 0 24 24">
              <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
            <span class="text-[9px] font-black uppercase">Pembayaran Berhasil</span>
          </div>
        </div>

        <!-- Order Info Grid -->
        <div class="space-y-2 text-[11px] mb-4">
          <div class="flex justify-between items-center">
            <span class="text-gray-400 font-bold uppercase">Pemesan</span>
            <span class="font-black text-gray-800">{{ order.User?.name || order.user?.name || 'Pelanggan' }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-gray-400 font-bold uppercase">ID Order</span>
            <span class="font-black text-gray-800">#{{ order.ID || order.id }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-gray-400 font-bold uppercase">Status</span>
            <span class="px-2 py-0.5 bg-green-500 text-white text-[8px] font-black rounded uppercase">
              {{ order.Status || order.status }}
            </span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-gray-400 font-bold uppercase">Pembayaran</span>
            <span class="font-black text-gray-800 uppercase">
              {{ (order.Payment || order.payment) === 'tunai' ? 'Tunai' : 'Non-Tunai' }}
            </span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-gray-400 font-bold uppercase">Metode</span>
            <span class="font-black text-gray-800 uppercase">
              {{ (order.OrderMethod || order.order_method) === 'antar' ? 'Diantar' : 'Bungkus' }}
            </span>
          </div>
        </div>

        <!-- Address if Delivery -->
        <div v-if="(order.OrderMethod || order.order_method) === 'antar'" class="bg-gray-50 p-3 rounded-xl mb-4 text-[10px] border border-gray-100">
          <div class="flex items-center gap-1.5 mb-1.5">
            <svg class="w-3 h-3 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"/>
            </svg>
            <span class="text-gray-500 font-black uppercase">Alamat Pengiriman</span>
          </div>
          <p class="font-semibold text-gray-700">{{ order.Address || order.address }}</p>
        </div>

        <!-- Note if exists -->
        <div v-if="order.Note || order.note" class="bg-yellow-50 p-3 rounded-xl mb-4 text-[10px] border border-yellow-100">
          <div class="flex items-center gap-1.5 mb-1.5">
            <svg class="w-3 h-3 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
            </svg>
            <span class="text-yellow-700 font-black uppercase">Catatan</span>
          </div>
          <p class="font-semibold text-gray-700 italic">"{{ order.Note || order.note }}"</p>
        </div>

        <!-- Order Items -->
        <div class="border-t border-b border-dashed border-gray-200 py-3 mb-4 space-y-2.5">
          <div v-for="item in (order.Items || order.items)" :key="item.ID || item.id" class="flex justify-between items-start">
            <div class="flex-1">
              <p class="text-xs font-black text-gray-800 uppercase mb-0.5">{{ item.Name || item.name }}</p>
              <p class="text-[9px] text-gray-400 font-semibold">
                {{ item.Qty || item.qty }}x {{ formatPrice(item.Price || item.price) }}
              </p>
            </div>
            <span class="font-black text-xs text-gray-800">{{ formatPrice((item.Price || item.price) * (item.Qty || item.qty)) }}</span>
          </div>
        </div>

        <!-- Price Summary -->
        <div class="space-y-2 mb-5 text-[11px] font-bold">
          <div class="flex justify-between text-gray-500">
            <span class="uppercase">Subtotal</span>
            <span class="font-black">{{ formatPrice(subtotal) }}</span>
          </div>
          <div v-if="(order.OrderMethod || order.order_method) === 'antar'" class="flex justify-between text-red-500">
            <span class="uppercase">Ongkir</span>
            <span class="font-black">{{ formatPrice(10000) }}</span>
          </div>
          <div class="flex justify-between text-gray-500">
            <span class="uppercase">Pajak (11%)</span>
            <span class="font-black">{{ formatPrice(tax) }}</span>
          </div>
          <div class="flex justify-between items-center pt-3 border-t-2 border-gray-200 mt-2">
            <span class="text-gray-700 uppercase font-black text-[11px]">Total Bayar</span>
            <span class="text-xl font-black text-red-600">{{ formatPrice(order.Total || order.total) }}</span>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="space-y-2.5 no-print">
          <button @click="printInvoice" class="w-full bg-gradient-to-r from-red-500 to-red-600 text-white py-3 rounded-xl font-black uppercase text-[10px] tracking-widest hover:from-red-600 hover:to-red-700 transition-all shadow-lg shadow-red-200 flex items-center justify-center gap-2">
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"/>
            </svg>
            Cetak Struk
          </button>
          <NuxtLink to="/" class="block w-full text-center bg-gray-100 text-gray-600 py-3 rounded-xl font-black uppercase text-[10px] tracking-widest hover:bg-gray-200 transition-all">
            Kembali ke Beranda
          </NuxtLink>
        </div>

        <!-- Footer Note -->
        <div class="mt-4 text-center text-[9px] text-gray-400 font-semibold">
          <p>Terima kasih telah memesan di <span class="font-black text-red-500">TENGKO</span></p>
          <p class="mt-0.5">Simpan struk ini sebagai bukti pembayaran</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { apiFetch } from '~/utils/api'

const route = useRoute()
const order = ref(null)

const getOrderId = () => {
  const raw = route.query.id
  if (!raw) return null
  return raw.toString().replace('ORDER-', '')
}

const fetchInvoice = async () => {
  const orderId = getOrderId()
  if (!orderId) return

  try {
    const data = await apiFetch(`/api/order/${orderId}`)
    order.value = data
    console.log('Data Order Masuk:', data)
  } catch (err) {
    console.error('Gagal ambil data:', err)
  }
}

const subtotal = computed(() => {
  if (!order.value) return 0
  const total = order.value.Total || order.value.total
  const method = order.value.OrderMethod || order.value.order_method
  const base = Math.round(total / 1.11)
  return method === 'antar' ? base - 10000 : base
})

const tax = computed(() => {
  if (!order.value) return 0
  const total = order.value.Total || order.value.total
  const method = order.value.OrderMethod || order.value.order_method
  const shipping = method === 'antar' ? 10000 : 0
  return total - (subtotal.value + shipping)
})

const formatPrice = (p) =>
  new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0
  }).format(p || 0)

const printInvoice = () => window.print()

onMounted(fetchInvoice)
</script>

<style>
@media print {
  .no-print {
    display: none !important;
  }
  body {
    background: white !important;
  }
}
</style>