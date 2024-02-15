package server

import (
	"fmt"
	"log/slog"

	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"go-api-boilerplate/config"
	"go-api-boilerplate/version"
)

var mode string = gin.DebugMode

type Server struct {
	WorkDir       string
	Configuration config.Config
}

func init() {
	switch mode {
	case gin.DebugMode:
	case gin.ReleaseMode:
	case gin.TestMode:
	default:
		mode = gin.DebugMode
	}

	gin.SetMode(mode)
}

var defaultAllowOrigins = []string{
	"localhost",
	"127.0.0.1",
	"0.0.0.0",
}

func NewServer() (*Server, error) {
	workDir, err := os.MkdirTemp("", "go-api-boilerplate-service")
	if err != nil {
		return nil, err
	}

	return &Server{
		WorkDir: workDir,
	}, nil
}

func Serve(ln net.Listener) error {
	s, err := NewServer()
	if err != nil {
		return err
	}
	r := s.GenerateRoutes()

	slog.Info(fmt.Sprintf("Listening on %s (version %s)", ln.Addr(), version.Version))
	srvr := &http.Server{
		Handler: r,
	}

	// listen for a ctrl+c and stop any loaded llm
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		// unload all loaded python codes
		// if loaded.runner != nil {
		// 	loaded.runner.Close()
		// }
		os.RemoveAll(s.WorkDir)
		os.Exit(0)
	}()

	return srvr.Serve(ln)
}
