package devbox

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
}

type ResolveRequest struct {
	Name    string
	Version string
}

type ResolveResponse struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Summary string `json:"summary,omitempty"`
	// Systems
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (client *Client) Resolve(request ResolveRequest) (*ResolveResponse, error) {
	u, err := url.Parse("https://search.devbox.sh/v2/resolve")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("name", request.Name)
	q.Add("version", request.Version)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	rsp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	jsonDecoder := json.NewDecoder(rsp.Body)
	defer func(r io.ReadCloser) {
		_ = r.Close()
	}(rsp.Body)

	var response ResolveResponse
	err = jsonDecoder.Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
