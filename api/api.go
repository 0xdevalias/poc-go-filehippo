package api

func (c *Client) ProgramDefinitions() (*ProgramDefinitions, error) {
	r, err := c.restyClient.R().
		SetResult(ProgramDefinitions{}).
		Get("/api/v1/ProgramDefinitions")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*ProgramDefinitions), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

func (c *Client) ScanResults(scanRequest ScanRequest) (*ScanResults, error) {
	r, err := c.restyClient.R().
		SetBody(scanRequest).
		SetResult(ScanResults{}).
		Post("/api/v1/ScanResults")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*ScanResults), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

func (c *Client) InstallerPrograms(irReq InstallerProgramsRequest) (*InstallerProgramsResponse, error) {
	r, err := c.restyClient.R().
		SetBody(irReq).
		SetResult(InstallerProgramsResponse{}).
		Post("/api/v1/InstallerPrograms")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*InstallerProgramsResponse), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}
