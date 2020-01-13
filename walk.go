package main

import (
	"os"
	"fmt"
	"path/filepath"
	"strings"
)


const maxFileWalkDepth = 7


func discoverPaths() []string {
	var paths []string

  cwd, err := os.Getwd()

  if err != nil {
    fmt.Println(err.Error())
  }

	filepath.Walk(
		cwd,
		func(path string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

      if !fileInfo.IsDir() {
        return nil
      }

			isDotFile := strings.HasPrefix(fileInfo.Name(), ".")

			if isDotFile {
        return filepath.SkipDir
			} else if strings.Count(path, "/") > maxFileWalkDepth {
				return filepath.SkipDir
			} else {
				paths = append(paths, path)
			}

			return nil
		},
	)

	return paths
}

