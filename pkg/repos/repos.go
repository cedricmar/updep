package repos

import "github.com/cedricmar/updep/pkg/depgraph"

type RepositoryAdapter interface {
	depgraph.GraphBuilder
}
