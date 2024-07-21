import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { resolve } from 'path'
import router from '@tanstack/router-plugin/vite';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [router(), react()],
  resolve: {
    alias: {
      "@": resolve(__dirname, "./src"),
    },
  },
})
