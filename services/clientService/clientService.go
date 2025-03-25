package clientService

import (
	"com.movie.app/data-access/client"
	protoServices "com.movie.app/services/servicespb"
	"context"
	"log"
)

type ClientProtoServiceServer struct {
	protoServices.UnimplementedClientProtoServiceServer
}

func CreateClient(c client.Client) client.Client {
	return client.AddClient(c)
}

func EditClient(c client.Client) {
	client.UpdateClient(c)
}

func DeleteClient(c client.Client) {
	client.DeleteClient(c)
}

func GetClientByEmail(email string) (client.Client, error) {
	return client.GetClientByEmail(email)
}

func GetClients() (clients []client.Client) {
	return client.GetAllClients()
}

func (csp *ClientProtoServiceServer) GetClientById(ctx context.Context, request *protoServices.IdRequest) (*protoServices.ClientResponse, error) {
	log.Println("GetClientById from ClientProtoServiceServer called with request: ", request)
	foundClient, err := client.GetClientById(int(request.Id))

	if err != nil {
		log.Println("ClientProtoServiceServer: An error occured while searching for client ID %s", err)
		return nil, err
	}

	log.Println("GetClientById from ClientProtoServiceServer found this Client: ", foundClient)
	return &protoServices.ClientResponse{
		Id:        int32(foundClient.Id),
		FirstName: foundClient.FirstName,
		LastName:  foundClient.LastName,
		Email:     foundClient.Email,
	}, nil
}

func (csp *ClientProtoServiceServer) IsClientExist(ctx context.Context, request *protoServices.EmailRequest) (*protoServices.IsExistResponse, error) {
	log.Println("IsClientExist from ClientProtoServiceServer called with email: ", request.Email)
	exist, _ := client.IsClientExist(request.Email)

	return &protoServices.IsExistResponse{
		Exist: exist,
	}, nil
}
