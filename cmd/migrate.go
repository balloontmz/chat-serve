package cmd

import (
	"fmt"

	"github.com/balloontmz/chat-serve/app/models"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dbMigrateCmd)
}

var (
	dbMigrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "run the db migrate",
		Long:  `run migrate use tags like`,
		Run:   runMigrate,
	}
)

func init() {
}

//runServe 启动网络服务
func runMigrate(cmd *cobra.Command, args []string) {
	fmt.Print("================启动数据库迁移================")
	models.Migrate()
	fmt.Print("================数据库迁移完成================")
}
