package updater

import (
	"testing"

	"github.com/cedricmar/updep/pkg/tagger"
)

func TestTest(t *testing.T) {

	infoBool := true

	tg, _ := tagger.NewTag("logger@1.0")

	f := Flags{
		Info: &infoBool,
		Tag:  tg,
	}

	Update(f)
}
