module com.movie.app/services/clientService

go 1.24.1

replace com.movie.app/data-access/client => ../../data-access/client

replace com.movie.app/data-access/handler => ../../data-access/handler

require com.movie.app/data-access/client v0.0.0-00010101000000-000000000000

require (
	com.movie.app/data-access/handler v0.0.0-00010101000000-000000000000 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.0 // indirect
)
