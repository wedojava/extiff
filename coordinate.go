package main

import "github.com/lukeroth/gdal"

type Coordinate struct {
	TopLeftX float64
	TopLeftY float64
	XSize    float64
	YSize    float64
	WE       float64
	NS       float64
	Points   [8]float64 // order by topleft topright bottomleft bottomright
}

func (c *Coordinate) ExtractCoordinates(filename string) error {
	dataset, err := gdal.Open(filename, gdal.ReadOnly)
	if err != nil {
		return err
	}
	// Reference: https://gdal.org/tutorials/geotransforms_tut.html?highlight=coordinate
	gt := dataset.GeoTransform()
	c.TopLeftX = gt[0]
	c.TopLeftY = gt[3]
	c.XSize = float64(dataset.RasterXSize()) // num of columns
	c.YSize = float64(dataset.RasterYSize()) // num of rows
	c.WE = gt[1]
	c.NS = gt[5]
	// t: top, b: bottom, l: left, r: right
	c.Points = [8]float64{
		c.TopLeftX, c.TopLeftY, // top left x and y
		c.TopLeftX + c.XSize*c.WE, c.TopLeftY, // top right x and y
		c.TopLeftX, c.TopLeftY + c.YSize*c.NS, // bottom left x and y
		c.TopLeftX + c.XSize*c.WE, c.TopLeftY + c.YSize*c.NS, // bottom right x and y
	}
	return nil
}
