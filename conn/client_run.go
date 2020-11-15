package conn

import (
	"github.com/hashicorp/nomad/api"
)

// Run is parsing the HCL code and registering the Job into Nomad
func (co *Client) Run(job *api.Job) (jr *api.JobRegisterResponse, err error) {
	jr, _, err = co.jobs.Register(job, nil)
	if err != nil {
		return nil, err
	}

	return jr, nil
}
