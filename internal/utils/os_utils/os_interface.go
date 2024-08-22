package os_utils

type OSUtils interface {
	InitGoMod(moduleName string) (err error)
	ValidateFileName(name string) (result bool)
	ValidateDirectoryName(name string) (err error)
	DoesFileExist() (result bool)
	IsGoInstalled() (err error)
}
