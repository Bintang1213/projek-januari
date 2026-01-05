<template>
  <div class="min-h-screen bg-white font-sans text-gray-900">
    <ClientOnly>
      <Toaster position="top-right" />
    </ClientOnly>

    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>

<script setup>
import 'vue3-toastify/dist/index.css'
</script>
