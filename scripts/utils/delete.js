import { unlink } from 'node:fs/promises';
import { error } from './logs.js';

/**
 * @param { string } path
 */

export default async function del(path) {
    await unlink(path, { recursive: true }).catch((err) => error(err));
}
