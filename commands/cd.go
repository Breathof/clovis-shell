package commands

import (
	"os"
	"os/user"

	"github.com/clovis-shell/common"
)

//Cd change directory to que path
func Cd(input []string) {
	usr, err := user.Current()
	common.HandleError(err)

	if len(input) > 1 {
		err = os.Chdir(input[1])
	} else {
		err = os.Chdir(usr.HomeDir)
	}

	common.HandleError(err)
}
