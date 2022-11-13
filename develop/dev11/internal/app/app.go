package app

import (
	"context"
	"dev11/internal/repository/memory"
	"dev11/internal/server"
	"dev11/internal/service"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	eventRepo := memory.NewEventRepo()
	calendarService := service.NewCalendarService(eventRepo)
	handler := server.NewHandler(&calendarService)
	s := server.NewServer(handler.CreateRoutes())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		s.Start()
	}()

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	s.Stop(ctx)

	fmt.Println("Server stopped")
}
