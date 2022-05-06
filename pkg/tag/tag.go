package tag

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Tag struct {
	Key     string
	Name    string
	Version string
}

// NewTag parses the cli provided tag key
// @TODO - write a test for this
func NewTag(key string) (Tag, error) {
	if !strings.Contains(key, "@") {
		return Tag{}, errors.New("malformed dependency tag, use <repo_name>@<version_tag>")
	}
	t := strings.Split(key, "@")
	return Tag{
		Key:     key,
		Name:    t[0],
		Version: t[1],
	}, nil
}

func BumpVersion(tag string) (string, error) {
	// Search for the last number in the string and increment it
	idx, oldLen, val := findLastNum(tag)
	if idx < 0 {
		return "", errors.New("tag cannot be bumped")
	}
	nextSegIdx := idx + oldLen
	return tag[:idx] + val + tag[nextSegIdx:], nil
}

// if idx = -1 then it failed
func findLastNum(tag string) (idx, length int, val string) {
	idx = len(tag) - 1
	for ; idx >= 0; idx-- {
		if unicode.IsDigit(rune(tag[idx])) {
			// Stack digits
			val = fmt.Sprintf("%s%s", string(tag[idx]), val)
			continue
		}
		// We are done consuming digits
		if len(val) > 0 {
			break
		}
	}

	d, err := strconv.Atoi(val)
	if err != nil {
		return -1, 0, ""
	}

	idx++
	d++

	return idx, len(val), strconv.Itoa(d)
}
