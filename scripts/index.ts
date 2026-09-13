#!/usr/bin/env node

import { spawnSync } from 'node:child_process';
import { argv, exit } from 'node:process';
import { file } from './utils/constants.js';
import download from './utils/download.js';
import { error, info } from './utils/logs.js';

/**
 * The package name and version, replaced at build time by tsdown's `define`.
 */
declare const NAME: string;
declare const VERSION: string;

info(`Welcome to ${NAME}@${VERSION}!`);

await download().catch(error);

const { status, error: err } = spawnSync(file, argv.slice(2), {
    stdio: 'inherit',
});
if (err) error(err);
exit(status ?? 1);
