import { defineConfig } from 'vite'
import graphql from '@rollup/plugin-graphql'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react(), graphql()],
})