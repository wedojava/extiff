package extiff

import (
	"fmt"
	"testing"
)

func ExampleExtract() {
	tt := Tiff{Name: "test.tif", FilePath: "./example/test/test.tif"}
	tt.Extract()
	fmt.Printf("w-e: %f, n-s: %f\nMinX\t%f\nMinY\t%f\nMaxX\t%f\nMaxY\t%f\n",
		tt.WE, tt.NS, tt.MinX, tt.MinY, tt.MaxX, tt.MaxY)
	// Output:
	// w-e: 1.194620, n-s: -1.194361
	// MinX	11559833.372064
	// MinY	4312093.976939
	// MaxX	11563312.106151
	// MaxY	4314964.027273
}

func TestContains(t *testing.T) {
	tt := Tiff{Name: "test.tif", FilePath: "./example/test/test.tif"}
	tt.Extract()
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

func TestSetArea(t *testing.T) {
	as, err := ReadArea("./example/config.txt")
	if err != nil {
		t.Errorf("ReadArea err: %v", err)
	}
	tif := Tiff{FilePath: "./example/test/test.tif"}
	err = tif.Extract()
	if err != nil {
		t.Errorf("Extract err: %v", err)
	}
	tif.SetArea(as)
	i := 0
	for _, ta := range tif.Areas {
		if ta.Name == "Shot1" {
			i++
		}
		if ta.Name == "Shot2" {
			i++
		}
	}
	if i != 2 {
		t.Errorf("tif areas not match all: %v", tif.Areas)
	}
}
