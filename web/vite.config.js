import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import UnoCSS from 'unocss/vite'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import Inspect from 'vite-plugin-inspect'

// https://vitejs.dev/config/
export default defineConfig({
  server:{
    host:'0.0.0.0',
    proxy:{

        "/api":{
          target:'https://127.0.0.1:2828',
          changeOrigin:true,
          secure: false,      
          ws: true,
      }
    }

 
  },
  plugins: [
    vue(),
    UnoCSS(),
    Inspect(),
    AutoImport({
      imports: ['vue'],
      resolvers: [
        ElementPlusResolver(),
      ],

    }),
    Components({
      resolvers: [
        ElementPlusResolver()
      ],
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
})
