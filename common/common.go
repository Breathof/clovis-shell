package common

import (
	"fmt"
	"os"
	"os/exec"
)

//HandleError prints the error err in std.error
func HandleError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		println()
	}
}

//AssignStd assigns the Stdout and Stderr
func AssignStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
