package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/clovis-shell/commands"
	"github.com/clovis-shell/common"
	"github.com/clovis-shell/prompt"
)

func main() {
	inputReader()

}

func inputReader() {
	reader := bufio.NewReader(os.Stdin)
	for {
		prompt.PrintPrompt()
		input, err := reader.ReadString('\n')
		common.HandleError(err)

		command := strings.TrimSuffix(input, "\n")
		execCommand(strings.TrimSpace(command))
	}

}

func execCommand(commandString string) {

	switch commandString {
	case "ls":
		if len(os.Args) < 2 {
			commands.Ls("./")
		} else {
			commands.Ls(os.Args[1])
		}
	case "clear":
		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		c.Run()
	case "exit":
		os.Exit(0)
	}
}
