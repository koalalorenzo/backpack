package conn

import (
	"github.com/hashicorp/nomad/api"
)

// Plan is parsing the HCL code and registering the Job into Nomad
func (co *Client) Plan(job *api.Job, diff bool) (plan *api.JobPlanResponse, err error) {
	plan, _, err = co.jobs.Plan(job, diff, nil)
	if err != nil {
		return nil, err
	}

	return plan, nil
}
