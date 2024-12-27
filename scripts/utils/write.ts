import { writeFile } from 'node:fs/promises';
import { error } from './logs.ts';

export default async function write(path: string, buffer: Buffer) {
    await writeFile(path, buffer).catch((err) => error(err));
}
