package movie

import (
	"com.movie.app/data-access/handler"
	"fmt"
	"log"
)

type Movie struct {
	Id    int
	Title string
	Price float64
}

// NewMovie creates a new Movie with the given parameters
func NewMovie(title string, price float64) Movie {
	return Movie{
		Title: title,
		Price: price,
	}
}

// NewMovieWithId creates a new Movie with the given parameters
func NewMovieWithId(id int, title string, price float64) Movie {
	return Movie{
		Id:    id,
		Title: title,
		Price: price,
	}
}

// String returns a string representation of the Movie
func (m Movie) String() string {
	return fmt.Sprintf("Movie{Id: %d, Title: %s, Price: %.2f}", m.Id, m.Title, m.Price)
}

func GetAllMovies() (movies []Movie) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	// Query the database
	rows, err := db.Query("SELECT * FROM movie")
	if err != nil {
		log.Fatalf("An error occured while querying the database %s", err)
	}
	// Iterate over the rows
	for rows.Next() {
		var movie Movie
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Price)
		if err != nil {
			log.Fatalf("An error occured while scanning the rows %s", err)
		}
		movies = append(movies, movie)
	}
	return movies
}

func GetMovie(id int) (Movie, error) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	var movie Movie
	// Query the database
	err := db.QueryRow("SELECT * FROM movie WHERE Id = ?", id).Scan(&movie.Id, &movie.Title, &movie.Price)
	if err != nil {
		log.Println("An error occured while querying the database %s", err)
		return Movie{}, err
	}
	return movie, nil
}

func GetMovieByTitle(title string) (Movie, error) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	var movie Movie
	// Query the database
	err := db.QueryRow("SELECT * FROM movie WHERE Title = ?", title).Scan(&movie.Id, &movie.Title, &movie.Price)
	if err != nil {
		log.Println("An error occured while querying the database %s", err)
		return Movie{}, err
	}

	return movie, nil
}

func AddMovie(movie Movie) Movie {
	exist, foundMovie := IsExist(movie)

	if exist {
		return foundMovie
	}

	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	// Insert into the database
	result, err := db.Exec("INSERT INTO movie (Title, Price) VALUES (?, ?)", movie.Title, movie.Price)
	if err != nil {
		log.Fatalf("An error occured while inserting into the database %s", err)
	}
	// Get the last inserted Id
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("An error occured while getting the last inserted Id %s", err)
	}
	movie.Id = int(id)
	return movie
}

func UpdateMovie(movie Movie) {
	exist, _ := IsExist(movie)
	if !exist {
		log.Println("The movie does not exist")
		return
	}

	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	// Update the database
	_, err := db.Exec("UPDATE movie SET Title = ?, Price = ? WHERE Id = ?", movie.Title, movie.Price, movie.Id)
	if err != nil {
		log.Println("An error occured while updating the database %s", err)
		return
	}
}

func DeleteMovie(movie Movie) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	// Delete from the database
	_, err := db.Exec("DELETE FROM movie WHERE Id = ?", movie.Id)
	if err != nil {
		log.Fatalf("An error occured while deleting from the database %s", err)
	}
}

func IsExist(movie Movie) (bool, Movie) {
	foundMovie, err := GetMovieByTitle(movie.Title)
	if err == nil {
		return true, foundMovie
	}

	return false, movie
}
