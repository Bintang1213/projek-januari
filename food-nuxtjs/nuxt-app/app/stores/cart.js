import { defineStore } from 'pinia'
import { apiFetch } from '~/utils/api'

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: []
  }),
  getters: {
    totalItems: (state) => state.items.reduce((total, item) => total + item.quantity, 0),
    totalPrice: (state) => state.items.reduce((total, item) => {
      // Akses melalui item.menu sesuai relasi GORM di backend
      const price = item.menu ? item.menu.price : 0
      return total + (price * item.quantity)
    }, 0)
  },
  actions: {
    async fetchCart() {
      // Jangan fetch jika tidak ada token
      const token = process.client ? localStorage.getItem('token') : null
      if (!token) {
        this.items = []
        return
      }

      try {
        const data = await apiFetch('/api/cart')
        this.items = data || []
      } catch (err) {
        this.items = []
      }
    },
    async addToCart(menuId) {
      try {
        await apiFetch('/api/cart', {
          method: 'POST',
          body: { menu_id: menuId }
        })
        await this.fetchCart()
      } catch (err) {
        alert("Silakan login terlebih dahulu")
      }
    },
    async decrementQuantity(menuId) {
      try {
        await apiFetch('/api/cart/decrement', {
          method: 'POST',
          body: { menu_id: menuId }
        })
        await this.fetchCart()
      } catch (err) {
        console.error("Gagal mengurangi item")
      }
    },
    clearCart() {
      this.items = []
    }
  }
})