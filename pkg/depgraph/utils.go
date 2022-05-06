package depgraph

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type FixtureLoader struct {
	file string
}

func NewFixtureLoader(name string) FixtureLoader {
	return FixtureLoader{
		file: "../fixtures/" + name + ".yml",
	}
}

func LoadDependencyGraphFromFixture(name string) (Graph, error) {
	ldr := NewFixtureLoader(name)
	return ldr.BuildGraph()
}

func (f FixtureLoader) BuildGraph() (Graph, error) {
	dat, err := ioutil.ReadFile(f.file)
	if err != nil {
		return Graph{}, err
	}

	g := Graph{}
	err = yaml.Unmarshal(dat, &g)
	if err != nil {
		return Graph{}, err
	}

	return g, nil
}
