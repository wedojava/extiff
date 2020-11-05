package extiff

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetTifs walk the target folder (default is `.`), catch files have suffix `.tif`
func GetTifs(dir string) (files []string, err error) {
	if dir == "" {
		dir = "./"
	}
	_, err = os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		return
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".tif") {
			fmt.Printf("deal with file or dir: %q\n", path)
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walking the path %s error : %v\n", dir, err)
	}
	return
}
