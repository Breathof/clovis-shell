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

	command := strings.Split(commandString, " ")
	args := command[1:]
	switch command[0] {
	case "ls":
		commands.Ls(command)
	case "clear":
		commands.Clear()
	case "cd":
		commands.Cd(command)
	case "exit":
		os.Exit(0)
	case "pwd":
		commands.Pwd()
	case "":
	default:
		cmd := exec.Command(command[0], args...)
		common.AssignStd(cmd)
		err := cmd.Run()
		common.HandleError(err)
	}

}
