import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

const apiBaseUrl = process.env.VITE_API_URL || 'http://127.0.0.1:3000'

console.log(process.env.VITE_API_URL)

export default defineConfig({
  plugins: [react()],
  base : '/',
  server: {
    proxy: {
      "/api": {
        target: apiBaseUrl,
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
});
