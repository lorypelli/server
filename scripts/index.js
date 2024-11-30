import { execFileSync } from 'node:child_process';
import { chmodSync, existsSync } from 'node:fs';
import { dir, extension, file } from './utils/constants.js';
import create from './utils/create.js';
import del from './utils/delete.js';
import download from './utils/download.js';
import write from './utils/write.js';

const url = `https://github.com/lorypelli/server/releases/latest/download/server_${process.platform}${extension}`;
const buffer = await download(url);
if (!existsSync(dir)) {
    await create(dir);
}
if (existsSync(file)) {
    await del(file);
}
await write(file, buffer);
chmodSync(file, 0x777);
execFileSync(file, { stdio: 'inherit' });
