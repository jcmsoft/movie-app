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

// AddClientFavoriteMovie adds a movie to a client's favorites
func AddClientFavoriteMovie(theClient client.Client, favoriteMovie movie.Movie) {
	conn, _ := OpenConnection()
	defer CloseConnection(conn)

	protoClient := protoServices.NewClientProtoServiceClient(conn)
	protoMovie := protoServices.NewMovieProtoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	clientResp, err := protoClient.IsClientExist(ctx, &protoServices.EmailRequest{Email: theClient.Email})
	if err != nil {
		log.Fatalf("Error calling IsClientExist: %v", err)
	}

	movieResp, err := protoMovie.IsMovieExist(ctx, &protoServices.IdRequest{Id: int32(favoriteMovie.Id)})
	if err != nil {
		log.Fatalf("Error calling IsMovieExist: %v", err)
	}

	if !clientResp.Exist || !movieResp.Exist {
		log.Fatalf("Client OR Movie does not exist")
	}

	if IsFavorite(theClient, favoriteMovie) {
		log.Fatalf("%s is already in %s's favorites", favoriteMovie.Title, theClient.FirstName)
	}

	favorite.AddFavoriteMovie(theClient, favoriteMovie)
}

// RemoveFavorite removes a movie from a client's favorites
func RemoveFavorite(c client.Client, m movie.Movie) {
	favorite.DeleteFavoriteMovie(c, m)
}

// GetFavoritesByClient retrieves all favorite movies for a specific client
func GetFavoritesByClient(c client.Client) []movie.Movie {
	return favorite.GetClientFavoriteMovies(c)
}

// GetFavoritesByMovie retrieves all clients who have favorited a specific movie
func GetFavoritesByMovie(m movie.Movie) []client.Client {
	panic("unimplemented")
}

// IsFavorite checks if a movie is in a client's favorites
func IsFavorite(c client.Client, m movie.Movie) bool {
	return favorite.IsFavorite(c, m)
}

func OpenConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	return conn, err
}

func CloseConnection(conn *grpc.ClientConn) {
	func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing connection: ", err)
		}
	}(conn)
}
