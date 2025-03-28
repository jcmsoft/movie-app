package main

import (
	"com.movie.app/data-access/client"
	"com.movie.app/data-access/favorite"
	"com.movie.app/data-access/movie"
	"fmt"
)

func printItems[T any](items []T) {
	for index, item := range items {
		fmt.Printf("# %d : %v\n", index+1, item)
	}
}

func main() {
	// create client
	mike := client.NewClient("Mike", "Smith", "mike@smith.com")
	mike = client.AddClient(mike)
	jon, _ := client.GetClientByEmail("jon@wick.com")
	denzel, _ := client.GetClientById(3)

	aquaman, _ := movie.GetMovieByTitle("Aquaman")
	johnWick, _ := movie.GetMovie(1)
	blackPanther, _ := movie.GetMovie(2)
	equalizer, _ := movie.GetMovie(10)

	favorite.AddFavoriteMovie(mike, johnWick)
	favorite.AddFavoriteMovie(mike, blackPanther)
	favorite.AddFavoriteMovie(jon, aquaman)
	favorite.AddFavoriteMovie(denzel, equalizer)

	fmt.Println("================MOVIES LIST====================")
	printItems(movie.GetAllMovies())

	fmt.Println("================CLIENTS LIST====================")
	printItems(client.GetAllClients())

	fmt.Println("================CLIENTS FAVORITE MOVIES====================")
	fmt.Printf("%s's favorite movies: %v \n", mike.FirstName, favorite.GetClientFavoriteMovies(mike.Id))
	fmt.Printf("%s's favorite movies: %v \n", jon.FirstName, favorite.GetClientFavoriteMovies(jon.Id))
	fmt.Printf("%s's favorite movies: %v \n", denzel.FirstName, favorite.GetClientFavoriteMovies(denzel.Id))
}
