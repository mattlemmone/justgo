package desktop

import (
  "fmt"
  "os"
  "github.com/mitchellh/go-homedir"
  "strings"
  "path/filepath"
)

const (
  MaxFileWalkDepth = 7
)

type Linux struct { }

func (l *Linux) DesktopApplications() []*File {
  var apps []*File

  appDir := "/usr/share/applications/"
  filePattern := "*.desktop"
  glob := fmt.Sprintf("%s%s", appDir, filePattern)
  appPaths, _ := filepath.Glob(glob)

  for _, appPath := range appPaths {
    app := File{
      Type: ApplicationFileType,
      Path: appPath,
      Name: desktopApplicationNameFromPath(appPath),
    }    
    
    apps = append(apps, &app)
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
          // fmt.Printf("skipping: %s\n", fileInfo.Name())
          return filepath.SkipDir
        } 
      } else if strings.Count(path, "/") > MaxFileWalkDepth {
          // fmt.Printf("skipping long boy: %s\n", fileInfo.Name())
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

  // for _, appPath := range appPaths {
  //   app := File{
  //     Type: ApplicationFileType,
  //     Path: appPath,
  //     Name: desktopApplicationNameFromPath(appPath),
  //   }    
  //   fmt.Println(app)
  // }

  return files
}


// desktopApplicationNameFromPath converts a binary ".desktop" filepath to the binary's name
// e.g.: /usr/share/applications/python3.6.desktop -> python3.6 
func desktopApplicationNameFromPath(path string) string {
  slashIdx := strings.LastIndex(path, "/")
  appName := path[slashIdx + 1:]

  return strings.TrimSuffix(appName, ".desktop")
}

