import chalk from 'chalk';

/**
 * @param { string } msg
 */

export function error(msg) {
    console.log(chalk.bold.bgRed('  ERROR  '), chalk.bold.redBright(msg));
    process.exit(1);
}
