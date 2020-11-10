package extiff

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	// "github.com/lukeroth/gdal"
	gdal "github.com/tingold/gdal-1"
)

type Tiff struct {
	MinX, MinY, MaxX, MaxY float64
	WE, NS                 float64
	Env                    gdal.Envelope
	Name, FilePath, Dir    string
	Areas                  []Area
	Siblings               []string
}

type Coordinate struct {
	X float64
	Y float64
}

// GetTifs walk the target folder (default is `.`), catch files have suffix `.tif`
func GetTifs(dir string) (ts []*Tiff, err error) {
	if dir == "" {
		dir = "./"
	}
	_, err = os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		return
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".tif") {
			ts = append(ts,
				&Tiff{Name: info.Name(),
					FilePath: path,
					Dir:      filepath.Dir(path)})
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walking the path %s error : %v\n", dir, err)
	}
	return
}

func (t *Tiff) Extract() error {
	dataset, err := gdal.Open(t.FilePath, gdal.ReadOnly)
	if err != nil {
		return err
	}
	defer dataset.Close()
	// Reference: https://gdal.org/tutorials/geotransforms_tut.html?highlight=coordinate
	gt := dataset.GeoTransform()
	xSize := float64(dataset.RasterXSize()) // num of columns
	ySize := float64(dataset.RasterYSize()) // num of rows
	t.WE, t.NS = gt[1], gt[5]
	t.MinX, t.MaxX = gt[0], gt[0]+xSize*t.WE
	t.MinY, t.MaxY = gt[3], gt[3]+ySize*t.NS
	if t.MinX > t.MaxX {
		t.MinX, t.MaxX = t.MaxX, t.MinX
	}
	if t.MinY > t.MaxY {
		t.MinY, t.MaxY = t.MaxY, t.MinY
	}
	t.Env.SetMinX(t.MinX)
	t.Env.SetMinY(t.MinY)
	t.Env.SetMaxX(t.MaxX)
	t.Env.SetMaxY(t.MaxY)
	return nil
}

func (t *Tiff) Contains(c *Coordinate) bool {
	// env2 := gdal.Envelope{}
	// env2.SetMinX(c.X)
	// env2.SetMaxX(c.X)
	// env2.SetMinY(c.Y)
	// env2.SetMaxY(c.Y)
	// return t.Env.Contains(env2)
	return t.MinX <= c.X && t.MaxX >= c.X && t.MinY <= c.Y && t.MaxY >= c.Y
}

func (t *Tiff) SetArea(as []Area) {
	for _, a := range as {
		if a.Env.Intersects(t.Env) {
			t.Areas = append(t.Areas, a)
		}
	}
}

func (t *Tiff) GetSiblings() {
	ss := []string{}
	prename := strings.Split(t.Name, ".tif")[0]
	fs, _ := ioutil.ReadDir(t.Dir)
	for _, f := range fs {
		fname := f.Name()
		if strings.HasPrefix(fname, prename) {
			ss = append(ss, fname)
		}
	}
	t.Siblings = ss
}

func (t *Tiff) Rename() error {
	prefix := ""
	as := t.Areas
	for _, a := range as {
		prefix = "[" + a.Name + "]" + prefix
	}
	err := os.Rename(t.FilePath, filepath.Join(t.Dir, prefix+t.Name))
	if err != nil {
		return err
	}
	t.GetSiblings()
	for _, s := range t.Siblings {
		err = os.Rename(filepath.Join(t.Dir, s), filepath.Join(t.Dir, prefix+s))
		if err != nil {
			return err
		}
	}
	prename := strings.Split(t.Name, ".tif")[0]
	if strings.Contains(t.Dir, prename) {
		os.Rename(t.Dir, strings.ReplaceAll(t.Dir, prename, prefix+prename))
	}
	return nil
}
