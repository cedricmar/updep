package git

import (
	"fmt"
	"strings"

	"github.com/cedricmar/updep/pkg/cli"
)

func GetLatestTag() (string, error) {
	cmd := "git describe --abbrev=0 --tags"

	out, err := cli.RunCommand(cmd)
	if err != nil {
		//fmt.Fprintln(os.Stderr, err.Error())
		return "", err
	}

	out = strings.Trim(out, "\n")
	out = strings.Trim(out, " ")

	return out, nil
}

func Tag(v string) (string, error) {
	tagCMD := fmt.Sprintf("git tag -a %s -m \"tagging %s\"", v, v)

	_, err := cli.RunCommand(tagCMD)
	if err != nil {
		//fmt.Fprintln(os.Stderr, err.Error())
		return "", err
	}

	pushCMD := fmt.Sprintf("git push origin %s", v)

	out, err := cli.RunCommand(pushCMD)
	if err != nil {
		return "", err
	}

	return out, nil
}
