package fileutil

import (
	"fmt"
	"io/ioutil"
)

var listFilePrefix string = "  "

func listAllFileByName(level int, pathSeparator string, fileDir string) {
	files, _ := ioutil.ReadDir(fileDir)

	tmpPrefix := ""
	for i := 1; i < level; i++ {
		tmpPrefix = tmpPrefix + listFilePrefix
	}

	for _, onefile := range files {
		if onefile.IsDir() {
			fmt.Printf("\033[34m %s %s \033[0m \n", tmpPrefix, onefile.Name())
			//fmt.Println(tmpPrefix, onefile.Name(), "目录:")
			listAllFileByName(level+1, pathSeparator, fileDir+pathSeparator+onefile.Name())
		} else {
			fmt.Println(tmpPrefix, onefile.Name())
		}
	}

}

//func main() {
//	srcDir := "C:\\Users\\Administrator\\Desktop\\logs"
//	pathSeparator := string(os.PathSeparator)
//	level := 1
//	listAllFileByName(level, pathSeparator, srcDir)
//}
