package updater

import (
	"fmt"

	"github.com/cedricmar/updep/pkg/tagger"
)

func DepUpdate(t tagger.Tag) error {
	// @TODO - find the type somehow
	kind := "go"
	// Go
	if kind == "go" {
		return goUp(t)
	}
	// Java
	if kind == "java" {
		return javaUp(t)
	}
	return fmt.Errorf("no strategy found to update %s", t.Key)
}

func goUp(t tagger.Tag) error {
	return nil
}

func javaUp(t tagger.Tag) error {
	return nil
}
