import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    globals: true,
    environment: 'node',
    include: process.env.DAY 
      ? [`src/day/${process.env.DAY.padStart(2, '0')}/**/*.test.ts`]
      : ['src/**/*.test.ts'],
  },
});
