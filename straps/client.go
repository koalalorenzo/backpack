package straps

import (
	"github.com/hashicorp/nomad/api"
)

// GetClient will provide a enw client with the default configuration
func GetClient() (*api.Client, error) {
	return api.NewClient(api.DefaultConfig())
}
