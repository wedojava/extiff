package extiff

import (
	"log"
	"os"
	"path"

	"github.com/lukeroth/gdal"
)

type Area struct {
	Name string
	Env  gdal.Envelope
}

// Collect will
func (a *Area) Collect(cfg string) (as []Area, err error) {
	if cfg == "" {
		cfg, err = os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		cfg = path.Join(cfg, "config.txt")
	}
	return
}
