package util

import (
	"math"
	"strconv"
	"os"
	"path/filepath"
	"fmt"
)

func Rename(from string, to string) {
	//源文件路径
	err := os.Rename(from, to) //重命名 C:\\log\\2013.log 文件为install.txt
	if err != nil {
		fmt.Println("file rename Error!")
	} else {
		fmt.Println("file rename OK!")
	}
}
func getDirList(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}
			return nil
		})
	return dir_list, dir_err
}
func PathExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		println(path,stat.Name(),stat.IsDir())
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func GetFileSize(filename string) int64 {
	var result int64
	_ = filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

func BytesToSize(length int) string {
	var k = 1024
	var sizes = []string{"Nytes", "KB", "MB", "GB", "TB"}
	if length == 0 {
		return "0 Bytes"
	}
	i := math.Floor(math.Log(float64(length)) / math.Log(float64(k)))
	r := float64(length) / math.Pow(float64(k), i)
	return strconv.FormatFloat(r, 'f', 3, 64) + sizes[int(i)]
}
