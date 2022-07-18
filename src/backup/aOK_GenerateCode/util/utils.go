package util

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func C_WFile(content, filename string) {
	// workingfolderDestination := readFolder(filename)
	// for _, wfd := range workingfolderDestination {
	// 	if strings.Contains(wfd.Name(), filename) {
	// fmt.Printf("wfd.Name(): %v\n", wfd.Name())
	f, createErr := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if createErr != nil {
		log.Printf("cannot create file:%v", createErr)
	}
	defer f.Close()

	_, writeErr := f.WriteString(content)
	if writeErr != nil {
		log.Printf("cannot write to file:%v", writeErr)
		// logger.Error("cannot write to file:", writeErr)
	}
	// 	} else {
	// 		log.Printf("no directory named:%v", wfd.Name())
	// 	}
	// }

}

func EnsureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
	// if cherr := os.Chmod(dirName, 0700); cherr != nil {
	// 	log.Printf("cannot chmod to file:%v", cherr)
	// }
}
func EnsureDir2(fileName string) error {
	dirName := filepath.Dir(fileName)
	// if _, serr := os.Stat(dirName); serr != nil {
	merr := os.MkdirAll(dirName, os.ModePerm)
	if merr != nil || os.IsExist(merr) {
		return merr
	}
	// }
	if cherr := os.Chmod(dirName, 0700); cherr != nil {
		log.Printf("cannot chmod to file:%v", cherr)
	}
	return nil
}
func ReadFolder(filename string) []fs.FileInfo {
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		log.Printf("no such file or directory:%v", err)
		// log.Fatal(err)
	}
	return files
}
