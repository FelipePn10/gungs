package main

import (
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/FelipePn10/gungs/config"
	"github.com/FelipePn10/gungs/internal/database"
	"github.com/gin-gonic/gin"
)

type application struct {
	config *config.Config
	logger *slog.Logger
	db     *database.DB
}

func (app *application) traceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		app.logger.Info("request completed",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.Int64("duration_ms", duration.Milliseconds()),
			slog.String("client_ip", c.ClientIP()),
			slog.Int("status", c.Writer.Status()),
		)
	}
}

func (app *application) mount() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(app.traceMiddleware())

	// Health check global
	r.GET("/health", app.healthHandler)

	return r
}

func (app *application) healthHandler(c *gin.Context) {
	resp := map[string]any{
		"status":    "ok",
		"timestamp": time.Now().UTC(),
		"service":   "core-api",
	}
	c.JSON(http.StatusOK, resp)
}

func (app *application) run(h *gin.Engine) error {
	addr := app.config.ServerPort
	if addr == "" {
		addr = "5030"
	}
	if !strings.HasPrefix(addr, ":") {
		addr = ":" + addr
	}

	srv := &http.Server{
		Addr:         addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", addr)
	return srv.ListenAndServe()
}
