package main

import "github.com/lukeroth/gdal"

type Tiff struct {
	TopLeftX float64
	TopLeftY float64
	XSize    float64
	YSize    float64
	WE       float64
	NS       float64
	Points   []Coordinate // order by topleft topright bottomleft bottomright
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
	// Reference: https://gdal.org/tutorials/geotransforms_tut.html?highlight=coordinate
	gt := dataset.GeoTransform()
	t.TopLeftX = gt[0]
	t.TopLeftY = gt[3]
	t.XSize = float64(dataset.RasterXSize()) // num of columns
	t.YSize = float64(dataset.RasterYSize()) // num of rows
	t.WE = gt[1]
	t.NS = gt[5]
	// t: top, b: bottom, l: left, r: right
	t.Points = []Coordinate{
		{t.TopLeftX, t.TopLeftY},                               // top left x and y
		{t.TopLeftX + t.XSize*t.WE, t.TopLeftY},                // top right x and y
		{t.TopLeftX, t.TopLeftY + t.YSize*t.NS},                // bottom left x and y
		{t.TopLeftX + t.XSize*t.WE, t.TopLeftY + t.YSize*t.NS}, // bottom right x and y
	}
	return nil
}
