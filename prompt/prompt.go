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
	path := pathSlice[len(pathSlice)-2]
	prompt = "(" + user.Username + ") " + path + "$ "

	return prompt
}

//PrintPrompt print the user prompt
func PrintPrompt() {
	prompt := GetPrompt()

	print(prompt)
}
