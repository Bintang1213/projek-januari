<template>
  <section class="py-20 px-8 bg-white text-center">
    <div class="max-w-7xl mx-auto">
      <div class="mb-12">
        <h2 class="text-3xl md:text-4xl font-black text-gray-900 mb-2">Testimonials About Us</h2>
        <div class="w-12 h-1 bg-red-500 mx-auto rounded-full mb-6"></div>
        <p class="text-gray-500 max-w-2xl mx-auto italic">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Porttitor scelerisque nisl.
        </p>
      </div>

      <div class="relative flex justify-center items-center">
        <button class="hidden md:block absolute left-0 lg:left-10 text-gray-300 hover:text-red-500 transition">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </button>

        <div class="bg-white border border-gray-100 shadow-[0_20px_50px_rgba(0,0,0,0.05)] rounded-[3rem] p-8 md:p-16 max-w-4xl w-full flex flex-col md:flex-row items-center gap-8 md:gap-12 relative">
          
          <div class="flex-shrink-0">
            <div class="w-32 h-32 md:w-40 md:h-40 bg-yellow-400 rounded-full flex items-center justify-center overflow-hidden border-4 border-white">
              <img src="/by.png" alt="Melissa Carson" class="w-full h-full object-cover object-top scale-125 pt-2" />
            </div>
          </div>

          <div class="text-left flex-1">
            <p class="text-gray-600 text-base md:text-lg leading-relaxed mb-6">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sagittis non augue elementum vehicula id. Urna vulputate nunc aliquet sapien. Sagittis non augue elementum vehicula id. Urna vulputate nunc aliquet sapien.
            </p>
            <div>
              <h4 class="text-red-500 font-bold text-xl">Bintang</h4>
              <p class="text-gray-400 text-sm">opejhfia</p>
            </div>
          </div>
        </div>

        <button class="hidden md:block absolute right-0 lg:right-10 text-gray-300 hover:text-red-500 transition">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>
      </div>
    </div>
  </section>
</template>