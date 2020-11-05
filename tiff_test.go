package main

import (
	"fmt"
	"testing"
)

func ExampleExtract() {
	filename := "./example/test/L18/test.tif"
	tt := Tiff{}
	tt.Extract(filename)
	fmt.Printf("w-e: %f, n-s: %f\nUpper left\t[%f, %f]\nUpper right\t[%f, %f]\nBottom left\t[%f, %f]\nBottom right\t[%f, %f]\n",
		tt.WE, tt.NS,
		tt.Points[0].X, tt.Points[0].Y,
		tt.Points[1].X, tt.Points[1].Y,
		tt.Points[2].X, tt.Points[2].Y,
		tt.Points[3].X, tt.Points[3].Y)
	// Output:
	// w-e: 1.194620, n-s: -1.194361
	// Upper left	[11559833.372064, 4312093.976939]
	// Upper right	[11563312.106151, 4312093.976939]
	// Bottom left	[11559833.372064, 4314964.027273]
	// Bottom right	[11563312.106151, 4314964.027273]
}

func TestContains(t *testing.T) {
	filename := "./example/test/L18/test.tif"
	tt := Tiff{}
	tt.Extract(filename)
	tcs := []struct {
		c *Coordinate
		b bool
	}{
		{&Coordinate{11559934.372064, 4313963.027273}, true}, // point: in
		{&Coordinate{11559832.0, 4314965.0}, false},          // point: left
		{&Coordinate{11559835.0, 4311962.0}, false},          // point: up
		{&Coordinate{11563313.0, 4314965.0}, false},          // point: right
		{&Coordinate{11563312.0, 4315092.0}, false},          // point: bottom
	}

	for _, tc := range tcs {
		got := tt.Contains(tc.c)
		if got != tc.b {
			fmt.Printf("tt.MinX is: \t%f\nc.c.X is: \t%f\n", tt.MinX, tc.c.X)
			fmt.Printf("tt.MinY is: \t%f\ntc.c.Y is: \t%f\n", tt.MinY, tc.c.Y)
			fmt.Printf("tt.MaxX is: \t%f\ntc.c.X is: \t%f\n", tt.MaxX, tc.c.X)
			fmt.Printf("tt.MaxY is: \t%f\ntc.c.Y is: \t%f\n", tt.MaxY, tc.c.Y)
			t.Error("error occur, check the output info.")
		}
	}
}
