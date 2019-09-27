package main

import (
	"path/filepath"
	"os"
	"../../util"
	"strings"
)

var dirpath = "../mzitu"

func main() {
	//重命名不规范的图片名
	filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				filepath.Walk(path,
					func(path2 string, f2 os.FileInfo, err2 error) error {
						if f2 == nil {
							return err2
						}
						if !(f2.IsDir()) {
							if len(f2.Name()) > 6 {
								index := strings.LastIndex(path2, "\\")
								util.Rename(path2, path2[0:index+1]+ path2[len(path2)-6:len(path2)])
							}
						}
						return nil
					})
			}
			return nil
		})
}
