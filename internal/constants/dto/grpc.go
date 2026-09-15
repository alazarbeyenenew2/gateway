package dto

import (
	authGenerated "github.com/alazarbeyenenew2/common/generated/auth/auth"
)

type GRPCClients struct {
	Auth authGenerated.AuthServiceClient
}
