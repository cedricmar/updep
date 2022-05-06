package updater

import (
	"testing"

	"github.com/cedricmar/updep/pkg/tag"
)

func TestTest(t *testing.T) {

	infoBool := true

	tg, _ := tag.NewTag("logger@1.0")

	f := Flags{
		Info: &infoBool,
		Tag:  tg,
	}

	Update(f)
}
