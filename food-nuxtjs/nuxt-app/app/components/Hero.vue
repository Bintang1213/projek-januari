<template>
  <section class="bg-[#FFF9F9] -mt-40 pt-20 md:pt-20 overflow-hidden">
    <div class="max-w-7xl mx-auto flex flex-col md:flex-row items-center px-8 md:px-16 pb-12 md:pb-20">
      
      <div class="md:w-1/2 space-y-6 md:space-y-8 z-10 text-left">
        <div class="inline-flex items-center bg-red-50 text-red-500 px-5 py-2 rounded-full text-sm font-bold shadow-sm">
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
          </svg>
          Lebih dari Sekedar Cepat
        </div>
        <h1 class="text-5xl md:text-7xl font-black leading-tight text-gray-900">
          Tengko <span class="text-red-500">Badog</span>
        </h1>
        <p class="text-gray-500 text-lg md:text-xl leading-relaxed max-w-md">
          Tugas kami adalah mengisi perut Anda dengan makanan lezat dan pengiriman yang cepat.
        </p>
        <div class="flex flex-wrap gap-6 pt-2">
          <button class="bg-red-500 text-white px-10 py-4 rounded-full font-extrabold shadow-2xl shadow-red-200 hover:scale-105 transition">
            Mulai Sekarang
          </button>
          <button class="flex items-center space-x-4 font-bold text-gray-700 hover:text-red-500 transition">
            <div class="w-12 h-12 flex items-center justify-center bg-white rounded-full shadow-lg text-red-500">
              <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8 5v14l11-7z"/>
              </svg>
            </div>
            <span>Proses Pemesanan</span>
          </button>
        </div>
      </div>
      
      <div class="md:w-1/2 relative mt-12 md:mt-0 flex justify-center md:justify-end items-end">
        <div class="absolute w-[300px] h-[300px] md:w-[500px] md:h-[500px] bg-red-500 rounded-full bottom-0 translate-y-10 -z-0"></div>
        
        <img 
          src="/by.png" 
          alt="Koki Makanan" 
          class="relative z-10 w-full max-w-lg h-auto drop-shadow-2xl -mb-10 md:-mb-20"
        />

        <div class="absolute bottom-10 left-0 md:-left-10 bg-white p-4 rounded-2xl shadow-2xl flex items-center space-x-4 z-20 animate-bounce">
          <div class="w-14 h-14 bg-orange-100 rounded-xl flex items-center justify-center">
            <span class="text-2xl font-black text-orange-500">Om</span>
          </div>
          <div>
            <p class="font-bold text-sm">Bintang</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>