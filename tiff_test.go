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
	// Upper left	[11559833.372064, 4314964.027273]
	// Upper right	[11563312.106151, 4314964.027273]
	// Bottom left	[11559833.372064, 4312093.976939]
	// Bottom right	[11563312.106151, 4312093.976939]
}

func TestContains(t *testing.T) {
	filename := "./example/test/L18/test.tif"
	tt := Tiff{}
	tt.Extract(filename)
	tcs := []struct {
		c *Coordinate
		b bool
	}{
		{&Coordinate{11559834.0, 4313963.0}, true},
		{&Coordinate{11559833.0, 4314965.0}, false}, // left
		{&Coordinate{11559834.0, 4314963.0}, false}, // up
		{&Coordinate{11563313.0, 4314965.0}, false}, // right
		{&Coordinate{11563312.0, 4312092.0}, false}, // bottom
	}

	for _, tc := range tcs {
		got := tt.Contains(tc.c)
		if got != tc.b {
			t.Errorf("want: %v, got: %v", tc.b, got)
		}
	}
}
