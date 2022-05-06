package cli

import (
	"os/exec"
	"strings"
)

type Commander interface {
	GetLatestTag() (string, error)
	Tag(v string) error
}

func RunCommand(cmd string) (string, error) {
	args := explodeCMD(cmd)

	run := exec.Command(args[0], args[1:]...)
	stdout, err := run.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}

func explodeCMD(cmd string) []string {
	return strings.Split(cmd, " ")
}
