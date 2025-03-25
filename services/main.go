package main

import (
	"com.movie.app/data-access/client"
	"com.movie.app/data-access/movie"
	"com.movie.app/services/clientService"
	"com.movie.app/services/favoriteService"
	"com.movie.app/services/movieService"
	protoServices "com.movie.app/services/servicespb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//TestServicesCRUDOperations()
	StartServerAndRegisterServices()
}

func TestServicesCRUDOperations() {
	println("================== TESTING CLIENT SERVICE ==================")

	// Use the service to create the client
	createdClient := clientService.CreateClient(client.NewClient("John", "Doe", "john@doe.com"))
	fmt.Printf("New Client Added: %v\n", createdClient)

	// Edit the client
	createdClient.FirstName = "Jane"
	clientService.EditClient(createdClient)
	fmt.Printf("New Client Edited: %v\n", createdClient)

	// Delete the client
	fmt.Println("Deleting Client: ", createdClient)
	clientService.DeleteClient(createdClient)

	// Get all clients
	fmt.Println("All remaining Clients: ", clientService.GetClients())

	println("================== TESTING MOVIE SERVICE ==================")

	// Use the service to create the movie
	createdMovie := movieService.CreateMovie(movie.NewMovie("Inception", 25.00))
	fmt.Println("Created Movie: ", createdMovie)

	// Edit the movie
	createdMovie.Price = 30.00
	movieService.EditMovie(createdMovie)
	fmt.Println("Edited Movie: ", createdMovie)

	// Delete the movie
	fmt.Println("Deleting Movie: ", createdMovie)
	movieService.DeleteMovie(createdMovie)

	// Get all movies
	fmt.Println("All remaining Movies: ", movieService.GetMovies())

	println("================== TESTING FAVORITE SERVICE ==================")
	// Use the service to add a favorite movie
	chadwick, _ := clientService.GetClientByEmail("chadwick@boseman.com")
	blackPanther, _ := movieService.GetMovieByTitle("Black Panther")

	favoriteService.AddClientFavoriteMovie(chadwick, blackPanther)
	fmt.Printf("Added %s to %s's Favorite Movies\n", blackPanther.Title, chadwick.FirstName)

	fmt.Printf("Favorite Movies for %s: %v\n", chadwick.FirstName, favoriteService.GetFavoritesByClient(chadwick))

	//remove favorite movie
	favoriteService.RemoveFavorite(chadwick, blackPanther)
	fmt.Printf("Removed %s from  %s's favorites movies\n ", blackPanther.Title, chadwick.FirstName)

	// is favorite
	fmt.Printf("Is %s part of %s's favorite movies? %v\n", blackPanther.Title, chadwick.FirstName, favoriteService.IsFavorite(chadwick, blackPanther))
}

func StartServerAndRegisterServices() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// Start the gRPC server
	server := grpc.NewServer()
	protoServices.RegisterClientProtoServiceServer(server, &clientService.ClientProtoServiceServer{})
	protoServices.RegisterMovieProtoServiceServer(server, &movieService.MovieProtoServiceServer{})

	log.Println("Starting gRPC Server on port :50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
