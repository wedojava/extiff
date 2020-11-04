package main

import (
	"fmt"
)

func ExampleExtractCoordinates() {
	filename := "./example/test/L18/test.tif"
	c := Coordinate{}
	c.ExtractCoordinates(filename)
	fmt.Printf("w-e: %f, n-s: %f\nUpper left\t[%f, %f]\nUpper right\t[%f, %f]\nBottom left\t[%f, %f]\nBottom right\t[%f, %f]\n",
		c.WE, c.NS,
		c.Points[0], c.Points[1],
		c.Points[2], c.Points[3],
		c.Points[4], c.Points[5],
		c.Points[6], c.Points[7])
	// Output:
	// w-e: 1.194620, n-s: -1.194361
	// Upper left	[11559833.372064, 4314964.027273]
	// Upper right	[11563312.106151, 4314964.027273]
	// Bottom left	[11559833.372064, 4312093.976939]
	// Bottom right	[11563312.106151, 4312093.976939]
}
