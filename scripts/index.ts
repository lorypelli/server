#!/usr/bin/env node

import chalk from 'chalk';
import { execFileSync, ExecFileSyncOptions } from 'node:child_process';
import { log } from 'node:console';
import { existsSync } from 'node:fs';
import { argv, platform } from 'node:process';
import { promisify } from 'node:util';
import { dir, extension, file } from './utils/constants.ts';
import create from './utils/create.ts';
import del from './utils/delete.ts';
import download from './utils/download.ts';
import json from './utils/json.ts';
import { error } from './utils/logs.ts';
import set from './utils/permissions.ts';
import write from './utils/write.ts';

(async () => {
    const exec = promisify<string, string[], ExecFileSyncOptions, void>(
        execFileSync,
    );
    log(
        chalk.bold.bgBlue('  INFO  '),
        chalk.bold.blueBright(`Welcome to fcy@${json.version}!`),
    );
    const url = `https://github.com/lorypelli/server/releases/latest/download/server_${platform}${extension}`;
    const buffer = await download(url);
    if (!existsSync(dir)) await create(dir);
    if (existsSync(file)) await del(file);
    await write(file, buffer);
    await set(file, 0o777);
    exec(file, argv.slice(2), { stdio: 'inherit' }).catch((err) => error(err));
})();
