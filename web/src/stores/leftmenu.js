import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useLeftMenuStore = defineStore(
  'left_menu',
  () => {
    const marginLeft = ref('200')
    const ExpandStatus = ref('showAside')
    return { marginLeft,ExpandStatus }
  },
  {
    persist: true
  }
)
