package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hookRootCmd = &cobra.Command{
	Use:   "hookroot",
	Short: "hookroot 子命令短描述, 命令 hook",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hookroot run...")
	},
	// PersistentPreRunE: func(cmd *cobra.Command, args []string) error {},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Run 函数之前执行...
		fmt.Println("Run 函数之前执行...PersistentPreRun, 可以被子命令继承")
	},
	// PersistentPostRunE: func(cmd *cobra.Command, args []string) error {},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Run 函数之后执行...
		fmt.Println("Run 函数之后执行...PersistentPostRun, 可以被子命令继承")
	},
	// PreRunE: func(cmd *cobra.Command, args []string) error {},
	PreRun: func(cmd *cobra.Command, args []string) {
		// Run 函数之前执行...
		fmt.Println("Run 函数之前执行...PreRun")
	},
	// PostRunE: func(cmd *cobra.Command, args []string) error {},
	PostRun: func(cmd *cobra.Command, args []string) {
		// Run 函数之后执行...
		fmt.Println("Run 函数之后执行...PostRun")
	},
}

var hookSubCmd = &cobra.Command{
	Use:   "hooksub",
	Short: "hooksub",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hooksub run...")
	},
}

func init() {
	rootCmd.AddCommand(hookRootCmd)

	hookRootCmd.AddCommand(hookSubCmd)
}
