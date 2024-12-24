import { fetch } from 'node-fetch-native';
import { Buffer } from 'node:buffer';
import { error } from './logs.js';

/**
 * @param { string } url
 */

export default async function download(url) {
    /** @type { Response } */
    const res = await fetch(url).catch((err) => error(err));
    if (!res.ok) error(`${res.status} - ${res.statusText}`);
    return Buffer.from(await res.arrayBuffer().catch((err) => error(err)));
}
