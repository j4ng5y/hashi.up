package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func runAll(ccmd *cobra.Command, args []string) {
	downloadCMD(ccmd, args)
	installCMD(ccmd, args)
}

func Run() {
	var (
		terraformVerFlag, vaultVerFlag, consulVerFlag, nomadVerFlag, vagrantVerFlag, packerVerFlag string

		mainCmd = &cobra.Command{
			Use:     "hashi.up",
			Version: "0.1.0",
			Args:    cobra.NoArgs,
			Run:     runAll,
		}
		downloadCmd = &cobra.Command{
			Use:  "download",
			Args: cobra.NoArgs,
			Run:  downloadCMD,
		}
		installCmd = &cobra.Command{
			Use:  "install",
			Args: cobra.NoArgs,
			Run:  installCMD,
		}
		updateCmd = &cobra.Command{
			Use:  "update",
			Args: cobra.NoArgs,
			Run:  updateCMD,
		}
		uninstallCmd = &cobra.Command{
			Use:  "uninstall",
			Args: cobra.NoArgs,
			Run:  uninstallCMD,
		}
	)

	mainCmd.AddCommand(downloadCmd, installCmd, updateCmd, uninstallCmd)

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
