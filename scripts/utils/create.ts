import { mkdir } from 'node:fs/promises';
import { error } from './logs.ts';

export default async function create(path: string) {
    await mkdir(path, { recursive: true }).catch((err) => error(err));
}
