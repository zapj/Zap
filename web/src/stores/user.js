import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('userstore', {
    state:()=>{
        return {
            userInfo:{
                username:""
            }
        }
    },
    persist:{
        storage: sessionStorage,
    }
})
