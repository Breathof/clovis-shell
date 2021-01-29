package prompt

import (
	"os"
	"os/user"
	"strings"

	"github.com/clovis-shell/common"
)

//GetPrompt returns the user prompt
func GetPrompt() string {
	var prompt string
	user, err := user.Current()
	common.HandleError(err)
	dir, err := os.Getwd()

	pathSlice := strings.Split(dir, string(os.PathSeparator))
	path := strings.Join(pathSlice[len(pathSlice)-2:], string(os.PathSeparator))
	prompt = "(" + user.Username + ") " + path + "$ "

	return prompt
}

//PrintPrompt print the user prompt
func PrintPrompt() {
	prompt := GetPrompt()

	print(prompt)
}
