package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/clovis-shell/common"
)

//Ls function list the contents of a direcory
func Ls(input string) {
	var files []os.FileInfo
	var err error
	if input != "" {
		files, err = ioutil.ReadDir(input)

	} else {
		files, err = ioutil.ReadDir("./")
	}

	common.HandleError(err)

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
