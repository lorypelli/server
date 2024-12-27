import { unlink } from 'node:fs/promises';
import { error } from './logs.ts';

export default async function del(path: string) {
    await unlink(path).catch((err) => error(err));
}
