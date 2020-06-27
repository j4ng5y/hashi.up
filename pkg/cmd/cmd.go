package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func argsContains(args []string, contains string) bool {
	for _, v := range args {
		if strings.ToLower(v) == strings.ToLower(contains) {
			return true
		}
	}
	return false
}

func updateCMD(ccmd *cobra.Command, args []string)    {}
func uninstallCMD(ccmd *cobra.Command, args []string) {}

// Run will run the CLI
//
// Arguments:
//     None
//
// Returns:
//     None
func Run() {
	var (
		terraformVersion, consulVersion, nomadVersion, vagrantVersion, vaultVersion, packerVersion string
		mainCmd                                                                                    = &cobra.Command{
			Use:     "hashi.up",
			Version: "0.4.0",
			Args:    cobra.NoArgs,
			Run: func(ccmd *cobra.Command, args []string) {
				ccmd.Help()
			},
		}
		installCmd = &cobra.Command{
			Use:   "install",
			Short: "Install one, multiple, or all hashicorp tools.",
			Example: `
  # Install all hashicorp tools at their latest versions
  hashi up install all

  # Install the latest version of Terraform
  hashi.up install terraform

  # Install just Terraform and Packer at their latest versions
  hashi.up install terraform packer

  # Install a specific version of Terraform	
  hashi.up install terraform --terraform-version=0.11.0

  # Full options
  hashi.up install <[terraform[consul[nomad[vagrant[vault[packer]]]]]]|all> [--terraform-version=<version|latest>[--nomad-version=<version|latest>[--consul-version=<version|latest>[--vagrant-version=<version|latest>[--vault-version=<version|latest>[--packer-version=<version|latest>]]]]]]`,
			Args: cobra.OnlyValidArgs,
			ValidArgs: []string{
				"terraform",
				"consul",
				"nomad",
				"vagrant",
				"vault",
				"packer",
				"all",
			},
			Run: installCMD,
		}
		updateCmd = &cobra.Command{
			Use:   "update",
			Short: "Check for and subsiquently update one, multiple, or all installed hashicorp tools.",
			Example: `
  # Update all hashicorp tools to their latest versions
  hashi.up update all

  # Update to the latest version of Terraform
  hashi.up update terraform

  # Update just Terraform and Packer to their latest versions
  hashi.up update terraform packer

  # Update to a specific version of Terraform
  hashi.up update terraform --terraform-version=0.12.28

  # Full options
  hashi.up update <[terraform[consul[nomad[vagrant[vault[packer]]]]]]|all> [--terraform-version=<version|latest>[--nomad-version=<version|latest>[--consul-version=<version|latest>[--vagrant-version=<version|latest>[--vault-version=<version|latest>[--packer-version=<version|latest>]]]]]]`,
			Args: cobra.OnlyValidArgs,
			ValidArgs: []string{
				"terraform",
				"consul",
				"nomad",
				"vagrant",
				"vault",
				"packer",
				"all",
			},
			Run: updateCMD,
		}
		uninstallCmd = &cobra.Command{
			Use:   "uninstall",
			Short: "Uninstall one, multiple, or all installed hashicorp tools.",
			Example: `
  # Uninstall all hashicorp tools
  hashi.up unistall all

  # Uninstall Terraform
  hashi.up uninstall terraform

  # Uninstall Terraform and Packer
  hashi.up uninstall terraorm packer

  # Full options
  hashi.up uninstall <[terraform[consul[nomad[vagrant[vault[packer]]]]]]|all>`,
			Args: cobra.OnlyValidArgs,
			ValidArgs: []string{
				"terraform",
				"consul",
				"nomad",
				"vagrant",
				"vault",
				"packer",
				"all",
			},
			Run: uninstallCMD,
		}
	)

	mainCmd.AddCommand(installCmd, updateCmd, uninstallCmd)

	installCmd.Flags().StringVar(&terraformVersion, "terraform-version", "latest", "the terraform version to install")
	installCmd.Flags().StringVar(&consulVersion, "consul-version", "latest", "the consul version to install")
	installCmd.Flags().StringVar(&nomadVersion, "nomad-version", "latest", "the nomad version to install")
	installCmd.Flags().StringVar(&vagrantVersion, "vagrant-version", "latest", "the vagrant version to install")
	installCmd.Flags().StringVar(&vaultVersion, "vault-version", "latest", "the vault version to install")
	installCmd.Flags().StringVar(&packerVersion, "packer-version", "latest", "the packer version to install")

	updateCmd.Flags().StringVar(&terraformVersion, "terraform-version", "latest", "the terraform version to update to")
	updateCmd.Flags().StringVar(&consulVersion, "consul-version", "latest", "the consul version to update to")
	updateCmd.Flags().StringVar(&nomadVersion, "nomad-version", "latest", "the nomad version to update to")
	updateCmd.Flags().StringVar(&vagrantVersion, "vagrant-version", "latest", "the vagrant version to update to")
	updateCmd.Flags().StringVar(&vaultVersion, "vault-version", "latest", "the vault version to update to")
	updateCmd.Flags().StringVar(&packerVersion, "packer-version", "latest", "the packer version to update to")

	if err := mainCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
