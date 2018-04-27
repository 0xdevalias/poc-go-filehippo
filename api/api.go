package api

import "time"

func (c *Client) ProgramDefinitions() (*ProgramDefinitions, error) {
	req := c.restyClient.R().
		SetResult(ProgramDefinitions{})
		//SetHeader("AccessToken", c.accessTokenProducer()).
		//SetHeader("RequestTime", time.Now().UTC().Format(time.RFC3339Nano)).

	// Trying to control for header case..
	req.Header["AccessToken"] = append(req.Header["AccessToken"], c.accessTokenProducer())
	req.Header["RequestTime"] = append(req.Header["RequestTime"], time.Now().UTC().Format(time.RFC3339Nano))

	r, err := req.Get("/api/v1/ProgramDefinitions")
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
		SetHeader("AccessToken", c.accessTokenProducer()).
		SetHeader("RequestTime", time.Now().UTC().Format(time.RFC3339Nano)).
		SetHeader("Expect", "100-continue").
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
		SetHeader("AccessToken", c.accessTokenProducer()).
		SetHeader("RequestTime", time.Now().UTC().Format(time.RFC3339Nano)).
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
