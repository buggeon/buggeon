import { defineConfig } from 'vite'
import graphql from '@rollup/plugin-graphql'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react(), graphql()],
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: (content, url) => {
          if (url.includes('variables.scss')) {
            return content;
          }
          return `@use "/src/variables.scss" as *;\n${content}`;
        }
      }
    }
  }
})