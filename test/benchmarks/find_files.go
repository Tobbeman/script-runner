package benchmarks

import (
	"os"
	"path/filepath"
	"strings"
)

func walk(dirS string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirS, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			files = append(files, strings.ReplaceAll(path, dirS, "")[1:])
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func homeBaked(dirS string) ([]string, error) {
	dir, err := os.Open(dirS)
	if err != nil {
		return nil, err
	}
	files, err := recurseGetFilenames(dir)
	if err != nil {
		return nil, err
	}
	err = dir.Close()
	if err != nil {
		return nil, err
	}
	return files, nil
}

func recurseGetFilenames(dir *os.File) ([]string, error) {
	var files []string
	list, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for _, file := range list {
		if file.IsDir() {
			nestedDir, err := os.Open(dir.Name() + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			nestedFiles, err := recurseGetFilenames(nestedDir)
			if err != nil {
				return nil, err
			}
			err = nestedDir.Close()
			if err != nil {
				return nil, err
			}
			for _, nestedFile := range nestedFiles {
				files = append(files, file.Name()+"/"+nestedFile)
			}
		} else {
			files = append(files, file.Name())
		}
	}
	return files, nil
}
