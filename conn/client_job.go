package conn

import (
	"github.com/hashicorp/nomad/api"
)

// Run is parsing the HCL code and registering the Job into Nomad
func (co *Client) GetJob(code string) (*api.Job, error) {
	return co.c.Jobs().ParseHCL(code, true)
}
