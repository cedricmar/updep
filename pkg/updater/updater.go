package updater

import (
	"flag"
	"fmt"
	"os"

	"github.com/cedricmar/updep/pkg/cli"
	"github.com/cedricmar/updep/pkg/planner"
	"github.com/cedricmar/updep/pkg/repos/github"
	"github.com/cedricmar/updep/pkg/tag"
)

// @TODO - move elsewhere
const (
	DevRoot = "~/dev/my-project" // local
	GitPath = "github.com/cedricmar"
)

type Flags struct {
	Info *bool
	Tag  tag.Tag
}

func Update(f Flags) {
	tag := f.Tag

	gha := github.GithubAdapter{}
	g, err := gha.BuildGraph()
	if err != nil {
		fmt.Fprint(os.Stderr, "Could not generate the graph from provided github config:\n\n")
		os.Exit(1)
	}

	// Building the update plan
	plan := planner.Plan(tag.Key, g)

	// Dry run?
	if *f.Info {
		planner.Fprint(os.Stdout, plan)
		os.Exit(0)
	}

	// @TODO - run the real thing

	tnum := cli.StringPrompt("Ticket number:")
	tname := cli.StringPrompt("Ticket name:")

	fmt.Fprintf(os.Stdout, "[%s] %s\n", tnum, tname)

	// - Prompt for MR name (derive the branch name from this)
	//
	// FOR deps in this step:
	//     1) Clone and cd into the folder
	//     2) Create branch and update the code (+ push)
	//        - Go strategy
	//        - Java strategy
	//        - C# strategy
	//        - etc...
	//     3) Create the MR
	//     4) Reset
	//        - go back to develop branch
	//        - cd back to dev root

	// FOR deps in this step (again):
	//     5) Wait for the pipeline to run / merge of the MRs, then:
	//        - Suggest tags and prompt for tag names (memorize those)
	//        - Create MRs

	// Wait for the pipelines to run again
	// Now we can go to the next step of the plan
	// (we now might have to update more than one dependency)

	os.Exit(0)
}

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of updep:\n\n")
	fmt.Fprintf(os.Stderr, "  updep <options> <repo_name>@<version_tag>\n\n")
	fmt.Fprint(os.Stderr, "  Options:\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}
