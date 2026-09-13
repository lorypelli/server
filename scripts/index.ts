#!/usr/bin/env node

import { spawnSync } from 'node:child_process';
import { argv, exit } from 'node:process';
import { file } from './utils/constants.ts';
import download from './utils/download.ts';
import { error, info } from './utils/logs.ts';

/**
 * The package version, replaced at build time by tsdown's `define`.
 */
declare const VERSION: string;

info(`Welcome to fcy@${VERSION}!`);

await download().catch(error);

const { status, error: err } = spawnSync(file, argv.slice(2), {
    stdio: 'inherit',
});
if (err) error(err);
exit(status ?? 1);
