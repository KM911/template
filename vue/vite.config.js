import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },

  server: {
    port: 3000,
    // 保持一致就好了 
    proxy: {
      "/hello": {
        // target: "https://www.bilibili.com/",
        // target: "http://10.21.91.5:4000/", 
        // target : "http://127.0.0.1:4000/",
        target : "http://localhost:4000/",
        changeOrigin:true, 
        // rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
})
