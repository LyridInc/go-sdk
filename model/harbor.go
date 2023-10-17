package model

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type IHarborClient interface {
	ListProjects() ([]byte, error)
	GetProject(projectName string) ([]byte, error)
	GetProjectRepositories(projectName string) ([]byte, error)
	GetRepository(projectName, repositoryName string) ([]byte, error)
	GetRepositoryArtifacts(projectName, repositoryName, tag string) ([]byte, error)
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

func (c *HarborClient) GetRepositoryArtifacts(projectName, repositoryName, tag string) ([]byte, error) {
	path := "/projects/" + projectName + "/repositories/" + repositoryName + "/artifacts?"
	if tag != "" {
		path = path + "q=tags%253D~" + tag
	}
	request, err := http.NewRequest("GET", c.BaseURL+path, nil)
	if err != nil {
		return nil, err
	}

	return c.doHttpRequest(request)
}
