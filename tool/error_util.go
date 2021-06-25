package tool

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// CheckErr prints the msg with the prefix 'Error:' and exits with error code 1. If the msg is nil, it does nothing.
func CheckErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Errorf("err: %+v", err), err)
		os.Exit(1)
	}
}
