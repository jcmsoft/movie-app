package client

import (
	"com.movie.app/data-access/handler"
	"fmt"
	"log"
)

type Client struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}

// NewClient creates a new Client with the given parameters
func NewClient(firstName, lastName, email string) Client {
	return Client{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

// NewClientWithId creates a new Client with the given parameters
func NewClientWithId(id int, firstName, lastName, email string) Client {
	return Client{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

// String returns a string representation of the Client
func (c Client) String() string {
	return fmt.Sprintf("Client{id: %d, FirstName: %s, LastName: %s, Email: %s}",
		c.Id, c.FirstName, c.LastName, c.Email)
}

func GetAllClients() (clients []Client) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	// Query the database
	rows, err := db.Query("SELECT * FROM client")
	if err != nil {
		log.Fatalf("An error occured while querying the database %s", err)
	}
	// Iterate over the rows
	for rows.Next() {
		var client Client
		err = rows.Scan(&client.Id, &client.FirstName, &client.LastName, &client.Email)
		if err != nil {
			log.Fatalf("An error occured while scanning the rows %s", err)
		}
		clients = append(clients, client)
	}
	return clients
}

func GetClientById(id int) (Client, error) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	var client Client
	// Query the database
	err := db.QueryRow("SELECT * FROM client WHERE Id = ?", id).Scan(&client.Id, &client.FirstName, &client.LastName, &client.Email)
	if err != nil {
		log.Fatalf("An error occured while querying the database %s", err)
		return Client{}, err
	}
	return client, nil
}

func GetClientByEmail(email string) (Client, error) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	var client Client
	// Query the database
	err := db.QueryRow("SELECT * FROM client WHERE email = ?", email).Scan(&client.Id, &client.FirstName, &client.LastName, &client.Email)
	if err != nil {
		// Return an empty client and the error if no rows found or other error
		return Client{}, err
	}

	return client, nil
}

func AddClient(client Client) Client {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	// Insert the client into the database
	result, err := db.Exec("INSERT INTO client (first_name, last_name, email) VALUES (?, ?, ?)", client.FirstName, client.LastName, client.Email)
	if err != nil {
		log.Fatalf("An error occurred while inserting the client into the database: %s", err)
	}

	// Get the ID of the newly inserted client
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("An error occurred while getting the last insert ID: %s", err)
	}

	client.Id = int(id)
	return client
}

func IsClientExist(email string) (bool, Client) {
	client, err := GetClientByEmail(email)

	if err == nil {
		log.Println("This client is already registered")
		return true, client
	}

	return false, Client{}
}

func UpdateClient(client Client) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)

	// Update the client in the database
	_, err := db.Exec("UPDATE client SET first_name = ?, last_name = ?, email = ? WHERE Id = ?", client.FirstName, client.LastName, client.Email, client.Id)
	if err != nil {
		log.Fatalf("An error occured while updating the client in the database %s", err)
	}
}

func DeleteClient(client Client) {
	db := handler.GetDatabaseConnection()
	defer handler.CloseDBConnection(db)
	// Delete the client from the database
	_, err := db.Exec("DELETE FROM client WHERE Id = ?", client.Id)
	if err != nil {
		log.Fatalf("An error occured while deleting the client from the database %s", err)
	}
}
