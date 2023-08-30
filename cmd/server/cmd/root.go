package cmd

import (
	"github.com/spf13/cobra"
)

func init() {}

var RootCmd = &cobra.Command{
	Use:   "1panel",
	Short: "1Panel ，一款现代化的 Linux 面板",
	RunE: func(cmd *cobra.Command, args []string) error {
		server.Start()
	},
}
