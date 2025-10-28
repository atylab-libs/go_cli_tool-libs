package cli_pkg_mock

func (m *CliPkgMock) Read(prompt string) (string, error) {
	args := m.Called(prompt)
	return args.String(0), args.Error(1)
}
