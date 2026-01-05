<template>
  <footer class="bg-white pt-20 pb-10 px-8 border-t border-gray-50">
    <div class="max-w-7xl mx-auto grid grid-cols-1 md:grid-cols-4 gap-12 mb-16">
      <div class="space-y-6">
        <div class="text-2xl font-black tracking-tighter text-gray-900">F<span class="text-red-500">OO</span>D</div>
        <p class="text-gray-500 text-sm leading-relaxed max-w-xs">Our job is to filling your tummy with delicious food and with fast and free delivery.</p>
      </div>
      <div v-for="title in ['About', 'Company', 'Support']" :key="title">
        <h4 class="font-bold mb-6 text-gray-900">{{ title }}</h4>
        <ul class="space-y-4 text-gray-500 text-sm">
          <li class="hover:text-red-500 cursor-pointer transition">bshbdhs</li>
          <li class="hover:text-red-500 cursor-pointer transition">adsa</li>
          <li class="hover:text-red-500 cursor-pointer transition">asdj</li>
        </ul>
      </div>
    </div>
    <div class="text-center text-gray-400 text-xs border-t border-gray-50 pt-10">
      © 2026 Food Delivery Service. All rights reserved.
    </div>
  </footer>
</template>