package movieService

import (
	"com.movie.app/data-access/movie"
	protoServices "com.movie.app/services/servicespb"
	"context"
	"log"
)

type MovieProtoServiceServer struct {
	protoServices.UnimplementedMovieProtoServiceServer
}

// CreateMovie adds a new movie to the database
func CreateMovie(m movie.Movie) movie.Movie {
	return movie.AddMovie(m)
}

// EditMovie updates an existing movie in the database
func EditMovie(m movie.Movie) {
	movie.UpdateMovie(m)
}

// DeleteMovie removes a movie from the database
func DeleteMovie(m movie.Movie) {
	movie.DeleteMovie(m)
}

// GetMovieByTitle retrieves a movie by its title
func GetMovieByTitle(title string) (movie.Movie, error) {
	return movie.GetMovieByTitle(title)
}

// GetMovies retrieves all movies from the database
func GetMovies() (movies []movie.Movie) {
	return movie.GetAllMovies()
}

func (mps *MovieProtoServiceServer) GetMovieById(ctx context.Context, request *protoServices.IdRequest) (*protoServices.MovieResponse, error) {
	log.Println("GetMovieId from MovieProtoServiceServer called with request: ", request)
	foundMovie, err := movie.GetMovie(int(request.Id))

	if err != nil {
		log.Println("MovieProtoServiceServer: An error occured while searching for movie ID %s", err)
		return nil, err
	}

	log.Println("GetMovieId from MovieProtoServiceServer found this Movie: ", foundMovie)
	return &protoServices.MovieResponse{
		Id:    int32(foundMovie.Id),
		Title: foundMovie.Title,
		Price: float32(foundMovie.Price),
	}, nil

}

func (mps *MovieProtoServiceServer) IsMovieExist(ctx context.Context, request *protoServices.IdRequest) (*protoServices.IsExistResponse, error) {
	log.Println("IsMovieExist from MovieProtoServiceServer called with ID: ", request.Id)
	_, err := movie.GetMovie(int(request.Id))

	if err != nil {
		log.Fatalf("MovieProtoServiceServer: An error occured while searching for movie ID: %d, %s", request.Id, err)
	}

	return &protoServices.IsExistResponse{
		Exist: true,
	}, nil
}
