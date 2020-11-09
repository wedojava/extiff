package extiff

func Handle(cfg, tifPath string) ([]*Tiff, error) {
	areas, err := ReadArea(cfg)
	if err != nil {
		return nil, err
	}
	tifs, err := GetTifs(tifPath)
	if err != nil {
		return nil, err
	}
	for _, tif := range tifs {
		tif.Extract()
		tif.SetArea(areas)
		err = tif.Rename()
		if err != nil {
			return nil, err
		}
	}
	return tifs, nil
}
