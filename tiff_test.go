package main

import (
	"fmt"
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
