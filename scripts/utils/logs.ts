import chalk from 'chalk';
import { log } from 'node:console';
import { exit } from 'node:process';

export function error(msg: string): never {
    log(chalk.bold.bgRed('  ERROR  '), chalk.bold.redBright(msg));
    exit(1);
}
