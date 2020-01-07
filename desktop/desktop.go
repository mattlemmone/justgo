package desktop

import (
	"bufio"
	"os"
)

type FileType int

const (
	DefaultFileType FileType = iota
	ApplicationFileType
)

type Application struct {
	Name     string
	IconPath string
	Path     string
	Exec     string
}

type File struct {
	Type FileType
	Name string
	Path string
}

type OperatingSystem interface {
	DesktopApplications() []File
	UserFiles() []File
}

func readFileLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
