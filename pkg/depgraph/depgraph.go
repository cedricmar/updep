package depgraph

type GraphBuilder interface {
	BuildGraph() (Graph, error)
}

type Graph map[string][]string

func LoadDependencyGraph(repo GraphBuilder) (Graph, error) {
	gr, err := repo.BuildGraph()
	if err != nil {
		return Graph{}, err
	}
	return gr, nil
}
