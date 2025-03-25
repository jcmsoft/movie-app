module com.movie.app/data-access/client

go 1.24.1

replace com.movie.app/data-access/handler => ../handler

require com.movie.app/data-access/handler v0.0.0-00010101000000-000000000000

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.0 // indirect
)
