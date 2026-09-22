import { defineConfig } from 'vite'
import graphql from '@rollup/plugin-graphql'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react(), graphql()],
  server: {
    proxy: {
        "/api": "http://localhost:9187",
        "/files": "http://localhost:9187",
        "/ws":    { target: "ws://localhost:9187", ws: true }
    }
  }
})