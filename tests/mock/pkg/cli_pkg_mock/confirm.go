package cli_pkg_mock

func (c *CliPkgMock) Confirm(label string) (bool, error) {
	args := c.Called(label)
	return args.Bool(0), args.Error(1)
}
