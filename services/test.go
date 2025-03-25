package main

import (
	"com.movie.app/data-access/client"
	"com.movie.app/data-access/movie"
	"com.movie.app/services/favoriteService"
)

func main() {
	denzel := client.NewClientWithId(3, "Denzel", "Washington", "denzel@washington.com")
	equalizer := movie.NewMovieWithId(11, "Equalizer II", 34.98)

	favoriteService.AddClientFavoriteMovie(denzel, equalizer)
}
