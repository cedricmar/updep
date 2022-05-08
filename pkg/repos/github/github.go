package github

import (
	"github.com/cedricmar/updep/pkg/repos"
)

type GithubAdapter struct{}

func (g GithubAdapter) UpdateDepRegistry(reg repos.DepStorer) error {
	// @TODO - implement

	// IF .config has an ORG (or several)

	// List organization repositories
	// GET /orgs/{org}/repos
	// https://docs.github.com/en/rest/repos/repos#list-organization-repositories

	// ELSE

	// List repositories for the authenticated user
	// GET /user/repos
	// https://docs.github.com/en/rest/repos/repos#list-repositories-for-the-authenticated-user

	// FOR all repositories

	// We need to figure out in which language it is written
	// Then we get its version (tag)
	// With these informations we update the Registry record for this repo

	return nil
}
