import { tmpdir } from 'node:os';
import { normalize } from 'node:path';

export const dir = tmpdir();

export const extension = process.platform == 'win32' ? '.exe' : '';

export const file = normalize(`${dir}/server${extension}`);
