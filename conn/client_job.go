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

func (co *Client) GetJobAllocations(jobId string) (alls []*api.Allocation, err error) {
	allList, _, err := co.jobs.Allocations(jobId, false, nil)
	if err != nil {
		return
	}

	alls = []*api.Allocation{}
	for _, al := range allList {
		alloc, _, err := co.alloc.Info(al.ID, nil)
		if err != nil {
			return []*api.Allocation{}, err
		}
		alls = append(alls, alloc)
	}
	return
}
