package server

import (
	"github.com/pachyderm/pachyderm/src/client/admin"
	"github.com/pachyderm/pachyderm/src/server/pkg/log"
)

// APIServer represents and APIServer
type APIServer interface {
	admin.APIServer
}

// NewAPIServer returns a new admin.APIServer
func NewAPIServer(address string) APIServer {
	return &apiServer{
		Logger:  log.NewLogger("admin.API"),
		address: address,
	}
}
