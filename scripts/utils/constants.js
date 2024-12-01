import { tmpdir } from 'node:os';
import { normalize } from 'node:path';
import { platform } from 'node:process';

export const dir = tmpdir();

export const extension = platform == 'win32' ? '.exe' : '';

export const file = normalize(`${dir}/server${extension}`);
