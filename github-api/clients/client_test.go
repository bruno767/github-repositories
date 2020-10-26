package clients

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGithubClient_GetRepositories_200(t *testing.T) {
	// given
	givenGitRepositories := []GitRepository{
		{
			Id:          123,
			Name:        "testName",
			Full_name:   "testFullName",
			Private:     false,
			Description: "TestDesc",
			Html_url:    "www.someurl.com",
			Language:    "go",
		},
	}
	handler := http.NewServeMux()
	handler.HandleFunc("/users/facebook/repos", func(w http.ResponseWriter, r *http.Request) {

		value, _ := json.Marshal(givenGitRepositories)
		_, _ = w.Write(value)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	client := GithubClient{Url: srv.URL}

	// when
	gitRepositories, err := client.GetRepositories()
	if err != nil {
		t.Error(err)
	}

	// then
	if !reflect.DeepEqual(givenGitRepositories, gitRepositories) {
		t.Errorf("Error the response: %v \n didnt match with the given: %v", gitRepositories, givenGitRepositories)
	}
}

func TestGithubClient_GetRepositories_MissingFields_200(t *testing.T) {
	// given
	givenGitRepositories := []GitRepository{
		{
			Id:        123,
			Name:      "testName",
			Full_name: "testFullName",
			Private:   false,
		},
	}
	handler := http.NewServeMux()
	handler.HandleFunc("/users/facebook/repos", func(w http.ResponseWriter, r *http.Request) {

		value, _ := json.Marshal(givenGitRepositories)
		_, _ = w.Write(value)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	client := GithubClient{Url: srv.URL}

	// when
	gitRepositories, err := client.GetRepositories()
	if err != nil {
		t.Error(err)
	}

	// then
	if !reflect.DeepEqual(givenGitRepositories, gitRepositories) {
		t.Errorf("Error the response: %v \n didnt match with the given: %v", gitRepositories, givenGitRepositories)
	}
}

func TestGithubClient_GetRepositories_EmptyReturn_200(t *testing.T) {
	// given
	handler := http.NewServeMux()
	handler.HandleFunc("/users/facebook/repos", func(w http.ResponseWriter, r *http.Request) {

		value, _ := json.Marshal([]GitRepository{})
		_, _ = w.Write(value)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	client := GithubClient{Url: srv.URL}

	// when
	gitRepositories, err := client.GetRepositories()
	if err != nil {
		t.Error(err)
	}

	// then
	if !reflect.DeepEqual([]GitRepository{}, gitRepositories) {
		t.Errorf("Error the response didnt match with the given: %v", gitRepositories)
	}
}

func TestGithubClient_GetRepositories_BadRequest(t *testing.T) {
	// given
	handler := http.NewServeMux()
	handler.HandleFunc("/users/facebook/repos", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		value, _ := json.Marshal([]GitRepository{})
		_, _ = w.Write(value)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	client := GithubClient{Url: srv.URL}

	// when
	_, err := client.GetRepositories()

	// then
	if err == nil {
		t.Error("We should have returned an error while requesting")
	}

}

func TestGithubClient_GetRepositories_WrongContract(t *testing.T) {
	// given
	handler := http.NewServeMux()
	handler.HandleFunc("/users/facebook/repos", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte{1, 2, 3})
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	client := GithubClient{Url: srv.URL}

	// when
	_, err := client.GetRepositories()

	// then
	if err == nil {
		t.Error("We should have returned an error while requesting")
	}

}

func TestGithubClient_GetCommits_200(t *testing.T) {
	// given
	givenGitCommits := []GitCommit{
		{
			NodeId: "123",
			Commit: Commit{
				Author:  Author{Name: "author"},
				Message: "message",
			},
			Html_url: "www.someurl.com",
		},
	}
	handler := http.NewServeMux()
	handler.HandleFunc("/repos/facebook/repo1/commits", func(w http.ResponseWriter, r *http.Request) {

		value, _ := json.Marshal(givenGitCommits)
		_, _ = w.Write(value)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	client := GithubClient{Url: srv.URL}

	// when
	gitCommits, err := client.GetCommits("repo1")
	if err != nil {
		t.Error(err)
	}

	// then
	if !reflect.DeepEqual(givenGitCommits, gitCommits) {
		t.Errorf("Error the response: %v \n didnt match with the given: %v", gitCommits, givenGitCommits)
	}
}
