<template>
  <section class="bg-[#FFF9F9] -mt-40 pt-24 overflow-hidden">
    <div class="max-w-7xl mx-auto flex flex-col md:flex-row items-center px-8 md:px-16 pb-12 md:pb-20">
      
      <div class="md:w-1/2 space-y-6 md:space-y-8 z-10 text-left">
        <div class="inline-flex items-center bg-red-50 text-red-500 px-5 py-2 rounded-full text-sm font-bold shadow-sm">
          More than Faster 🍎
        </div>
        <h1 class="text-5xl md:text-7xl font-black leading-tight text-gray-900">
          Be The Fastest In Delivering Your <span class="text-red-500">Food</span>
        </h1>
        <p class="text-gray-500 text-lg md:text-xl leading-relaxed max-w-md">
          Our job is to filling your tummy with delicious food and with fast and free delivery.
        </p>
        <div class="flex flex-wrap gap-6 pt-2">
          <button class="bg-red-500 text-white px-10 py-4 rounded-full font-extrabold shadow-2xl shadow-red-200 hover:scale-105 transition">
            Get Started
          </button>
          <button class="flex items-center space-x-4 font-bold text-gray-700 hover:text-red-500 transition">
            <div class="w-12 h-12 flex items-center justify-center bg-white rounded-full shadow-lg text-red-500">▶</div>
            <span>Order Process</span>
          </button>
        </div>
      </div>
      
      <div class="md:w-1/2 relative mt-12 md:mt-0 flex justify-center md:justify-end items-end">
        <div class="absolute w-[300px] h-[300px] md:w-[500px] md:h-[500px] bg-red-500 rounded-full bottom-0 translate-y-10 -z-0"></div>
        
        <img 
          src="/by.png" 
          alt="Chef Food" 
          class="relative z-10 w-full max-w-lg h-auto drop-shadow-2xl -mb-10 md:-mb-20"
        />

        <div class="absolute bottom-10 left-0 md:-left-10 bg-white p-4 rounded-2xl shadow-2xl flex items-center space-x-4 z-20 animate-bounce">
          <div class="w-14 h-14 bg-orange-100 rounded-xl flex items-center justify-center text-3xl">🍔</div>
          <div>
            <p class="font-bold text-sm">American Burger</p>
            <div class="text-xs text-yellow-400">★★★★★</div>
            <p class="font-black text-red-500">$7.75</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>