package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userLicense string
var configFile string

var rootCmd = &cobra.Command{
	Use:   "mycobra",
	Short: "命令短描述",
	Long:  "命令长描述",
}

func init() {
	// 每个命令执行之前调用的方法
	cobra.OnInitialize(initConfig)

	// 持久标志: 当前命令及子命令都可用
	rootCmd.PersistentFlags().Bool("viper", true, "是否使用 viper 配置文件")               // 参数完整形式
	rootCmd.PersistentFlags().StringP("author", "a", "ll", "作者")                   // 参数简写形式
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "配置文件")         // 参数指针完整形式
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "授权信息") // 参数指针简写形式
	rootCmd.PersistentFlags().String("host", "", "host")
	rootCmd.PersistentFlags().Int("port", 6379, "port")

	// 本地标志: 当前命令可用
	rootCmd.Flags().StringP("source", "s", "", "来源") // 本地标志

	// 绑定 viper 配置文件和命令标志
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.SetConfigType("yaml")
		viper.SetConfigName(configFile)
		viper.AddConfigPath(home)
	}

	viper.AutomaticEnv() // 自动绑定环境变量
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("使用 viper config file:", viper.ConfigFileUsed())
}

func Execute() {
	rootCmd.Execute()
}
