package github

import "github.com/cedricmar/updep/pkg/repos"

func (gh *GithubAdapter) UpdateDepRegistry(reg repos.DepStorer) error {
	// @TODO - implement

	// IF .config has an ORG (or several)
	orgs, err := gh.GetOrganizations()
	if err != nil {
		return err
	}

	// @TODO - if more than 1 organization maybe specify
	// in the cli command, or loop through all of them and
	// get the repos

	if len(orgs) > 0 {

	} else {

	}

	// List organization repositories
	// GET /orgs/{org}/repos
	// https://docs.github.com/en/rest/repos/repos#list-organization-repositories

	// ELSE

	// List repositories for the authenticated user
	// GET /user/repos
	// https://docs.github.com/en/rest/repos/repos#list-repositories-for-the-authenticated-user

	// FOR all repositories

	// Check if it is there locally (and up-to-date) or go fetch Github:
	// We need to figure out in which language it is written
	// Then we get its version (tag)
	// With these informations we update the Registry record for this repo

	return nil
}
