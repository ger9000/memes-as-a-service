package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ger9000/memes-as-a-service/internal/shared/config"
)

func NewServer(handler http.Handler) *http.Server {
	conf := config.GetInstance()
	timeout := func() time.Duration {
		if conf.Server.Timeout > 0 {
			return conf.Server.Timeout
		}
		return 5 * time.Second
	}()

	return &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Server.Port),
		Handler:        http.TimeoutHandler(handler, timeout, ""),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
