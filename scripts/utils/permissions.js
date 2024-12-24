import { chmod } from 'node:fs/promises';
import { error } from './logs.js';

/**
 * @param { string } path
 * @param { number } permissions
 */

export default async function set(path, permissions) {
    await chmod(path, permissions).catch((err) => error(err));
}
