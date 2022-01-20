package eject_test

import (
	"testing"

	"github.com/Triment/eject"
)

func TestCreateContext(t *testing.T) {
	ctx := eject.CreateContext()
	ctx.Params["hh"] = "k"
}
