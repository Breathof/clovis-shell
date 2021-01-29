package commands

import (
	"fmt"
	"os"

	"github.com/clovis-shell/common"
)

//Pwd print working directory
func Pwd() {
	path, err := os.Getwd()
	common.HandleError(err)
	fmt.Println(path)
	println()
}
