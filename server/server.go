package server

import (
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"furina/config"
	"furina/logger"
	"furina/render"
)

func NewServer() *gin.Engine {
	gin.SetMode(config.GetConfig().Server.RunMode)
	r := gin.New()
	r.Use(ginLogger(), ginRecovery())
	r.Static("/static", "static")
	render.Routes(r)
	return r
}

func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.URL.Path
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		if !strings.HasPrefix(p, "/static") {
			logger.Info(
				"web",
				"path", p,
				"method", c.Request.Method,
				"cost", cost,
			)
		}
	}
}

func ginRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if brokenPipe {
					logger.Error(
						c.Request.URL.Path,
						slog.Any("error", err),
					)
					c.Error(err.(error))
					c.Abort()
					return
				}

				logger.Error(
					"panic",
					"path", c.Request.URL.Path,
					"error", err,
					"stack", string(debug.Stack()),
				)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
