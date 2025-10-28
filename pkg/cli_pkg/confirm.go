package cli_pkg

import "github.com/manifoldco/promptui"

func (c *CliPkg) Confirm(label string) (bool, error) {
	items := []string{"はい", "いいえ"}
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return index == 0, nil
}
