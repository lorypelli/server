import { createRequire } from 'node:module';

const url = import.meta.url;
const require = createRequire(url);

export default require(
    url.endsWith('.ts') ? '../../package.json' : '../package.json',
);
