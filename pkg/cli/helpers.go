package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func Slugify(str string) string {
	str = strings.ToLower(str)
	// @TODO - finish this
	return str
}
