package devbox

import (
	"encoding/json"
	"fmt"
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

type PkgRequest struct {
	Name string
}

type ResolveResponse struct {
	Name    string            `json:"name,omitempty"`
	Version string            `json:"version,omitempty"`
	Summary string            `json:"summary,omitempty"`
	Systems map[string]System `json:"systems,omitempty"`
}

type System struct {
	FlakeInstallable FlakeInstallable `json:"flake_installable,omitempty"`
	LastUpdated      string           `json:"last_updated,omitempty"`
	Outputs          []Output         `json:"outputs,omitempty"`
}

type FlakeInstallable struct {
	Ref      Ref    `json:"ref,omitempty"`
	AttrPath string `json:"attr_path,omitempty"`
}

type Ref struct {
	Type  string `json:"type,omitempty"`
	Owner string `json:"owner,omitempty"`
	Repo  string `json:"repo,omitempty"`
	Rev   string `json:"rev,omitempty"`
}

type PkgResponse struct {
	Name        string    `json:"name,omitempty"`
	Summary     string    `json:"summary,omitempty"`
	HomepageUrl string    `json:"homepage_url,omitempty"`
	License     string    `json:"string,omitempty"`
	Releases    []Release `json:"releases,omitempty"`
}

type Release struct {
	Version          string     `json:"version,omitempty"`
	LastUpdated      string     `json:"last_updated,omitempty"`
	Platforms        []Platform `json:"platforms,omitempty"`
	PlatformsSummary string     `json:"platforms_summary,omitempty"`
	OutputsSummary   string     `json:"outputs_summary,omitempty"`
}

type Platform struct {
	Arch          string   `json:"arch,omitempty"`
	Os            string   `json:"os,omitempty"`
	System        string   `json:"system,omitempty"`
	AttributePath string   `json:"attribute_path,omitempty"`
	CommitHash    string   `json:"commit_hash,omitempty"`
	Date          string   `json:"date,omitempty"`
	Outputs       []Output `json:"outputs,omitempty"`
}

type Output struct {
	Name    string `json:"name,omitempty"`
	Path    string `json:"path,omitempty"`
	Default bool   `json:"default,omitempty"`
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

	var response ResolveResponse
	err = client.doRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *Client) Pkg(request PkgRequest) (*PkgResponse, error) {
	u, err := url.Parse("https://search.devbox.sh/v2/pkg")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("name", request.Name)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	var response PkgResponse
	err = client.doRequest(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *Client) doRequest(httpRequest *http.Request, response any) error {
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(httpResponse.Body)

	if httpResponse.StatusCode != 200 {
		return fmt.Errorf("http status %d", httpResponse.StatusCode)
	}

	jsonDecoder := json.NewDecoder(httpResponse.Body)
	defer func(r io.ReadCloser) {
		_ = r.Close()
	}(httpResponse.Body)

	err = jsonDecoder.Decode(&response)
	if err != nil {
		return err
	}

	return nil
}
