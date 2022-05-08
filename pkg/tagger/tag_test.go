package tagger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLastNum(t *testing.T) {
	tests := []struct {
		tag  string
		want string
	}{
		{tag: "1", want: "2"},
		{tag: "v1", want: "v2"},
		{tag: "1.0.1", want: "1.0.2"},
		{tag: "1.0.19", want: "1.0.20"},
		{tag: "v1.0.1-123.alpha.1", want: "v1.0.1-123.alpha.2"},
		{tag: "1.0.199*£.alpha.---", want: "1.0.200*£.alpha.---"},
		{tag: "v2.12.0-alpha.1", want: "v2.12.0-alpha.2"},
		{tag: "v2.12.0-alpha-9", want: "v2.12.0-alpha-10"},
		{tag: "1.alpha.acb", want: "2.alpha.acb"},
		{tag: "abc", want: ""},
		{tag: "99", want: "100"},
		{tag: "99v", want: "100v"},
		{tag: "v99", want: "v100"},
		{tag: "v99v", want: "v100v"},
		{tag: "-1", want: "-2"},
	}
	for _, tt := range tests {
		t.Run("Bump "+tt.tag, func(t *testing.T) {
			got, _ := BumpVersion(tt.tag)
			assert.Equal(t, tt.want, got, fmt.Sprintf("BumpVersion() = %v, want %v", got, tt.want))
		})
	}
}
