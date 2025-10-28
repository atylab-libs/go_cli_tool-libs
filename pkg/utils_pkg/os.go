package utils_pkg

import (
	"os"
	"os/exec"
)

type OsInterface interface {
	Exit(code int)
	Command(name string, args ...string) ([]byte, error)
	CommandRun(name string, args ...string) error
	UserHomeDir() (string, error)
	MkdirAll(path string, perm os.FileMode) error
	WriteFile(filename string, data []byte, perm os.FileMode) error
	Stat(name string) (os.FileInfo, error)
	IsNotExist(err error) bool
	ReadFile(filename string) ([]byte, error)
}

func NewOS() OsInterface {
	return &defaultOS{}
}

type defaultOS struct{}

func (d *defaultOS) Exit(code int) {
	os.Exit(code)
}

func (d *defaultOS) Command(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *defaultOS) CommandRun(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (d *defaultOS) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (d *defaultOS) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (d *defaultOS) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

func (d *defaultOS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (d *defaultOS) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (d *defaultOS) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
