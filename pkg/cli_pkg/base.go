package cli_pkg

type CliPkgInterface interface {
	SelectFromMap(label string, items []MenuItem) string
	Read(prompt string) (string, error)
	Confirm(label string) (bool, error)
}

type CliPkg struct{}

func NewCliPkg() *CliPkg {
	return &CliPkg{}
}
