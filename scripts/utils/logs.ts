import { log } from 'node:console';
import { exit } from 'node:process';
import { styleText } from 'node:util';

export function info(msg: string) {
    log(
        styleText(['bold', 'bgBlue'], '  INFO  '),
        styleText(['bold', 'blueBright'], msg),
    );
}

export function error(err: unknown): never {
    const msg = err instanceof Error ? err.message : String(err);
    log(
        styleText(['bold', 'bgRed'], '  ERROR  '),
        styleText(['bold', 'redBright'], msg),
    );
    exit(1);
}
