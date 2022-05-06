package github

import "github.com/cedricmar/updep/pkg/depgraph"

type GithubAdapter struct{}

func (g GithubAdapter) BuildGraph() (depgraph.Graph, error) {
	// @TODO - implement
	return depgraph.Graph{}, nil
}
