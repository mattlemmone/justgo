package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

const maxFileWalkDepth = 7
const blacklistedDirs = []map[string]bool{
    "build": true,
}

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

            if shouldSkipDirectory(path) {
                return filepath.SkipDir
            }

            paths = append(paths, path)

            return nil
        },
    )

    return paths
}

func shouldSkipDirectory(path string) bool {
    isDotFile := strings.HasPrefix(fileInfo.Name(), ".")

    if isDotFile {
        return true
    }

    exceedsMaxFolderDepth =  strings.Count(path, "/") > maxFileWalkDepth

    if exceedsMaxFolderDepth {
        return true
    }

    fileNameIdx := strings.LastIndex("/", path)

    if fileNameIdx != -1 {
        fileName := path[fileNameIdx+1:]

        if _, ok := blacklistedDirs[fileName]; ok {
            return true
        }
    }

    return false    
}
