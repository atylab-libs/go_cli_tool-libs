package cli_pkg_mock

import (
	"cli_tool/common/pkg/cli_pkg"
)

func (m *CliPkgMock) SelectFromMap(label string, items []cli_pkg.MenuItem) string {
	args := m.Called(label, items)
	return args.String(0)
}
