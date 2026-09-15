package initiator

import (
	"log"

	authGenerated "github.com/alazarbeyenenew2/common/generated/auth/auth"
	"github.com/alazarbeyenenew2/gateway/internal/constants/dto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCConnections struct {
	Auth *grpc.ClientConn
}

func InitGRPCConnections() *GRPCConnections {
	authClient, err := grpc.NewClient(viper.GetString("auth.grpc"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("unable to initialize client")
	}

	return &GRPCConnections{
		Auth: authClient,
	}
}

func InitGRPC() dto.GRPCClients {
	connections := InitGRPCConnections()
	return dto.GRPCClients{
		Auth: authGenerated.NewAuthServiceClient(connections.Auth),
	}

}
