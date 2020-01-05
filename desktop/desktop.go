package desktop

type FileType int

const (
  DefaultFileType FileType = iota
  ApplicationFileType
)

type File struct {
  Type FileType
  Name string
  Path string
}

type OperatingSystem interface {
  DesktopApplications() []File
  UserFiles() []File
}
