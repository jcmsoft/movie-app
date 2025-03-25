package favorite

import (
	"com.movie.app/data-access/client"
	"com.movie.app/data-access/handler"
	"com.movie.app/data-access/movie"
	"log"
)

func AddFavoriteMovie(client client.Client, movie movie.Movie) {
	if IsFavorite(client, movie) {
		return
	}

	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)
	// Insert the favorite movie into the database
	log.Printf("inserting favorite movie %s for client %d\n", movie.Title, client.Id)
	_, err := db.Exec("INSERT INTO favorite_movie (client_id, movie_id) VALUES (?, ?)", client.Id, movie.Id)
	if err != nil {
		log.Fatalf("An error occured while inserting the favorite movie into the database %s", err)
	}
}

func IsFavorite(client client.Client, movie movie.Movie) bool {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)
	//check if the movie is already marked as a favorite
	rows, err := db.Query("SELECT * FROM favorite_movie WHERE client_id = ? AND movie_id = ?", client.Id, movie.Id)
	if err != nil {
		log.Println("An error occured while querying the database %s", err)
		return false
	}

	if rows.Next() {
		log.Printf("This movie %s is already marked as favorite for this client", movie.Title)
		return true
	}
	return false
}

func DeleteFavoriteMovie(client client.Client, movie movie.Movie) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	// Delete the favorite movie from the database
	_, err := db.Exec("DELETE FROM favorite_movie WHERE client_id = ? AND movie_id = ?", client.Id, movie.Id)
	if err != nil {
		log.Fatalf("An error occured while deleting the favorite movie from the database %s", err)
	}
}

func GetClientFavoriteMovies(client client.Client) []movie.Movie {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	rows, err := db.Query("SELECT movie_id FROM favorite_movie WHERE client_id = ?", client.Id)
	if err != nil {
		log.Fatalf("An error occured while querying the database %s", err)
	}

	var movies []movie.Movie
	for rows.Next() {
		var movieId int
		err := rows.Scan(&movieId)
		if err != nil {
			log.Fatalf("An error occured while scanning the database %s", err)
		}

		foundMovie, _ := movie.GetMovie(movieId)
		if foundMovie != (movie.Movie{}) {
			movies = append(movies, foundMovie)
		}
	}
	return movies
}
