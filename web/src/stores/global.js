import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useGlobalStore = defineStore(
  'global_s',
  () => {
    
    return {  }
  },
  {
    persist: true
  }
)
