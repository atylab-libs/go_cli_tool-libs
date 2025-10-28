package cli_pkg_mock

import "github.com/atylab-libs/go_cli_tool-libs/pkg/cli_pkg"

func (m *CliPkgMock) SelectFromMap(label string, items []cli_pkg.MenuItem) string {
	args := m.Called(label, items)
	return args.String(0)
}
