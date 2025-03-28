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

func (csp *ClientProtoServiceServer) GetClientByEmail(ctx context.Context, request *protoServices.EmailRequest) (*protoServices.ClientResponse, error) {
	log.Println("GetClientByEmail from ClientProtoServiceServer called with email: ", request.Email)
	foundClient, err := client.GetClientByEmail(request.Email)

	if err != nil {
		log.Println("ClientProtoServiceServer: An error occured while searching for client email %s", err)
		return nil, err
	}

	log.Println("GetClientByEmail from ClientProtoServiceServer found this Client: ", foundClient)
	return &protoServices.ClientResponse{
		Id:        int32(foundClient.Id),
		FirstName: foundClient.FirstName,
		LastName:  foundClient.LastName,
		Email:     foundClient.Email,
	}, nil
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

func (csp *ClientProtoServiceServer) AddClient(ctx context.Context, request *protoServices.ClientRequest) (*protoServices.ClientResponse, error) {
	log.Println("AddClient from ClientProtoServiceServer called with request: ", request)
	_, err := client.GetClientByEmail(request.Email)
	if err == nil {
		log.Printf("ClientProtoServiceServer: Client with email %s already exists\n", request.Email)
		return nil, nil
	}
	addedClient := client.AddClient(client.NewClient(request.FirstName, request.LastName, request.Email))

	return &protoServices.ClientResponse{
		Id:        int32(addedClient.Id),
		FirstName: addedClient.FirstName,
		LastName:  addedClient.LastName,
		Email:     addedClient.Email,
	}, nil
}

func (csp *ClientProtoServiceServer) UpdateClient(ctx context.Context, request *protoServices.ClientRequest) (*protoServices.ClientResponse, error) {
	log.Println("UpdateClient from ClientProtoServiceServer called with request: ", request)
	foundClient, _ := client.GetClientById(int(request.Id))
	foundClient.FirstName = request.FirstName
	foundClient.LastName = request.LastName
	foundClient.Email = request.Email

	client.UpdateClient(foundClient)

	return &protoServices.ClientResponse{
		Id:        int32(foundClient.Id),
		FirstName: foundClient.FirstName,
		LastName:  foundClient.LastName,
		Email:     foundClient.Email,
	}, nil
}

func (csp *ClientProtoServiceServer) DeleteClient(ctx context.Context, request *protoServices.IdRequest) (*protoServices.StatusResponse, error) {
	log.Println("DeleteClient from ClientProtoServiceServer called with request: ", request)
	foundClient, err := client.GetClientById(int(request.Id))
	client.DeleteClient(foundClient)

	if err != nil {
		return &protoServices.StatusResponse{
			Success: false,
		}, err
	}

	return &protoServices.StatusResponse{
		Success: true,
	}, nil
}

func (csp *ClientProtoServiceServer) GetAllClients(empty *protoServices.EmptyRequest, stream protoServices.ClientProtoService_GetAllClientsServer) error {
	log.Println("GetAllClients from ClientProtoServiceServer called")
	clients := client.GetAllClients()

	for _, c := range clients {
		clientResponse := &protoServices.ClientResponse{
			Id:        int32(c.Id),
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Email:     c.Email,
		}

		if err := stream.Send(clientResponse); err != nil {
			return err
		}
	}

	return nil
}
