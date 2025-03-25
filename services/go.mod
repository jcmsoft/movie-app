module com.movie.app/services

go 1.24.1

replace com.movie.app/data-access/client => ../data-access/client

replace com.movie.app/data-access/handler => ../data-access/handler

require com.movie.app/data-access/client v0.0.0-00010101000000-000000000000

require (
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.5
)
