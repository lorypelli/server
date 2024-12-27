import { chmod } from 'node:fs/promises';
import { error } from './logs.ts';

export default async function set(path: string, permissions: number) {
    await chmod(path, permissions).catch((err) => error(err));
}
