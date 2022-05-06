package depgraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseList(t *testing.T) {
	gr, rgr := getDepGraphAndReversed(t)
	rev := gr.Reverse()

	fmt.Println("##### expected", rgr)
	fmt.Println("##### reversed", rev)
}

func getDepGraphAndReversed(t *testing.T) (adj, rev Graph) {
	ldr := NewFixtureLoader("simple")
	adj, err := LoadDependencyGraph(ldr)
	require.NoError(t, err)
	rev, err = LoadDependencyGraph(NewFixtureLoader("simple-reversed"))
	require.NoError(t, err)
	return
}

// func getReverseDepGraph() Graph {
// 	return Graph{
// 		"logger": {
// 			"event",
// 			"gateway",
// 			"load-balancer",
// 			"http-srv",
// 		},
// 		"metrics": {
// 			"gateway",
// 			"load-balancer",
// 			"http-srv",
// 		},
// 		"event": {
// 			"rest-api",
// 		},
// 		"rest-api": {
// 			"http-srv",
// 		},
// 		"gateway":       {},
// 		"load-balancer": {},
// 		"http-srv":      {},
// 	}
// }
