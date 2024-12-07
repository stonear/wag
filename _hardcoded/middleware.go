package server

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

// PanicMiddleware logs any panics. For now, we're continue throwing the panic up
// the stack so this may crash the process.
func PanicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			panicErr := recover()
			if panicErr == nil {
				return
			}
			var err error

			switch panicErr := panicErr.(type) {
			case string:
				err = fmt.Errorf(panicErr)
			case error:
				err = panicErr
			default:
				err = fmt.Errorf("unknown panic %#v of type %T", panicErr, panicErr)
			}

			logrus.WithContext(r.Context()).Error("panic", err, string(debug.Stack()))
			panic(panicErr)
		}()
		h.ServeHTTP(w, r)
	})
}
