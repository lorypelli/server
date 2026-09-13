import { defineConfig } from 'tsdown';
import json from './package.json' with { type: 'json' };

export default defineConfig({
    entry: ['scripts/index.ts'],
    minify: true,
    define: {
        NAME: JSON.stringify(json.name),
        VERSION: JSON.stringify(json.version),
    },
});
