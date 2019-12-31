package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "chat",
		Short: "这是一个 chat serve 端程序",
		Long: `
			这是一个 chat serve 端程序
		这是一些介绍
				`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Print("进入此处\n")
		// 	fmt.Println("viper 配置文件为:", viper.Get("a"))
		// },
	}
)

//Execute 执行主命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//初始化时添加配置,程序运行时能获取,此函数后无法获取.
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $pwd/config.ini)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找当前文件夹
		pwd, _ := os.Getwd()

		// 将当前文件夹下的 config 加入配置文件
		viper.AddConfigPath(pwd)
		viper.SetConfigName("config")
		viper.SetConfigType("json")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("读取配置文件错误,错误原因为:", err)
	}
}
