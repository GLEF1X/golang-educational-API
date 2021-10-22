package core

import (
	"fmt"
	"github.com/GLEF1X/golang-educational-API/adapters/database"
	"github.com/GLEF1X/golang-educational-API/adapters/repositories"
	"github.com/GLEF1X/golang-educational-API/api/controllers"
	log "github.com/GLEF1X/golang-educational-API/core/logging"
	"github.com/GLEF1X/golang-educational-API/serializers"
	"github.com/GLEF1X/golang-educational-API/service"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type FinalizeTask interface {
	Close()
}

type Server struct {
	Router            *fasthttprouter.Router
	cfg               *APIConfiguration
	HTTPServer        *fasthttp.Server
	finalizationTasks []FinalizeTask
}

func (s *Server) AddFinalizationTasks(tasks ...FinalizeTask) {
	s.finalizationTasks = append(s.finalizationTasks, tasks...)
}

func (s *Server) OnShutdown() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for _, task := range s.finalizationTasks {
			go task.Close()
		}
		wg.Done()
	}()
	wg.Wait()
	log.GetLogger().Debug("All finalization tasks were executed")
}

// NewServer creates a new HTTP Server
func NewServer() *Server {
	cfg := NewAPIConfiguration()
	r := fasthttprouter.New()
	h := fasthttp.CompressHandler(r.Handler)
	return &Server{
		HTTPServer: newHTTPServer(h),
		Router:     r,
		cfg:        cfg,
	}
}

// newHTTPServer creates a new HTTP Server
// TODO: configuration should be configurable
func newHTTPServer(h fasthttp.RequestHandler) *fasthttp.Server {
	return &fasthttp.Server{
		Handler:            h,
		ReadTimeout:        5 * time.Second,
		WriteTimeout:       10 * time.Second,
		MaxConnsPerIP:      500,
		MaxRequestsPerConn: 500,
	}
}

// Run starts the HTTP Server and performs a graceful shutdown
func (s *Server) Run() {
	// NOTE: Package reuseport provides a TCP net.Listener with SO_REUSEPORT support.
	// SO_REUSEPORT allows linear scaling Server performance on multi-CPU servers.

	// create a fast listener ;)
	ln, err := reuseport.Listen("tcp4", s.cfg.GetApplicationListenURL())
	if err != nil {
		log.GetLogger().Errorf("error in reuseport listener: %s", err)
	}

	// create a graceful shutdown listener
	duration := 5 * time.Second
	graceful := NewGracefulListener(ln, duration)

	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.GetLogger().Errorf("hostname unavailable: %s", err)
	}

	// Error handling
	listenErr := make(chan error, 1)

	/// Run Server
	go func() {
		log.GetLogger().Info("%s - Web Server starting on port %v", hostname, graceful.Addr())
		log.GetLogger().Info("%s - Press Ctrl+C to stop", hostname)
		listenErr <- s.HTTPServer.Serve(graceful)

	}()

	// SIGINT/SIGTERM handling
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	// Handle channels/graceful shutdown
	for {
		select {
		// If Server.ListenAndServe() cannot start due to errors such
		// as "port in use" it will return an error.
		case err := <-listenErr:
			if err != nil {
				log.GetLogger().Fatalf("listener error: %s", err)
			}
			os.Exit(0)
		// handle termination signal
		case <-osSignals:
			fmt.Printf("\n")
			log.GetLogger().Debug("%s - Shutdown signal received.\n", hostname)

			// Servers in the process of shutting down should disable KeepAlives
			m := &sync.Mutex{}
			m.Lock()
			// unsafe need to be synchronized by mutex
			s.HTTPServer.DisableKeepalive = true
			m.Unlock()

			// Attempt the graceful shutdown by closing the listener
			// and completing all inflight requests.
			if err := graceful.Close(); err != nil {
				log.GetLogger().Fatalf("error with graceful close: %s", err)
			}
			s.OnShutdown()
			log.GetLogger().Debug("%s - Server gracefully stopped.\n", hostname)
		}

	}
}

func (s *Server) InitializeCustomerController() *routes.CustomerController {
	db := database.NewConnectionPool(s.cfg.GetDatabaseConnectionURI())
	defer s.AddFinalizationTasks(db)
	customerRepository := repositories.NewUserRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	jsonSerializer := serializers.JsonSerializer{}
	return routes.NewCustomerController(customerService, jsonSerializer)
}
