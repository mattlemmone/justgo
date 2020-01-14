package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

const maxFileWalkDepth = 6

var blacklistedDirs = map[string]bool{
    "build": true,
    "node_modules": true,
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

            if shouldSkipDirectory(fileInfo, path) {
                return filepath.SkipDir
            }

            paths = append(paths, path)

            return nil
        },
    )

    fmt.Println(paths)

    return paths
}

func shouldSkipDirectory(fileInfo os.FileInfo, path string) bool {
    isDotFile := strings.HasPrefix(fileInfo.Name(), ".")
    exceedsMaxFolderDepth := strings.Count(path, "/") > maxFileWalkDepth

    if isDotFile || exceedsMaxFolderDepth || isBlacklisted(path) {
        return true
    }

    return false
}

func isBlacklisted(path string) bool {
    fileNameIdx := strings.LastIndex( path,"/")

    if fileNameIdx == -1 {
        return false
    }

    fileName := path[fileNameIdx+1:]

    if _, ok := blacklistedDirs[fileName]; ok {
        return true
    }

    return false
}
