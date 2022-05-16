package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetOrganizations returns the
func (gh *GithubAdapter) GetOrganizations() (orgs []Organization, err error) {
	// GET /organizations
	r, err := gh.httpGet("/organizations", &orgs)
	if err != nil {
		return
	}

	if _, ok := r.([]Organization); ok {
		err = errors.New("wrong type for organizations")
		return
	}

	return
}

func (gh *GithubAdapter) GetOrgRepositories(org string) (repos []Repository, err error) {
	// GET /orgs/{org}/repos
	url := fmt.Sprintf("/orgs/%s/repos", org)
	r, err := gh.httpGet(url, &repos)
	if err != nil {
		return
	}

	if _, ok := r.([]Repository); ok {
		err = errors.New("wrong type for repositories")
		return
	}

	return
}
func (gh *GithubAdapter) GetRepositories() (repos []Repository, err error) {
	// GET /user/repos
	r, err := gh.httpGet("/user/repos", &repos)
	if err != nil {
		return
	}

	if _, ok := r.([]Repository); ok {
		err = errors.New("wrong type for repositories")
		return
	}

	return
}

func (gh *GithubAdapter) httpGet(url string, ret interface{}) (interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, UrlBase+url, nil)
	if err != nil {
		return ret, err
	}
	req.Header.Set("Accept", HeaderAccept)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ret, err
	}
	defer res.Body.Close()

	dat, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(dat, ret)
	if err != nil {
		return ret, err
	}

	return ret, err
}
