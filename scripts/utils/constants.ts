import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { arch, platform } from 'node:process';
import { error } from './logs.js';

const goos: Partial<Record<NodeJS.Platform, string>> = {
    darwin: 'darwin',
    linux: 'linux',
    win32: 'windows',
};

const goarch: Partial<Record<NodeJS.Architecture, string>> = {
    arm64: 'arm64',
    x64: 'amd64',
};

const os = goos[platform];

const cpu = goarch[arch];

if (!os || !cpu) error(`Unsupported platform: ${platform}/${arch}`);

export const extension = platform == 'win32' ? '.exe' : '';

export const file = join(tmpdir(), `server${extension}`);

export const etag = `${file}.etag`;

export const url = `https://github.com/lorypelli/server/releases/latest/download/server_${os}_${cpu}${extension}`;
