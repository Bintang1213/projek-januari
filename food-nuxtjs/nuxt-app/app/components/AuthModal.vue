<script setup>
import { ref } from 'vue'
import { toast } from 'vue3-toastify'
import { apiFetch } from '~/utils/api'

const props = defineProps({ isOpen: Boolean })
const emit = defineEmits(['close', 'login-success'])

const router = useRouter()

const isLogin = ref(true)
const showPassword = ref(false)
const form = ref({ name: '', email: '', password: '', role: 'user' })

const handleAuth = async () => {
  const endpoint = isLogin.value ? '/api/login' : '/api/register'
  
  try {
    const data = await apiFetch(endpoint, { 
      method: 'POST', 
      body: form.value 
    })

    if (isLogin.value) {
      localStorage.setItem('token', data.token)
      localStorage.setItem('user', JSON.stringify(data))
      
      toast.success(`Selamat datang, ${data.name}!`, {
        theme: "colored",
        autoClose: 1500
      })

      emit('login-success', data)
      
      // AGAR HALUS: Tutup modal dulu, baru navigasi
      emit('close')
      
      // Beri jeda 300ms agar animasi penutupan modal selesai
      setTimeout(() => {
        if (data.role === 'admin') {
          router.push('/admin')
        } else {
          router.push('/')
        }
      }, 300)

    } else {
      toast.info('Pendaftaran berhasil! Silakan login.')
      isLogin.value = true
    }

  } catch (err) {
    const errorData = err.response?._data
    const msg = errorData?.message || 'Email atau Password salah!'

    toast.error(msg, {
      position: "top-right",
      theme: "colored",
      autoClose: 3000
    })
  }
}
</script>

<template>
  <Teleport to="body">
    <Transition name="fade">
      <div v-if="isOpen" class="fixed inset-0 z-[100] flex items-center justify-center p-4">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="$emit('close')"></div>
        
        <div class="bg-white w-full max-sm:max-w-xs max-w-sm rounded-2xl p-8 shadow-xl relative z-10 border border-gray-100">
          <button @click="$emit('close')" class="absolute top-4 right-4 text-gray-400 hover:text-red-500 transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>

          <div class="mb-6">
            <h2 class="text-2xl font-bold text-gray-900 leading-tight">
              {{ isLogin ? 'Masuk' : 'Daftar Akun' }}
            </h2>
            <p class="text-gray-500 text-sm mt-1">
              {{ isLogin ? 'Silakan masuk ke akun Anda.' : 'Buat akun untuk mulai memesan.' }}
            </p>
          </div>

          <form @submit.prevent="handleAuth" class="space-y-4">
            <div v-if="!isLogin">
              <label class="block text-sm font-semibold text-gray-700 mb-1">Nama Lengkap</label>
              <input v-model="form.name" type="text" placeholder="Nama Anda" class="w-full px-4 py-2.5 rounded-lg border border-gray-200 focus:border-red-500 outline-none transition text-sm" required />
            </div>
            
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1">Email</label>
              <input v-model="form.email" type="email" placeholder="contoh@mail.com" class="w-full px-4 py-2.5 rounded-lg border border-gray-200 focus:border-red-500 outline-none transition text-sm" required />
            </div>

            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1">Password</label>
              <div class="relative">
                <input :type="showPassword ? 'text' : 'password'" v-model="form.password" placeholder="••••••••" class="w-full px-4 py-2.5 rounded-lg border border-gray-200 focus:border-red-500 outline-none transition text-sm" required />
                <button type="button" @click="showPassword = !showPassword" class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400">
                  <svg v-if="!showPassword" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                  <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.542-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l18 18" />
                  </svg>
                </button>
              </div>
            </div>

            <button type="submit" class="w-full bg-red-500 text-white py-3 rounded-lg font-bold hover:bg-red-600 active:scale-[0.98] transition-all text-sm mt-2">
              {{ isLogin ? 'Masuk Sekarang' : 'Daftar Sekarang' }}
            </button>
          </form>

          <div class="mt-6 text-center">
            <button @click="isLogin = !isLogin" class="text-sm text-gray-500">
              {{ isLogin ? "Belum punya akun?" : "Sudah punya akun?" }} 
              <span class="text-red-500 font-bold hover:underline ml-1">
                 {{ isLogin ? 'Daftar' : 'Login' }}
              </span>
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.3s; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>