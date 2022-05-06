package planner

import (
	"fmt"
	"io"

	"github.com/cedricmar/updep/pkg/cli/color"
	"github.com/cedricmar/updep/pkg/depgraph"
	"github.com/cedricmar/updep/pkg/tag"
)

type step []string

type converter struct {
	deps  depgraph.Graph
	steps []step
}

func Fprint(w io.Writer, steps []step) {
	for i, step := range steps {
		fmt.Fprintln(w, color.Green)
		fmt.Fprintln(w, "STEP:", i+1)
		fmt.Fprintln(w, "Dependencies to update:")
		fmt.Fprint(w, color.Reset)
		//fmt.Fprint(w, color.Blue)
		for _, k := range step {
			fmt.Fprintln(w, " -", k)
		}
		//fmt.Fprint(w, color.Reset)
	}
}

func Plan(key string, g depgraph.Graph) ([]step, error) {
	// First we check the graph for circular dependencies
	if hasCycles, cycles := g.HasCycles(); hasCycles {
		return []step{}, fmt.Errorf("dependency graph has cycles: %v", cycles)
	}
	// @TODO

	// We need to reverse the dependency graph to traverse it
	gr := g.Reverse()

	t, _ := tag.NewTag(key)

	conv := converter{
		deps:  gr,
		steps: []step{},
	}

	// We traverse the tree to generate an update plan
	conv.contraintsPush(t.Name)

	return conv.steps, nil
}

func (c *converter) contraintsPush(name string) {
	num := 0
	steps := make(map[string]int)
	n, ok := c.deps[name]
	if !ok {
		return
	}
	numSteps := c.contraintsPushRec(num, n, steps, 0)

	gSteps := gatherSteps(steps, numSteps)
	c.steps = cleanupSteps(gSteps)
}

func (c *converter) contraintsPushRec(num int, deps []string, steps map[string]int, maxSteps int) int {
	num++
	for _, dep := range deps {
		if num > steps[dep] { // We push to the highest step
			steps[dep] = num
			if num > maxSteps {
				maxSteps = num
			}
		}
		if len(c.deps[dep]) > 0 {
			maxSteps = c.contraintsPushRec(num, c.deps[dep], steps, maxSteps)
		}
	}
	return maxSteps
}

func gatherSteps(steps map[string]int, num int) []step {
	plan := make([]step, num)
	for svc, st := range steps {
		plan[st-1] = append(plan[st-1], svc)
	}
	return plan
}

func cleanupSteps(steps []step) []step {
	cl := []step{}
	for _, step := range steps {
		if len(step) > 0 {
			cl = append(cl, step)
		}
	}
	return cl
}

// // descentDepGraph issues a correct plan, but not an optimal one
// func (c *converter) descentDepGraph(num int, deps []string) step {
// 	s := step{}

// 	num++
// 	// Collect all deps
// 	for _, dep := range deps {
// 		if len(c.deps[dep]) == 0 {
// 			continue
// 		}
// 		nexts := c.descentDepGraph(num, c.deps[dep])
// 		for _, n := range nexts {
// 			c.seen[n] = struct{}{}
// 		}
// 	}
// 	for _, dep := range deps {
// 		if _, ok := c.seen[dep]; !ok {
// 			s = append(s, dep)
// 		}
// 	}

// 	if len(s) > 0 {
// 		// Prepend
// 		c.steps = append([]step{s}, c.steps...)
// 	}

// 	return s
// }
