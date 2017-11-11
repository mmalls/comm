package comm

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/xtfly/log4g"
)

// StartHTTP Starting HTTP server, then Wait for interrupt signal to gracefully shutdown the server with a timeout duration
func StartHTTP(slog log.Logger, cfg *Common, h http.Handler, cb func()) {
	haddr := fmt.Sprintf("%s:%d", cfg.HTTPAddr, cfg.HTTPPort)
	server := http.Server{Addr: haddr, Handler: h}
	go func() {
		slog.Debugf("Starting server: %s", haddr)
		// service connections
		if err := server.ListenAndServe(); err != nil {
			slog.Errorf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	slog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Info("Server Shutdown:", err)
	}
	cb()
	slog.Info("Server exist")
}
