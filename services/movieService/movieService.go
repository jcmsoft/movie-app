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

func (mps *MovieProtoServiceServer) AddMovie(ctx context.Context, request *protoServices.MovieRequest) (*protoServices.MovieResponse, error) {
	log.Println("AddMovie from MovieProtoServiceServer called with request: ", request)
	_, err := movie.GetMovieByTitle(request.Title)
	if err == nil {
		log.Printf("MovieProtoServiceServer: Movie with title %s already exists\n", request.Title)
		return nil, nil
	}
	newMovie := movie.Movie{
		Title: request.Title,
		Price: float64(request.Price),
	}

	addedMovie := movie.AddMovie(newMovie)

	return &protoServices.MovieResponse{
		Id:    int32(addedMovie.Id),
		Title: addedMovie.Title,
		Price: float32(addedMovie.Price),
	}, nil
}

func (mps *MovieProtoServiceServer) UpdateMovie(ctx context.Context, request *protoServices.MovieRequest) (*protoServices.MovieResponse, error) {
	log.Println("UpdateMovie from MovieProtoServiceServer called with request: ", request)
	movieToUpdate, err := movie.GetMovie(int(request.Id))

	if err != nil {
		log.Printf("MovieProtoServiceServer: An error occured while searching for movie title %s\n", err)
		return nil, err
	}
	movieToUpdate.Title = request.Title
	movieToUpdate.Price = float64(request.Price)
	movie.UpdateMovie(movieToUpdate)

	return &protoServices.MovieResponse{
		Id:    int32(movieToUpdate.Id),
		Title: movieToUpdate.Title,
		Price: float32(movieToUpdate.Price),
	}, nil

}

func (mps *MovieProtoServiceServer) DeleteMovie(ctx context.Context, request *protoServices.IdRequest) (*protoServices.StatusResponse, error) {
	log.Println("DeleteMovie from MovieProtoServiceServer called with request: ", request)
	movieToDelete, err := movie.GetMovie(int(request.Id))

	if err != nil {
		log.Printf("MovieProtoServiceServer: An error occured while searching for movie ID %s\n", err)
		return nil, err
	}
	movie.DeleteMovie(movieToDelete)
	return &protoServices.StatusResponse{
		Success: true,
	}, nil
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

func (mps *MovieProtoServiceServer) GetMovieByTitle(ctx context.Context, request *protoServices.TitleRequest) (*protoServices.MovieResponse, error) {
	log.Println("GetMovieByTitle from MovieProtoServiceServer called with request: ", request)
	foundMovie, err := movie.GetMovieByTitle(request.Title)

	if err != nil {
		log.Println("MovieProtoServiceServer: An error occured while searching for movie title %s", err)
		return nil, err
	}

	log.Println("GetMovieByTitle from MovieProtoServiceServer found this Movie: ", foundMovie)
	return &protoServices.MovieResponse{
		Id:    int32(foundMovie.Id),
		Title: foundMovie.Title,
		Price: float32(foundMovie.Price),
	}, nil
}

func (mps *MovieProtoServiceServer) GetAllMovies(request *protoServices.EmptyRequest, stream protoServices.MovieProtoService_GetAllMoviesServer) error {
	log.Println("GetMovies from MovieProtoServiceServer called with request: ", request)
	movies := movie.GetAllMovies()

	for _, m := range movies {
		movieResponse := &protoServices.MovieResponse{
			Id:    int32(m.Id),
			Title: m.Title,
			Price: float32(m.Price),
		}

		if err := stream.Send(movieResponse); err != nil {
			return err
		}
	}
	return nil
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
