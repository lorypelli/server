import { createRequire } from 'node:module';
import { join } from 'node:path';
import { cwd } from 'node:process';

const require = createRequire(import.meta.url);

export default require(join(cwd(), 'package.json'));
