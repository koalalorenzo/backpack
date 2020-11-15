package conn

import (
	"github.com/hashicorp/nomad/api"
)

// Run is parsing the HCL code and registering the Job into Nomad
func (co *Client) Run(code string) (jr *api.JobRegisterResponse, err error) {
	japi := co.c.Jobs()
	job, err := japi.ParseHCL(code, true)
	if err != nil {
		return
	}

	jr, _, err = japi.Register(job, nil)
	if err != nil {
		return nil, err
	}

	return jr, nil
}
