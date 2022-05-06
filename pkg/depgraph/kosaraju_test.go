package depgraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKosaraju(t *testing.T) {
	assert.Equal(t,
		Graph{"1": {"1", "0", "2"}, "3": {"3"}},
		graph.kosaraju(),
	)
}

func TestKosarajuNoSCC(t *testing.T) {
	sccs := noscc.kosaraju()

	fmt.Println(sccs)
}

func TestKosarajuComplex(t *testing.T) {
	exp := Graph{
		"1":  {"1", "0", "3", "4", "2", "5", "6"},
		"7":  {"7", "9", "8"},
		"10": {"10"},
		"11": {"11"},
		"12": {"12"},
	}
	assert.Equal(t, exp, graph_complex.kosaraju())
}

func TestReverse(t *testing.T) {
	assert.Equal(t, reversed, graph.Reverse())
}

var graph = Graph{
	"0": {"1"},
	"1": {"2", "3"},
	"2": {"0"},
	"3": {"3"},
}

var reversed = Graph{
	"0": {"2"},
	"1": {"0"},
	"2": {"1"},
	"3": {"1", "3"},
}

var noscc = Graph{
	"0": {"1", "2", "3"},
	"1": {"4", "5"},
	"2": {"6", "7"},
	"3": {"7"},
	"4": {},
	"5": {},
	"6": {"8"},
	"7": {},
	"8": {},
}

var graph_complex = Graph{
	"0":  {"1"},
	"1":  {"2"},
	"2":  {"4", "6"},
	"3":  {"0", "5"},
	"4":  {"3"},
	"5":  {"4", "7"},
	"6":  {"4", "7"},
	"7":  {"8", "10"},
	"8":  {"9", "11"},
	"9":  {"7"},
	"10": {},
	"11": {"12"},
	"12": {},
}
