package favoriteService

import (
	"com.movie.app/data-access/client"
	"com.movie.app/data-access/favorite"
	"com.movie.app/data-access/movie"
	protoServices "com.movie.app/services/servicespb"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type FavoriteMovieProtoServiceServer struct {
	protoServices.UnimplementedFavoriteMovieProtoServiceServer
}

// AddClientFavoriteMovie adds a movie to a client's favorites
func (csp *FavoriteMovieProtoServiceServer) AddClientFavoriteMovie(ctx context.Context, request *protoServices.FavoriteMovieRequest) (*protoServices.StatusResponse, error) {
	log.Println("AddFavorite from FavoriteMovieProtoServiceServer called with request: ", request)
	success := AddFavoriteMovie(request.ClientId, request.MovieId)

	return &protoServices.StatusResponse{
		Success: success,
	}, nil
}

func (csp *FavoriteMovieProtoServiceServer) RemoveClientFavoriteMovie(ctx context.Context, request *protoServices.FavoriteMovieRequest) (*protoServices.StatusResponse, error) {
	log.Println("RemoveFavorite from FavoriteMovieProtoServiceServer called with request: ", request)
	success := removeFavoriteMovie(request.ClientId, request.MovieId)

	return &protoServices.StatusResponse{
		Success: success,
	}, nil
}

func (csp *FavoriteMovieProtoServiceServer) GetClientFavoriteMovies(request *protoServices.IdRequest, stream protoServices.FavoriteMovieProtoService_GetClientFavoriteMoviesServer) error {
	log.Println("GetFavoritesByClient from FavoriteMovieProtoServiceServer called with request: ", request)
	movies := favorite.GetClientFavoriteMovies(int(request.Id))

	for _, m := range movies {
		err := stream.Send(&protoServices.MovieResponse{
			Id:    int32(m.Id),
			Title: m.Title,
			Price: float32(m.Price),
		})
		if err != nil {
			log.Println("Error sending movie: ", err)
			return err
		}
	}
	return nil
}

// addFavoriteMovie adds a movie to a client's favorites
func AddFavoriteMovie(clientId int32, movieId int32) (success bool) {
	conn, _ := openConnection()
	defer closeConnection(conn)

	protoClient := protoServices.NewClientProtoServiceClient(conn)
	protoMovie := protoServices.NewMovieProtoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	clientResp, err := protoClient.GetClientById(ctx, &protoServices.IdRequest{Id: clientId})
	if err != nil {
		log.Printf("Error Getting the client: %v\n", err)
		return false
	}

	movieResp, err := protoMovie.GetMovieById(ctx, &protoServices.IdRequest{Id: movieId})
	if err != nil {
		log.Printf("Error getting the movie: %v", err)
		return false
	}

	foundClient := client.NewClientWithId(int(clientResp.Id), clientResp.FirstName, clientResp.LastName, clientResp.Email)
	foundMovie := movie.NewMovieWithId(int(movieResp.Id), movieResp.Title, float64(movieResp.Price))

	if IsFavorite(foundClient, foundMovie) {
		log.Printf("%s is already in %s's favorites", foundMovie.Title, foundClient.FirstName)
		return false
	}

	favorite.AddFavoriteMovie(foundClient, foundMovie)
	return true
}

// RemoveFavorite removes a movie from a client's favorites
func removeFavoriteMovie(clientId int32, movieId int32) (success bool) {
	conn, _ := openConnection()
	defer closeConnection(conn)

	protoClient := protoServices.NewClientProtoServiceClient(conn)
	protoMovie := protoServices.NewMovieProtoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	clientResp, err := protoClient.GetClientById(ctx, &protoServices.IdRequest{Id: clientId})
	if err != nil {
		log.Printf("Error Getting the client: %v\n", err)
		return false
	}

	movieResp, err := protoMovie.GetMovieById(ctx, &protoServices.IdRequest{Id: movieId})
	if err != nil {
		log.Printf("Error getting the movie: %v", err)
		return false
	}

	foundClient := client.NewClientWithId(int(clientResp.Id), clientResp.FirstName, clientResp.LastName, clientResp.Email)
	foundMovie := movie.NewMovieWithId(int(movieResp.Id), movieResp.Title, float64(movieResp.Price))

	if !IsFavorite(foundClient, foundMovie) {
		log.Printf("%s is not in %s's favorites", foundMovie.Title, foundClient.FirstName)
		return false
	}

	favorite.DeleteFavoriteMovie(foundClient, foundMovie)
	return true
}

// IsFavorite checks if a movie is in a client's favorites
func IsFavorite(c client.Client, m movie.Movie) bool {
	return favorite.IsFavorite(c, m)
}

func openConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	return conn, err
}

func closeConnection(conn *grpc.ClientConn) {
	func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing connection: ", err)
		}
	}(conn)
}
