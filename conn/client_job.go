package conn

import (
	"github.com/hashicorp/nomad/api"
)

// GetJobFromCode is parsing the HCL code and providing a api.Job{} pointer
func (co *Client) GetJobFromCode(code string) (*api.Job, error) {
	return co.jobs.ParseHCL(code, true)
}

func (co *Client) GetJobStatus(jobId string) (j *api.Job, err error) {
	j, _, err = co.jobs.Info(jobId, nil)
	return
}
