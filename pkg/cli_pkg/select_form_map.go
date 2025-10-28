package cli_pkg

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

type MenuItem struct {
	Key   string
	Label string
}

func (c *CliPkg) SelectFromMap(label string, items []MenuItem) string {
	maxKeyLen := 0
	for _, item := range items {
		if len(item.Key) > maxKeyLen {
			maxKeyLen = len(item.Key)
		}
	}

	keys := make([]string, len(items))
	labels := make([]string, len(items))

	for i, item := range items {
		padding := strings.Repeat(" ", maxKeyLen-len(item.Key))
		keys[i] = item.Key
		labels[i] = fmt.Sprintf("%s%s : %s", item.Key, padding, item.Label)
	}

	prompt := promptui.Select{
		Label: label,
		Items: labels,
	}

	index, _, err := prompt.Run()
	if err != nil {
		if err == promptui.ErrInterrupt {
			fmt.Println("🛑 GUI選択が中断されました。")
			os.Exit(0)
		}
		fmt.Println("選択に失敗しました:", err)
		os.Exit(1)
	}

	return keys[index]
}
