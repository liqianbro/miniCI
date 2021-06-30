package cmd

import (
	"fmt"
	"testing"
)

func TestCreateFile(t *testing.T) {
	strs := []string{"apis", "cmd"}
	_, err := initializeProject(strs)
	if err != nil {
		fmt.Println(err)
	}

}
