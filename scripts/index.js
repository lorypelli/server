#!/usr/bin/env node

import { execa } from 'execa';
import { existsSync } from 'node:fs';
import { chmod } from 'node:fs/promises';
import { argv, platform } from 'node:process';
import { dir, extension, file } from './utils/constants.js';
import create from './utils/create.js';
import del from './utils/delete.js';
import download from './utils/download.js';
import { error } from './utils/logs.js';
import write from './utils/write.js';

const url = `https://github.com/lorypelli/server/releases/latest/download/server_${platform}${extension}`;
const buffer = await download(url);
if (!existsSync(dir)) await create(dir);
if (existsSync(file)) await del(file);
await write(file, buffer);
chmod(file, 0o777).catch((err) => error(err));
execa(file, argv.slice(2), { stdio: 'inherit' }).catch((err) => error(err));
