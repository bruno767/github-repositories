package clients

import (
	"../cache"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GitRepository struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Full_name   string `json:"full_name"`
	Private     bool   `json:"private"`
	Description string `json:"description"`
	Html_url    string `json:"html_url"`
	Language    string `json:"language"`
}

type GitCommit struct {
	NodeId   string `json:"node_id"`
	Commit   Commit `json:"commit"`
	Html_url string `json:"html_url"`
}

type Commit struct {
	Author  Author `json:"author"`
	Message string `json:"message"`
}

type Author struct {
	Name string `json:"name"`
}

type GithubClient struct {
	Client          http.Client
	TtlCache        *cache.TtlCache
	GitRepositories *[]GitRepository
	GitCommits      *[]GitCommit
	Url             string
}

func (gc *GithubClient) GetRepositories() ([]GitRepository, error) {

	req, err := http.NewRequest(http.MethodGet, gc.Url+"/users/facebook/repos", nil)
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

func (gc *GithubClient) GetCommits(repo string) ([]GitCommit, error) {

	req, err := http.NewRequest(http.MethodGet, gc.Url+"/repos/facebook/"+repo+"/commits", nil)
	if err != nil {
		return nil, err
	}
	resp, err := gc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response recieved from git| status: %v| body: %v ", resp.StatusCode, gc.GitRepositories)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &gc.GitCommits)

	if err != nil {
		return nil, fmt.Errorf("failed reading response body : %s", err.Error())
	}

	return *gc.GitCommits, err
}
