package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func Run() {
	var (
		mainCmd = &cobra.Command{
			Use:     "hashi.up",
			Version: "0.1.0",
			Args:    cobra.NoArgs,
			Run:     func(ccmd *cobra.Command, args []string) {},
		}
	)

	if err := mainCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
