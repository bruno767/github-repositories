package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GitRepository struct {
	Id          int                    `json:"id"`
	Name        string                 `json:"name"`
	Full_name   string                 `json:"full_name"`
	Private     bool                   `json:"private"`
	Description string                 `json:"description"`
	Git_url     string                 `json:"git_url"`
	Language    string                 `json:"language"`
}

type GithubClient struct {
	Client          http.Client
	GitRepositories *[]GitRepository
	Url             string
}

func (gc *GithubClient) GetRepositories() ([]GitRepository, error) {

	req, err := http.NewRequest(http.MethodGet, gc.Url + "/users/facebook/repos", nil)
	if err != nil {
		return nil, err
	}
	resp, err := gc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &gc.GitRepositories)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body : %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response recieved from git| status: %v| body: %v ", resp.StatusCode, gc.GitRepositories)
	}

	return *gc.GitRepositories, err
}
