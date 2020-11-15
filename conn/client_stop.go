package conn

// Stop stops a specific Job
func (co *Client) Stop(jobID string, purge bool) (evalID string, err error) {
	evalID, _, err = co.jobs.Deregister(jobID, purge, nil)
	if err != nil {
		return "", err
	}

	return evalID, nil
}
