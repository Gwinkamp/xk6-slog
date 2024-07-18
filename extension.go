package xk6slog

import (
	"fmt"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/slog", new(SlogExt))
}

type SlogExt struct{}

// XLogger creates a new Logger instance with the given options.
// Usage in k6 scripts: `const logger = new Logger(...);`
func (b *SlogExt) XLogger(call goja.ConstructorCall, rt *goja.Runtime) *goja.Object {
	var opts LoggerOpts

	if len(call.Arguments) == 0 {
		opts = LoggerOpts{
			Output: "console",
			Format: "text",
			Level:  "INFO",
		}
	} else {
		err := rt.ExportTo(call.Argument(0), &opts)
		if err != nil {
			panic(fmt.Errorf("error reading argument: %w", err))
		}
	}

	return rt.ToValue(NewLogger(opts)).ToObject(rt)
}
