import { rm } from 'node:fs/promises';
import { error } from './logs.js';

/**
 * @param { string } path
 */

export default async function del(path) {
    await rm(path, { recursive: true }).catch((err) => error(err));
}
