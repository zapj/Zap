import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useGlobalStore = defineStore(
  'global_s',
  () => {
    const lastFilePath = ref('')
    const settings = ref({})
    return { lastFilePath, settings }
  },
  {
    persist: true
  }
)
