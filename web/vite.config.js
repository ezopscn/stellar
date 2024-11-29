import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0',
    port: 3000,
    open: true
  },
  resolve: {
    alias: [
      // 设置 @ 为 src 的别名
      { find: '@', replacement: '/src' }
    ]
  }
});