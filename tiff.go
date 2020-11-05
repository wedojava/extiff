package main

import (
	"github.com/lukeroth/gdal"
)

type Tiff struct {
	MinX, MinY, MaxX, MaxY float64
	WE, NS                 float64
	Points                 []Coordinate // order by topleft topright bottomleft bottomright
	Env                    gdal.Envelope
}

type Coordinate struct {
	X float64
	Y float64
}

func (t *Tiff) Extract(filename string) error {
	dataset, err := gdal.Open(filename, gdal.ReadOnly)
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
	// t: top, b: bottom, l: left, r: right
	t.Points = []Coordinate{
		{t.MinX, t.MinY}, // top left x and y
		{t.MaxX, t.MinY}, // top right x and y
		{t.MinX, t.MaxY}, // bottom left x and y
		{t.MaxX, t.MaxY}, // bottom right x and y
	}
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

func (t1 *Tiff) Intersection(t2 *Tiff) bool {
	env1 := gdal.Envelope{}
	env1.SetMinX(t1.MinX)
	env1.SetMinY(t1.MinY)
	env1.SetMaxX(t1.MaxX)
	env1.SetMaxY(t1.MaxY)

	env2 := gdal.Envelope{}
	env2.SetMinX(t2.MinX)
	env2.SetMinY(t2.MinY)
	env2.SetMaxX(t2.MaxX)
	env2.SetMaxY(t2.MaxY)
	return env1.Intersects(env2)
}
