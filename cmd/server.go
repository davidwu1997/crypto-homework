package cmd

import (
	"crypto/pkg/service/crypto"
	"fmt"
	"os"

	cobra "github.com/spf13/cobra"
)

// ServerCmd http server
var ServerCmd = &cobra.Command{
	Run:           runServerCmd,
	Use:           "server",
	Short:         "Start Crypto server",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func runServerCmd(cmd *cobra.Command, args []string) {
	app, err := crypto.Initialize()
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
