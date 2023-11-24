import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

const apiBaseUrl = 'http://127.0.0.1:3000'

export default defineConfig({
  base: '/micro-reddit/',
  plugins: [react()],
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
