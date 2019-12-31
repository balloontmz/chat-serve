package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run the serve",
	Long:  `run serve use tags like --port`,
	Run:   runServe,
}

func init() {
	serveCmd.Flags().String("port", "1323", "serve port default 1323")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
}

//runServe 启动网络服务
func runServe(cmd *cobra.Command, args []string) {
	fmt.Println("这里启动了服务器,当前获取的 port 为:", viper.GetString("port"))
}
