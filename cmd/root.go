package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Qian",
	Short: "fast mini_ci Go code",
	Long:  "fast mini_ci Go code-happy",
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
}
