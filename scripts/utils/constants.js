import { tmpdir } from 'node:os';

export const dir = tmpdir();

export const extension = process.platform == 'win32' ? '.exe' : '';

export const file = `${dir}/server${extension}`;
