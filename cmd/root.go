package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "miniCI",
	Short: "miniCI Go generate code",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err:%+v/n", err)
		os.Exit(1)
	}
}

// 初始化
func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("type", "t", "", "选择类型(web或cmd)")
}
