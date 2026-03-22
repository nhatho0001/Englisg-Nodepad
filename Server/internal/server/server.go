package server

import (
	"app-notepad/configs"
	"app-notepad/internal/services"
	"app-notepad/internal/store"
	"app-notepad/router"
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	engine *gin.Engine
	cfg    *configs.Configs
	query  *store.Queries
}

func NewServer(cfg *configs.Configs, db store.DBTX) *Server {
	r := gin.New()

	return &Server{engine: r, cfg: cfg, query: store.New(db)}
}

func ConectDB(ctx context.Context, cfg *configs.Configs) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, cfg.DataBaseURl())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (s *Server) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:           net.JoinHostPort(s.cfg.HOST, s.cfg.PORT),
		Handler:        s.engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	router.InitRouter(s.engine, services.NewUserService(s.query))
	go func() {
		slog.Info(fmt.Sprintf("Start server with Port : %v", server.Addr))
		if err := server.ListenAndServe(); err != nil {
			slog.Error(fmt.Sprintf("Api serve to Listen and server! %v", err))
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
