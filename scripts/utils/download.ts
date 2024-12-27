import { fetch } from 'node-fetch-native';
import { Buffer } from 'node:buffer';
import { error } from './logs.ts';

export default async function download(url: string) {
    const res = await fetch(url).catch((err) => error(err));
    if (!res.ok) error(`${res.status} - ${res.statusText}`);
    return Buffer.from(await res.arrayBuffer().catch((err) => error(err)));
}
