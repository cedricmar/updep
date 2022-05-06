package depgraph

func (g Graph) Reverse() Graph {
	rev := Graph{}

	for k, deps := range g {
		if _, ok := rev[k]; !ok {
			rev[k] = []string{}
		}
		for _, dep := range deps {
			if _, ok := rev[dep]; !ok {
				rev[dep] = []string{}
			}
			rev[dep] = append(rev[dep], k)
		}
	}

	return rev
}

func (g Graph) HasCycles() (bool, [][]string) {
	sccs := g.kosaraju()

	hasCycles := false
	res := [][]string{}
	for _, scc := range sccs {
		if len(scc) > 1 {
			hasCycles = true
			res = append(res, scc)
		}
	}
	return hasCycles, res
}

type rec struct {
	gr   Graph
	rgr  Graph
	seen map[string]struct{}
	res  map[string][]string
}

// Kosaraju algorithm
func (g Graph) kosaraju() Graph {
	r := rec{
		gr:   g,
		rgr:  g.Reverse(),
		seen: map[string]struct{}{},
		res:  Graph{},
	}

	lst := []string{}

	for u := range g {
		r.visit(u, &lst)
	}

	r.seen = map[string]struct{}{}
	for _, u := range lst {
		r.assign(u, u)
	}

	return r.res
}

func (r *rec) visit(u string, lst *[]string) {
	if _, ok := r.seen[u]; !ok {
		r.seen[u] = struct{}{}
		for _, v := range r.gr[u] {
			r.visit(v, lst)
			prependStr(lst, u)
		}
	}
}

func (r *rec) assign(u string, root string) {
	if _, ok := r.seen[u]; !ok {
		r.seen[u] = struct{}{}
		r.res[root] = append(r.res[root], u)
		for _, v := range r.rgr[u] {
			r.assign(v, root)
		}
	}
}

func prependStr(s *[]string, v string) {
	*s = append(*s, "")
	copy((*s)[1:], *s)
	(*s)[0] = v
}
