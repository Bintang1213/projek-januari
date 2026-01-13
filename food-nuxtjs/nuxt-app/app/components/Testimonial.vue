<template>
  <section class="py-20 px-8 bg-white">
    <div class="max-w-7xl mx-auto">
      <div class="text-center mb-16">
        <p class="text-red-500 font-bold text-sm mb-2 uppercase tracking-widest">Testimoni</p>
        <h2 class="text-4xl md:text-5xl font-black text-gray-900 mb-4">Apa Kata Pelanggan Kami</h2>
        <p class="text-gray-500 max-w-2xl mx-auto text-lg">
          Feedback asli dari orang-orang yang menyukai makanan kami
        </p>
      </div>

      <div class="relative flex justify-center items-center">
        <button class="hidden md:flex absolute left-0 lg:-left-4 w-14 h-14 bg-white rounded-full shadow-lg items-center justify-center text-gray-400 hover:text-red-500 hover:shadow-xl transition-all duration-300 z-10 hover:scale-110">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </button>

        <div class="bg-white shadow-2xl shadow-red-100 rounded-[2.5rem] p-8 md:p-12 max-w-5xl w-full relative overflow-hidden border border-gray-100">
          
          <!-- Quote Icon Background -->
          <div class="absolute top-8 right-8 text-red-100 opacity-20">
            <svg class="w-24 h-24" fill="currentColor" viewBox="0 0 24 24">
              <path d="M6 17h3l2-4V7H5v6h3zm8 0h3l2-4V7h-6v6h3z"/>
            </svg>
          </div>

          <div class="flex flex-col md:flex-row items-center gap-8 md:gap-12 relative z-10">
            
            <!-- Avatar -->
            <div class="flex-shrink-0">
              <div class="relative">
                <div class="w-32 h-32 md:w-40 md:h-40 bg-gradient-to-br from-red-400 to-red-600 rounded-full flex items-center justify-center overflow-hidden shadow-2xl shadow-red-200">
                  <img src="/by.png" alt="Bintang" class="w-full h-full object-cover object-top scale-125 pt-2" />
                </div>
                <!-- Badge Verified -->
                <div class="absolute -bottom-2 -right-2 bg-green-500 rounded-full p-2 shadow-lg">
                  <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                  </svg>
                </div>
              </div>
            </div>

            <!-- Content -->
            <div class="text-center md:text-left flex-1">
              <!-- Stars Rating -->
              <div class="flex justify-center md:justify-start gap-1 mb-4">
                <svg v-for="i in 5" :key="i" class="w-5 h-5 text-yellow-400 fill-current" viewBox="0 0 24 24">
                  <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
                </svg>
              </div>

              <p class="text-gray-600 text-base md:text-lg leading-relaxed mb-6 italic">
                "Pelayanannya cepat banget! Makanannya enak, hangat, dan selalu fresh. Delivery driver-nya juga ramah. Pokoknya the best deh untuk delivery food di area ini!"
              </p>
              
              <div>
                <h4 class="text-gray-900 font-black text-xl mb-1">Bintang</h4>
                <p class="text-gray-400 text-sm font-medium">Mahasiswa</p>
              </div>
            </div>
          </div>

          <!-- Dots Indicator -->
          <div class="flex justify-center gap-2 mt-8">
            <button class="w-2 h-2 rounded-full bg-red-500"></button>
            <button class="w-2 h-2 rounded-full bg-gray-200 hover:bg-gray-300"></button>
            <button class="w-2 h-2 rounded-full bg-gray-200 hover:bg-gray-300"></button>
          </div>
        </div>

        <button class="hidden md:flex absolute right-0 lg:-right-4 w-14 h-14 bg-white rounded-full shadow-lg items-center justify-center text-gray-400 hover:text-red-500 hover:shadow-xl transition-all duration-300 z-10 hover:scale-110">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>
      </div>

      <!-- Navigation Mobile -->
      <div class="flex md:hidden justify-center gap-4 mt-8">
        <button class="w-12 h-12 bg-white rounded-full shadow-lg flex items-center justify-center text-gray-400 hover:text-red-500 transition">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </button>
        <button class="w-12 h-12 bg-white rounded-full shadow-lg flex items-center justify-center text-gray-400 hover:text-red-500 transition">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>
      </div>
    </div>
  </section>
</template>