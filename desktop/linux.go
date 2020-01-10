package desktop

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	maxFileWalkDepth = 7
)

type Linux struct{}

func (l *Linux) LaunchApplication(s string) {
	// replace all occurences of %U (file input) so app launches with no params, and
	// instead, with the -d(etached) flag
	sanitized := strings.Replace(s, "%U", "-d", -1)
	args := append([]string{"bash",  "-c"}, sanitized)
	cmd := exec.Command("nohup", args...)
	// out, err := cmd.Output()

  fmt.Printf("> %s\n", sanitized)
	// println(string(out))

	cmd.Start()
}

func (l *Linux) DesktopApplications() []*Application {
	var apps []*Application

	appDirs := []string{
		"/usr/share/applications/",
		"/var/lib/snapd/desktop/applications/",
	}

	for _, appDir := range appDirs {
		filePattern := "*.desktop"
		glob := fmt.Sprintf("%s%s", appDir, filePattern)
		appPaths, _ := filepath.Glob(glob)

		for _, appPath := range appPaths {
			app := applicationFromDesktopFile(appPath)
			apps = append(apps, app)
		}
	}

	return apps
}

func (l *Linux) UserFiles() []*File {
	var files []*File

	appDir, _ := homedir.Dir()

	filepath.Walk(
		appDir,
		func(path string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			isDotFile := strings.HasPrefix(fileInfo.Name(), ".")

			if isDotFile {
				if fileInfo.IsDir() {
					return filepath.SkipDir
				}
			} else if strings.Count(path, "/") > maxFileWalkDepth {
				return filepath.SkipDir
			} else {
				file := File{
					Type: DefaultFileType,
					Path: path,
					Name: fileInfo.Name(),
				}

				files = append(files, &file)
			}

			return nil
		},
	)

	return files
}


func applicationFromDesktopFile(path string) *Application {
	params := map[string]string{}

	fileLines, _ := readFileLines(path)
  foundDesktopEntry := false

	for _, line := range fileLines {
    if !foundDesktopEntry && strings.Contains(line, "[Desktop Entry]") {
      foundDesktopEntry = true
      continue
    }

    // found non-default desktop entry
    if foundDesktopEntry && strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
      break
    }

		sepIdx := strings.Index(line, "=")

		if sepIdx == -1 {
			continue
		}

		key := line[:sepIdx]
		val := line[sepIdx+1:]
		params[key] = val
	}

	fmt.Printf("app name: %s\n", params["Name"])

	return &Application{
		Name:     params["Name"],
    Path:     path,
		IconPath: params["Icon"],
		Exec:     params["Exec"],
	}
}
