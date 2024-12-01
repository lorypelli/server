import chalk from 'chalk';
import { exit } from 'node:process';

/**
 * @param { string } msg
 */

export function error(msg) {
    console.log(chalk.bold.bgRed('  ERROR  '), chalk.bold.redBright(msg));
    exit(1);
}
