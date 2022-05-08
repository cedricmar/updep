package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cedricmar/updep/pkg/tagger"
	"github.com/cedricmar/updep/pkg/updater"
)

func main() {
	fs := updater.Flags{}
	fs.Info = flag.Bool("i", false, "info mode")

	flag.Usage = updater.Usage
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprint(os.Stderr, "Error: Too many tags\n\n")
		flag.Usage()
		os.Exit(1)
	}

	t, err := tagger.NewTag(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
		flag.Usage()
		os.Exit(1)
	}
	fs.Tag = t

	updater.Update(fs)
}
