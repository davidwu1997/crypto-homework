package cmd

import (
	"crypto/pkg/service/migrate"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// MigrateCmd migrate mysql
var MigrateCmd = &cobra.Command{
	Run:           runMigrateCmd,
	Use:           "migrate",
	Short:         "Migrate Database",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func runMigrateCmd(cmd *cobra.Command, args []string) {
	app, err := migrate.Initialize(cfgFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = app.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
