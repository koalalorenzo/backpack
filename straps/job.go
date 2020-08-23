package straps

import (
	"github.com/hashicorp/nomad/api"
)

// StrapJob is used to store the HCL code and the Job evaluation ID once it is
// registered in Nomad via the API.
type StrapJob struct {
	// JobEvalID is the Job ID once registered in Nomad
	JobEvalID string
	// Code is the HCL code of the Job
	Code string
}

// IsValid that is the question!
func (sj *StrapJob) IsValid(c *api.Client) bool {
	japi := c.Jobs()
	job, err := japi.ParseHCL(sj.Code, true)
	if err != nil {
		return false
	}

	res, _, err := japi.Validate(job, nil)
	if err != nil {
		return false
	}

	return len(res.Error) == 0 && len(res.Warnings) == 0 && len(res.ValidationErrors) == 0
}

// Run is parsing the HCL code and registering the Job into Nomad
func (sj *StrapJob) Run(c *api.Client) (err error) {
	japi := c.Jobs()
	job, err := japi.ParseHCL(sj.Code, true)
	if err != nil {
		return err
	}

	jobResponse, _, err := japi.Register(job, nil)
	if err != nil {
		return err
	}

	sj.JobEvalID = jobResponse.EvalID
	return nil
}
