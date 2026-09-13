import { createWriteStream } from 'node:fs';
import { access, chmod, readFile, writeFile } from 'node:fs/promises';
import { pipeline } from 'node:stream/promises';
import { etag, file, url } from './constants.ts';

export default async function download() {
    const headers = new Headers();
    const [exists, cached] = await Promise.all([
        access(file).then(
            () => true,
            () => false,
        ),
        readFile(etag, 'utf8').catch(() => ''),
    ]);
    if (exists && cached) headers.set('If-None-Match', cached);
    const res = await fetch(url, { headers });
    if (res.status == 304) return;
    if (!res.ok || !res.body)
        throw new Error(`${res.status} - ${res.statusText}`);
    await pipeline(res.body, createWriteStream(file));
    await chmod(file, 0o755);
    await writeFile(etag, res.headers.get('etag') ?? '');
}
