import { defineConfig } from 'tsup';

export default defineConfig({
    entry: ['scripts/index.ts'],
    outDir: './dist',
    target: 'node14.13.1',
    format: 'esm',
    minify: true,
    clean: true,
});
