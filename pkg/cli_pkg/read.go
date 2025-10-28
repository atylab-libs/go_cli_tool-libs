package cli_pkg

import (
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
)

func (c *CliPkg) Read(prompt string) (string, error) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          prompt,
		InterruptPrompt: "^C で中断されました。\n",
		EOFPrompt:       "終了します。\n",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	line, err := rl.Readline()
	if err == readline.ErrInterrupt {
		return "", io.EOF
	}
	if err != nil {
		return "", fmt.Errorf("入力中にエラーが発生しました: %v", err)
	}

	line = strings.TrimSpace(line)
	if line == "exit" {
		return "", io.EOF
	}
	return line, nil
}
