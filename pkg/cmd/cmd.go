package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func installCMD(ccmd *cobra.Command, args []string) {
	v := map[string]string{
		"terraform-version": "",
		"vault-version":     "",
		"consul-version":    "",
		"nomad-version":     "",
		"vagrant-version":   "",
		"packer-version":    "",
	}

	for key := range v {
		value, _ := ccmd.LocalFlags().GetString(key)
		v[key] = value
	}

	inst, _ := ccmd.LocalFlags().GetBool("install")

	install(inst, v)
}
func updateCMD(ccmd *cobra.Command, args []string)    {}
func uninstallCMD(ccmd *cobra.Command, args []string) {}

func Run() {
	var (
		terraformVerFlag, vaultVerFlag, consulVerFlag, nomadVerFlag, vagrantVerFlag, packerVerFlag string

		installFlag bool

		mainCmd = &cobra.Command{
			Use:     "hashi.up",
			Version: "0.3.0",
			Args:    cobra.NoArgs,
			Run: func(ccmd *cobra.Command, args []string) {
				ccmd.Help()
			},
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

	mainCmd.AddCommand(installCmd, updateCmd, uninstallCmd)

	installCmd.Flags().StringVar(&terraformVerFlag, "terraform-version", "latest", "The version of terraform to install")
	installCmd.Flags().StringVar(&vaultVerFlag, "vault-version", "latest", "The version of vault to install")
	installCmd.Flags().StringVar(&consulVerFlag, "consul-version", "latest", "The version of consul to install")
	installCmd.Flags().StringVar(&nomadVerFlag, "nomad-version", "latest", "The version of nomad to install")
	installCmd.Flags().StringVar(&vagrantVerFlag, "vagrant-version", "latest", "The version of vagrant to install")
	installCmd.Flags().StringVar(&packerVerFlag, "packer-version", "latest", "The version of packer to install")
	installCmd.Flags().BoolVar(&installFlag, "install", true, "whether or not to install the tool after downloading it")

	if err := mainCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
