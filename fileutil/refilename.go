package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReNameFile() {
	// 解析命令行参数
	//var dir string
	//flag.StringVar(&dir, "d", "", "directory path")
	//var pattern string
	//flag.StringVar(&pattern, "p", "", "name pattern, eg. newname%d")
	//flag.Parse()
	//if dir == "" || pattern == "" {
	//	flag.Usage()
	//	return
	//}

	dir := "C:\\Users\\Administrator\\Desktop\\logs"
	//pattern := "logz"

	// 遍历文件夹，获取文件路径
	paths := make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	// 遍历文件路径，修改文件名
	for i, path := range paths {
		//newPath := filepath.Join(filepath.Dir(path), filepath.Ext(path))
		//newPath := filepath.Join(filepath.Dir(path), filepath.VolumeName(path), ".zip")
		newPath := filepath.Join(filepath.Dir(path), strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))+".zip")
		os.Rename(path, newPath)
		fmt.Println(i, newPath)
	}
}
