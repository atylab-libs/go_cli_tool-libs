package cli_pkg_mock

import "github.com/stretchr/testify/mock"

type CliPkgMock struct {
	mock.Mock
}

func NewCliPkgMock() *CliPkgMock {
	return &CliPkgMock{}
}
