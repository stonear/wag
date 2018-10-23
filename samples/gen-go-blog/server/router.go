package server

// Code auto-generated. Do not edit.

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"time"

	"github.com/Clever/go-process-metrics/metrics"
	"github.com/gorilla/mux"
	"github.com/kardianos/osext"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/transport"
	"gopkg.in/Clever/kayvee-go.v6/logger"
	kvMiddleware "gopkg.in/Clever/kayvee-go.v6/middleware"
	"gopkg.in/tylerb/graceful.v1"
)

type contextKey struct{}

// Server defines a HTTP server that implements the Controller interface.
type Server struct {
	// Handler should generally not be changed. It exposed to make testing easier.
	Handler http.Handler
	addr    string
	l       logger.KayveeLogger
}

// Serve starts the server. It will return if an error occurs.
func (s *Server) Serve() error {

	go func() {
		metrics.Log("blog", 1*time.Minute)
	}()

	go func() {
		// This should never return. Listen on the pprof port
		log.Printf("PProf server crashed: %s", http.ListenAndServe(":6060", nil))
	}()

	dir, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	if err := logger.SetGlobalRouting(path.Join(dir, "kvconfig.yml")); err != nil {
		s.l.Info("please provide a kvconfig.yml file to enable app log routing")
	}

	if signalfxToken := os.Getenv("SIGNALFX_ACCESS_TOKEN"); signalfxToken != "" {
		ingestUrl := os.Getenv("SIGNALFX_INGEST_URL")
		if ingestUrl == "" {
			ingestUrl = "https://ingest.signalfx.com"
		}

		// Create a Jaeger HTTP Thrift transport
		transport := transport.NewHTTPTransport(ingestUrl+"/v1/trace",
			transport.HTTPBasicAuth("auth", signalfxToken))

		cfg, err := jaegercfg.FromEnv()
		if err != nil {
			log.Fatal("Could not parse Jaeger env vars: ", err.Error())
		}
		cfg.ServiceName = "blog"
		cfg.Sampler = &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeRateLimiting,
			Param: 500,
		}

		signalfxTracer, closer, err := cfg.NewTracer(jaegercfg.Reporter(jaeger.NewRemoteReporter(transport)))
		if err != nil {
			log.Fatal("Could not initialize jaeger tracer: ", err.Error())
		}
		defer closer.Close()

		opentracing.SetGlobalTracer(signalfxTracer)
	} else {
		s.l.Error("please set SIGNALFX_ACCESS_TOKEN to enable tracing")
	}

	s.l.Counter("server-started")

	// Give the sever 30 seconds to shut down
	return graceful.RunWithErr(s.addr, 30*time.Second, s.Handler)
}

type handler struct {
	Controller
}

func withMiddleware(serviceName string, router http.Handler, m []func(http.Handler) http.Handler) http.Handler {
	handler := router

	// Wrap the middleware in the opposite order specified so that when called then run
	// in the order specified
	for i := len(m) - 1; i >= 0; i-- {
		handler = m[i](handler)
	}
	handler = TracingMiddleware(handler)
	handler = PanicMiddleware(handler)
	// Logging middleware comes last, i.e. will be run first.
	// This makes it so that other middleware has access to the logger
	// that kvMiddleware injects into the request context.
	handler = kvMiddleware.New(handler, serviceName)
	return handler
}

// New returns a Server that implements the Controller interface. It will start when "Serve" is called.
func New(c Controller, addr string) *Server {
	return NewWithMiddleware(c, addr, []func(http.Handler) http.Handler{})
}

// NewWithMiddleware returns a Server that implemenets the Controller interface. It runs the
// middleware after the built-in middleware (e.g. logging), but before the controller methods.
// The middleware is executed in the order specified. The server will start when "Serve" is called.
func NewWithMiddleware(c Controller, addr string, m []func(http.Handler) http.Handler) *Server {
	router := mux.NewRouter()
	h := handler{Controller: c}

	l := logger.New("blog")

	router.Methods("GET").Path("/students/{student_id}/sections").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.FromContext(r.Context()).AddContext("op", "getSectionsForStudent")
		h.GetSectionsForStudentHandler(r.Context(), w, r)
		ctx := WithTracingOpName(r.Context(), "getSectionsForStudent")
		r = r.WithContext(ctx)
	})

	handler := withMiddleware("blog", router, m)
	return &Server{Handler: handler, addr: addr, l: l}
}
