package extiff

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/wedojava/gdal"
)

type Area struct {
	Name string
	Env  gdal.Envelope
}

// readArea config file and return slice of Area
func ReadArea(cfgPath string) (as []Area, err error) {
	if cfgPath == "" {
		cfgPath, err = os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		cfgPath = path.Join(cfgPath, "config.txt")
	}

	f, err := os.OpenFile(cfgPath, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		area, err := strconvArea(s.Text())
		if err != nil {
			return nil, err
		}
		as = append(as, area)
	}

	return
}

// strconvArea convert a string to Area object, the str looks like `testArea1: 123, 719, 131, 819`
func strconvArea(str string) (Area, error) {
	var area Area
	a := strings.Split(str, ":")
	if a[0] == "" {
		area.Name = "unknowArea"
	}
	area.Name = a[0]
	env := strings.Split(a[1], ",")
	s0, s1, s2, s3 := strings.TrimSpace(env[0]),
		strings.TrimSpace(env[1]),
		strings.TrimSpace(env[2]),
		strings.TrimSpace(env[3])
	if s0 == "" {
		s0 = "0"
	}
	if s1 == "" {
		s1 = "0"
	}
	if s2 == "" {
		s2 = "0"
	}
	if s3 == "" {
		s3 = "0"
	}
	minx, err := strconv.ParseFloat(s0, 64)
	miny, err := strconv.ParseFloat(s1, 64)
	maxx, err := strconv.ParseFloat(s2, 64)
	maxy, err := strconv.ParseFloat(s3, 64)
	if err != nil {
		return Area{}, err
	}
	area.Env.SetMinX(minx)
	area.Env.SetMinY(miny)
	area.Env.SetMaxX(maxx)
	area.Env.SetMaxY(maxy)
	return area, nil
}
