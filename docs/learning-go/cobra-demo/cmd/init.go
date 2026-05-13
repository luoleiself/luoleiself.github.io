package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "init 子命令短描述",
	Long:    "init 子命令长描述",
	Aliases: []string{"create"},
	// RunE 返回一个错误
	// Run 没有返回值
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init run...")
		fmt.Printf("author is %s\nviper is %v\nlicense %v\n",
			cmd.Flags().Lookup("author").Value,
			cmd.Flags().Lookup("viper").Value,
			cmd.Flags().Lookup("license").Value,
		)
		fmt.Printf("backend is %v\ndelete is %v\n",
			cmd.Flags().Lookup("backend").Value,
			cmd.PersistentFlags().Lookup("delete").Value,
		)

		fmt.Println("viper.GetString(\"author\")", viper.GetString("author"))
		fmt.Println("viper.GetString(\"host\")", viper.GetString("host"))
		fmt.Println("viper.GetInt(\"port\")", viper.GetInt("port"))

		fmt.Println("args", args)
	},
}

func init() {
	// 本地标志
	initCmd.Flags().BoolP("backend", "b", true, "后台运行")

	// 持久标志
	initCmd.PersistentFlags().BoolP("delete", "D", true, "强制删除")

	rootCmd.AddCommand(initCmd)
}
