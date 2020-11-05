package extiff

import (
	"testing"

	"github.com/lukeroth/gdal"
)

func TestCollect(t *testing.T) {
	a := Area{Name: "area1", Env: gdal.Envelope{}}
	a.Collect("")
	if false {
		t.Error("")
	}
}
