import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { arch, platform } from 'node:process';

export const extension = platform == 'win32' ? '.exe' : '';

export const file = join(tmpdir(), `server${extension}`);

export const etag = `${file}.etag`;

const suffix = arch == 'arm64' && platform != 'win32' ? '_arm64' : '';

export const url = `https://github.com/lorypelli/server/releases/latest/download/server_${platform}${suffix}${extension}`;
