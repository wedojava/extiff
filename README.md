# ExTiff

ExTiff is the tool to extract coordinates from tif file.

# For Developers

## References:

- [GDAL Geotransform Tutorial (with coordinate highlight)](https://gdal.org/tutorials/geotransforms_tut.html?highlight=coordinate)

## gdal prepare

1. `brew install gdal`, it'll eat long time.
2. `go get -u github.com/lukeroth/gdal`, maybe error occur this step.
3. fix error occur **could not determine kind of name for C.OPTGetParameterInfo** on the last step by [this line](https://github.com/lukeroth/gdal/issues/53#issuecomment-670553446), `go.mod` line should be:
```
replace github.com/lukeroth/gdal v0.0.0-20200603180320-4c5dda23594f => github.com/tingold/gdal-1 v0.0.0-20200805034744-092f31c3aae1
```
4. `go mod tidy`

## gdal tips

### A geotransform consists in a set of 6 coefficients

- GT(0) x-coordinate of the upper-left corner of the upper-left pixel.
- GT(1) w-e pixel resolution / pixel width.
- GT(2) row rotation (typically zero).
- GT(3) y-coordinate of the upper-left corner of the upper-left pixel.
- GT(4) column rotation (typically zero).
- GT(5) n-s pixel resolution / pixel height (negative value for a north-up image).

### Transformation from image coordinate space to georeferenced coordinate space

```
X_geo = GT(0) + X_pixel * GT(1) + Y_line * GT(2)
Y_geo = GT(3) + X_pixel * GT(4) + Y_line * GT(5)
```
Note that the pixel/line coordinates in the above are from (0.0,0.0) at the top left corner of the top left pixel to (width_in_pixels,height_in_pixels) at the bottom right corner of the bottom right pixel. The pixel/line location of the center of the top left pixel would therefore be (0.5,0.5).

### In case of north up images

- GT(2), GT(4) coefficients are zero.
- GT(1), GT(5) is the pixel size.
- GT(0), GT(3) position is the top left corner of the top left pixel of the raster.
