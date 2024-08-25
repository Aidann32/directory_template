package os_utils

type OSUtils interface {
	ValidateFileName(name string) (err error)
	ValidateDirectoryName(name string) (err error)
}
