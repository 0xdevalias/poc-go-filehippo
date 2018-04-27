package api

import (
	"fmt"

	"github.com/go-resty/resty"
)

const (
	BaseUrl string = "http://appmanager.filehippo.com/" // TODO: HTTPS?
)

type Client struct {
	restyClient *resty.Client

	clientID            string
	accessTokenProducer func() string
}

func DefaultClient(clientID string) *Client {
	c := resty.DefaultClient
	c.SetHostURL(BaseUrl)

	//c.SetHeader("User-Agent", "download_manager")
	//c.SetHeader("AppManagerVersion", "2.0.0.392")
	//c.SetHeader("ClientId", clientID)

	// Trying to control for header casing..
	c.Header["User-Agent"] = append(c.Header["User-Agent"], "download_manager")
	c.Header["AppManagerVersion"] = append(c.Header["AppManagerVersion"], "2.0.0.392")
	c.Header["ClientId"] = append(c.Header["ClientId"], clientID)

	client := Client{
		restyClient: c,
		clientID:    clientID,
	}
	client.accessTokenProducer = client.AccessToken

	return &client
}

func (c *Client) WithDebug() *Client {
	c.restyClient.SetDebug(true)
	return c
}

func (c *Client) WithProxy(proxyURL string) *Client {
	c.restyClient.SetProxy(proxyURL)
	return c
}

func errorResponseFormatter(resp *resty.Response) error {
	return fmt.Errorf("api returned an error response (%v): %s", resp.StatusCode(), string(resp.Body()))
}

func DefaultScanRequest(programs []Program) ScanRequest {
	return ScanRequest{
		Programs:   programs,
		ResVersion: "2",
		Parameters: ScanRequestParameters{
			HideBeta:     true,
			ShowDetailed: true,
			ShowAll:      true,
		},
		Time:          0.0,
		ClientVersion: "2.392",
		System: System{
			VersionString: "Microsoft Windows NT 6.2.9200.0",
			VersionMajor:  6,
			VersionMinor:  2,
			Revision:      0,
			ServicePack:   "",
			Bit:           64,
			IsServer:      false,
			GraphicsCards: []string{"Intel(R) HD Graphics 530", "NVIDIA GeForce GTX 1060"},
		},
		ScanType: "Manual",
	}
}
