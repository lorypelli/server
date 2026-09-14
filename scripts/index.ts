#!/usr/bin/env node

import { spawnSync } from 'node:child_process';
import process, { argv, exit } from 'node:process';
import { file } from './utils/constants.js';
import download from './utils/download.js';

await download();

const { status, error } = spawnSync(file, argv.slice(2), { stdio: 'inherit' });
if (error) throw error;
exit(status ?? 1);
