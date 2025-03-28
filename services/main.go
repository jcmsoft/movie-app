package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"com.movie.app/services/clientService"
	"com.movie.app/services/favoriteService"
	"com.movie.app/services/movieService"
	protoServices "com.movie.app/services/servicespb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	// Create a listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server
	server := grpc.NewServer()

	// Register services
	protoServices.RegisterClientProtoServiceServer(server, &clientService.ClientProtoServiceServer{})
	protoServices.RegisterMovieProtoServiceServer(server, &movieService.MovieProtoServiceServer{})
	protoServices.RegisterFavoriteMovieProtoServiceServer(server, &favoriteService.FavoriteMovieProtoServiceServer{})

	// Start server in a goroutine
	go func() {
		log.Printf("Starting gRPC Server on port %s", port)
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	gracefulShutdown(server)
}

func gracefulShutdown(server *grpc.Server) {
	// Create channel to listen for signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-ch

	log.Println("Shutting down gRPC server...")

	// Create a deadline for server shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully stop the server
	server.GracefulStop()

	log.Println("Server stopped gracefully")
}
