<template>
  <section class="py-20 px-8 bg-[#FFF9F9] rounded-[3rem] mx-4 md:mx-16 mb-20 overflow-hidden">
    <div class="max-w-7xl mx-auto flex flex-col md:flex-row items-center justify-between gap-12">
      <div class="md:w-1/2 space-y-6 text-center md:text-left">
        <p class="text-red-500 font-bold uppercase text-xs">Download our app</p>
        <h2 class="text-4xl md:text-5xl font-black leading-tight text-gray-900">Get the fastest food delivery app</h2>
        <p class="text-gray-500 text-lg">Available on iOS and Android for free. Order your favorite food now and get exclusive discounts!</p>
        <div class="flex flex-wrap justify-center md:justify-start gap-4 pt-4">
          <button class="bg-black text-white px-6 py-3 rounded-2xl flex items-center space-x-3 hover:bg-gray-800 transition shadow-xl">
            <span class="text-2xl">🍎</span>
            <div class="text-left leading-none"><p class="text-[10px] mb-1 opacity-70">Download on</p><p class="text-base font-bold">App Store</p></div>
          </button>
          <button class="bg-black text-white px-6 py-3 rounded-2xl flex items-center space-x-3 hover:bg-gray-800 transition shadow-xl">
            <span class="text-2xl">🤖</span>
            <div class="text-left leading-none"><p class="text-[10px] mb-1 opacity-70">Get it on</p><p class="text-base font-bold">Google Play</p></div>
          </button>
        </div>
      </div>
      <div class="md:w-1/2 relative flex justify-center">
        <div class="w-64 h-[500px] bg-white rounded-[3rem] border-[12px] border-gray-900 shadow-2xl overflow-hidden relative z-10">
           <div class="p-6 space-y-4">
              <div class="flex justify-between items-center mb-4 text-xs font-bold"><span>9:41</span><span>🔋</span></div>
              <div class="h-4 w-3/4 bg-gray-100 rounded"></div>
              <div class="h-32 w-full bg-red-50 rounded-2xl flex items-center justify-center text-3xl">🍔</div>
              <div class="space-y-2">
                 <div class="h-3 w-full bg-gray-50 rounded"></div>
                 <div class="h-3 w-5/6 bg-gray-50 rounded"></div>
              </div>
              <div class="h-10 w-full bg-red-500 rounded-xl"></div>
           </div>
        </div>
        <div class="absolute w-[400px] h-[400px] bg-red-500 rounded-full bottom-0 right-0 opacity-10 -z-0"></div>
      </div>
    </div>
  </section>
</template>