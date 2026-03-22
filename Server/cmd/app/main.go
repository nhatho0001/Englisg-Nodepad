package main

import (
	"app-notepad/configs"
	"app-notepad/internal/server"
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Load config env file error : %v", err)
	}
	db, err := server.ConectDB(ctx, cfg)
	if err != nil {
		log.Fatalf("Connect DB fail : %v", err)
	}
	defer db.Close(ctx)
	s := server.NewServer(cfg, db)
	fmt.Printf("Connect DataBase is successful!")
	if err := s.Start(ctx); err != nil {
		log.Fatal("Errorr to start server!")
	}
	fmt.Printf("Server ")
}
