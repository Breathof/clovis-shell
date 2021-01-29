package common

import (
	"fmt"
	"os"
)

//HandleError prints the error err in std.error
func HandleError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
