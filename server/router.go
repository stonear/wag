package server

var routerTemplateStr = `
package server

// Code auto-generated. Do not edit.

{{.ImportStatements}}

// Server defines a HTTP server that implements the Controller interface.
type Server struct {
	// Handler should generally not be changed. It exposed to make testing easier.
	Handler http.Handler
	addr string
	config serverConfig
}

type serverConfig struct{
	compressionLevel int
}

func CompressionLevel(level int) func(*serverConfig) {
	return func(c *serverConfig) {
		c.compressionLevel = level
	}
}

// Serve starts the server. It will return if an error occurs.
func (s *Server) Serve() error {
	go func() {
		// This should never return. Listen on the pprof port
		log.Printf("PProf server crashed: %%s", http.ListenAndServe("localhost:6060", nil))
	}()

	// Give the sever 30 seconds to shut down
	server := &http.Server{
		Addr:        s.addr,
		Handler:     s.Handler,
		IdleTimeout: 3 * time.Minute,
	}
	server.SetKeepAlivesEnabled(true)

	// Give the server 30 seconds to shut down gracefully after it receives a signal
	shutdown := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Signal(syscall.SIGTERM))
		sig := <-c
		logrus.Info("shutdown-initiated", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		defer close(shutdown)
		if err := server.Shutdown(ctx); err != nil {
			logrus.WithContext(ctx).Error("error-during-shutdown", err)
		}
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	// ensure we wait for graceful shutdown
	<-shutdown

	return nil
}

type handler struct {
	Controller
}

func withMiddleware(router http.Handler, m []func(http.Handler) http.Handler, config serverConfig) http.Handler {
	handler := router

	// compress everything
	handler = handlers.CompressHandlerLevel(handler, config.compressionLevel)

	// Wrap the middleware in the opposite order specified so that when called then run
	// in the order specified
	for i := len(m) - 1; i >= 0; i-- {
		handler = m[i](handler)
	}
	handler = PanicMiddleware(handler)

	return handler
}


// New returns a Server that implements the Controller interface. It will start when "Serve" is called.
func New(c Controller, addr string, options ...func(*serverConfig)) *Server {
	return NewWithMiddleware(c, addr, []func(http.Handler) http.Handler{}, options...)
}

// NewRouter returns a mux.Router with no middleware. This is so we can attach additional routes to the
// router if necessary
func NewRouter(c Controller) *mux.Router {
	return newRouter(c)
}

func newRouter(c Controller) *mux.Router {
	router := mux.NewRouter()

	h := handler{Controller: c}

	{{range $index, $val := .Functions}}
	router.Methods("{{$val.Method}}").Path("{{$val.Path}}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.{{$val.HandlerName}}Handler(r.Context(), w, r)
	})
	{{end}}
	return router
}

// NewWithMiddleware returns a Server that implemenets the Controller interface. It runs the
// middleware after the built-in middleware (e.g. logging), but before the controller methods.
// The middleware is executed in the order specified. The server will start when "Serve" is called.
func NewWithMiddleware(c Controller, addr string, m []func(http.Handler) http.Handler, options ...func(*serverConfig)) *Server {
	router := newRouter(c)

	return AttachMiddleware(router, addr, m, options...)
}

// AttachMiddleware attaches the given middleware to the router; this is to be used in conjunction with
// NewServer. It attaches custom middleware passed as arguments as well as the built-in middleware for
// logging, tracing, and handling panics. It should be noted that the built-in middleware executes first
// followed by the passed in middleware (in the order specified).
func AttachMiddleware(router *mux.Router, addr string, m []func(http.Handler) http.Handler, options ...func(*serverConfig)) *Server {
	// Set sane defaults, to be overriden by the varargs functions.
	// This would probably be better done in NewWithMiddleware, but there are services that call
	// AttachMiddleWare directly instead.
	config := serverConfig {
		compressionLevel: gzip.DefaultCompression,
	}
	for _, option := range options {
		option(&config)
	}


	handler := withMiddleware(router, m, config)
	return &Server{Handler: handler, addr: addr, config: config}
}`
