package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var checkArgsCmd = &cobra.Command{
	Use:   "check",
	Short: "检查参数短描述",
	Long:  "检查参数长描述",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少输入一个参数")
		}
		if len(args) > 2 {
			return errors.New("最多输入两个参数")
		}
		return nil
	},
	// Run 没有返回值
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check run...")
		fmt.Println("args", args)
		fmt.Println(args)
	},
}

var noArgsCmd = &cobra.Command{
	Use:  "noArgs",
	Args: cobra.NoArgs,
	// RunE 返回一个错误
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("noArgs run...")
		fmt.Println("args", args)
		return nil
	},
}

var onlyArgsCmd = &cobra.Command{
	Use:  "only",
	Args: cobra.OnlyValidArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("only run...")
		fmt.Println("args", args)
		return nil
	},
}

var exactArgsCmd = &cobra.Command{
	Use:  "exact",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("exact run...")
		fmt.Println("args", args)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkArgsCmd, noArgsCmd, onlyArgsCmd, exactArgsCmd)
}
