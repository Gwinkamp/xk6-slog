# xk6-slog

Extension for [k6](https://k6.io). Designed for logging during k6 tests. Uses [slog](https://pkg.go.dev/log/slog) module

## Requirements

* [Golang 1.22.4](https://go.dev/)
* [xk6](https://k6.io/blog/extending-k6-with-xk6/)

```shell
go install go.k6.io/xk6/cmd/xk6@latest
```

## Build

From local repository:

```shell
xk6 build --with xk6-slog=.
```

From remote repository:

```shell
xk6 build --with github.com/Gwinkamp/xk6-slog
```

## Usage

In load testing scenarios:

```javascript
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
```

To run this script, you need to run the k6 executable file, which was previously built with the `xk6 build` command

```shell
./k6 run scripts/example.js
```
