package deployer

import (
	"github.com/hashicorp/nomad/api"
	"golang.org/x/exp/errors/fmt"
)

var c *api.Client

func Deploy(hcl string) (err error) {
	c, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	japi := c.Jobs()
	job, err := japi.ParseHCL(hcl, true)
	if err != nil {
		return err
	}

	jobResponse, _, err := japi.Register(job, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Job ID: %s", jobResponse.EvalID)
	return nil
}
