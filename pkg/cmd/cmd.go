package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func Run() {
	var (
		terraformVerFlag, vaultVerFlag, consulVerFlag, nomadVerFlag, vagrantVerFlag, packerVerFlag string
		mainCmd                                                                                    = &cobra.Command{
			Use:     "hashi.up",
			Version: "0.1.0",
			Args:    cobra.NoArgs,
			Run:     downloadCMD,
		}
	)

	mainCmd.Flags().StringVar(&terraformVerFlag, "terraform-version", "latest", "The version of terraform to install")
	mainCmd.Flags().StringVar(&vaultVerFlag, "vault-version", "latest", "The version of vault to install")
	mainCmd.Flags().StringVar(&consulVerFlag, "consul-version", "latest", "The version of consul to install")
	mainCmd.Flags().StringVar(&nomadVerFlag, "nomad-version", "latest", "The version of nomad to install")
	mainCmd.Flags().StringVar(&vagrantVerFlag, "vagrant-version", "latest", "The version of vagrant to install")
	mainCmd.Flags().StringVar(&packerVerFlag, "packer-version", "latest", "The version of packer to install")

	if err := mainCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
