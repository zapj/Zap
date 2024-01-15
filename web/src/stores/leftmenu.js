import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useLeftMenuStore = defineStore(
  'left_menu',
  () => {
    const marginLeft = ref('0')
    return { marginLeft }
  },
  {
    persist: true
  }
)
