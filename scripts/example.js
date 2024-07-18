import { sleep } from "k6";
import slog from 'k6/x/slog';

export const options = {
  target: 1,
  duration: '10s',
};

export default function () {
  const consoleLogger = new slog.Logger({
    output: 'console',
    format: 'text',
    level: 'DEBUG',
  });

  const fileLogger = new slog.Logger({
    output: 'file',
    filepath: './k6_logs.log',
    format: 'json',
    level: 'WARNING',
  });

  consoleLogger.debug('example log');
  fileLogger.warn("test message", { operation: 'main', message: 'Example log', test: 123 });

  sleep(1);
}