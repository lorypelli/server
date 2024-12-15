import chalk from 'chalk';
import { log } from 'node:console';
import { exit } from 'node:process';

/**
 * @param { string } msg
 */

export function error(msg) {
    log(chalk.bold.bgRed('  ERROR  '), chalk.bold.redBright(msg));
    exit(1);
}
