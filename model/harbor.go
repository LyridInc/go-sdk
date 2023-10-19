package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Artifact struct {
	ID                int          `json:"id"`
	Accessories       interface{}  `json:"accessories"`
	AdditionLinks     AdditionLink `json:"addition_links"`
	Digest            string       `json:"digest"`
	ExtraAttrs        Attribute    `json:"extra_attrs"`
	Icon              string       `json:"icon"`
	Labels            interface{}  `json:"labels"`
	ManifestMediaType string       `json:"manifest_media_type"`
	MediaType         string       `json:"media_type"`
	ProjectID         int          `json:"project_id"`
	PullTime          *time.Time   `json:"pull_time"`
	PushTime          *time.Time   `json:"push_time"`
	RepositoryID      int          `json:"repository_id"`
	Size              uint64       `json:"size"`
	Tags              []BuildTag   `json:"tags"`
	Type              string       `json:"type"`
}

type AdditionLink struct {
	BuildHistory    Link `json:"build_history"`
	Vulnerabilities Link `json:"vulnerabilities"`
}

type Link struct {
	Absolute bool   `json:"absolute"`
	Href     string `json:"href"`
}

type Attribute struct {
	Architecture string      `json:"architecture"`
	Author       string      `json:"author"`
	Config       interface{} `json:"config"`
	Created      *time.Time  `json:"created"`
	OS           string      `json:"os"`
}

type BuildTag struct {
	ArtifactID   int        `json:"artifact_id"`
	ID           int        `json:"id"`
	Immutable    bool       `json:"immutable"`
	Name         string     `json:"name"`
	PullTime     *time.Time `json:"pull_time"`
	PushTime     *time.Time `json:"push_time"`
	RepositoryID int        `json:"repository_id"`
}

type IHarborClient interface {
	ListProjects() ([]byte, error)
	GetProject(projectName string) ([]byte, error)
	GetProjectRepositories(projectName string) ([]byte, error)
	GetRepository(projectName, repositoryName string) ([]byte, error)
	GetRepositoryArtifacts(projectName, repositoryName, tag string) (*[]Artifact, error)
}

type HarborClient struct {
	BaseURL  string
	Username string
	Password string
}

func NewHarborClient(baseURL, username, password string) IHarborClient {
	return &HarborClient{
		BaseURL:  baseURL,
		Username: username,
		Password: password,
	}
}

func (c *HarborClient) doHttpRequest(request *http.Request) ([]byte, error) {
	request.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.Username+":"+c.Password)))

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("error %d: resource not found for %s", response.StatusCode, request.URL)
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	return b, err
}

func (c *HarborClient) ListProjects() ([]byte, error) {
	request, err := http.NewRequest("GET", c.BaseURL+"/projects", nil)
	if err != nil {
		return nil, err
	}
	return c.doHttpRequest(request)
}

func (c *HarborClient) GetProject(projectName string) ([]byte, error) {
	request, err := http.NewRequest("GET", c.BaseURL+"/projects/"+projectName, nil)
	if err != nil {
		return nil, err
	}
	return c.doHttpRequest(request)
}

func (c *HarborClient) GetProjectRepositories(projectName string) ([]byte, error) {
	request, err := http.NewRequest("GET", c.BaseURL+"/projects/"+projectName+"/repositories", nil)
	if err != nil {
		return nil, err
	}
	return c.doHttpRequest(request)
}

func (c *HarborClient) GetRepository(projectName, repositoryName string) ([]byte, error) {
	request, err := http.NewRequest("GET", c.BaseURL+"/projects/"+projectName+"/repositories/"+repositoryName, nil)
	if err != nil {
		return nil, err
	}
	return c.doHttpRequest(request)
}

func (c *HarborClient) GetRepositoryArtifacts(projectName, repositoryName, tag string) (*[]Artifact, error) {
	path := "/projects/" + projectName + "/repositories/" + repositoryName + "/artifacts?"
	if tag != "" {
		path = path + "q=tags%253D~" + tag
	}
	request, err := http.NewRequest("GET", c.BaseURL+path, nil)
	if err != nil {
		return nil, err
	}

	b, err := c.doHttpRequest(request)
	if err != nil {
		return nil, err
	}

	artifacts := []Artifact{}
	json.Unmarshal(b, &artifacts)

	return &artifacts, nil
}
