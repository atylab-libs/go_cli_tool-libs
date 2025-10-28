package utils_pkg_mock

import (
	"os"

	"github.com/stretchr/testify/mock"
)

type OsMock struct {
	mock.Mock
}

func (o *OsMock) Exit(code int) {
	o.Called(code)
}

func (o *OsMock) Command(name string, args ...string) ([]byte, error) {
	callArgs := o.Called(append([]interface{}{name}, stringSliceToIface(args)...)...)
	return callArgs.Get(0).([]byte), callArgs.Error(1)
}

func (o *OsMock) CommandRun(name string, args ...string) error {
	callArgs := o.Called(append([]interface{}{name}, stringSliceToIface(args)...)...)
	return callArgs.Error(0)
}

func stringSliceToIface(s []string) []interface{} {
	out := make([]interface{}, len(s))
	for i, v := range s {
		out[i] = v
	}
	return out
}

func (o *OsMock) UserHomeDir() (string, error) {
	callArgs := o.Called()
	return callArgs.String(0), callArgs.Error(1)
}

func (o *OsMock) MkdirAll(path string, perm os.FileMode) error {
	callArgs := o.Called(path, perm)
	return callArgs.Error(0)
}

func (o *OsMock) WriteFile(filename string, data []byte, perm os.FileMode) error {
	callArgs := o.Called(filename, data, perm)
	return callArgs.Error(0)
}

func (o *OsMock) Stat(name string) (os.FileInfo, error) {
	callArgs := o.Called(name)
	return callArgs.Get(0).(os.FileInfo), callArgs.Error(1)
}

func (o *OsMock) IsNotExist(err error) bool {
	callArgs := o.Called(err)
	return callArgs.Bool(0)
}

func (o *OsMock) ReadFile(filename string) ([]byte, error) {
	callArgs := o.Called(filename)
	return callArgs.Get(0).([]byte), callArgs.Error(1)
}
