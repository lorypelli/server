import { createRequire } from 'node:module';

const require = createRequire(import.meta.url);

export default require(
    import.meta.filename.endsWith('.ts')
        ? '../../package.json'
        : '../package.json',
);
