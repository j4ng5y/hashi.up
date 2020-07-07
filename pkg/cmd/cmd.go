package cmd

import (
	"log"
	"sync"

	"github.com/j4ng5y/hashi.up/pkg/products"

	"github.com/spf13/cobra"
)

// Run runs the CLI
func Run() {
	var (
		downloadOnly, allTools                                                                                                               bool
		consulInstallVersion, nomadInstallVersion, packerInstallVersion, terraformInstallVersion, vagrantInstallVersion, vaultInstallVersion string
		consulUpdatedVersion, nomadUpdatedVersion, packerUpdatedVersion, terraformUpdatedVersion, vagrantUpdatedVersion, vaultUpdatedVersion string
		uninstallConsul, uninstallNomad, uninstallPacker, uninstallTerraform, uninstallVagrant, uninstallVault                               bool

		hashiupCommand = &cobra.Command{
			Use:     "hashi.up",
			Args:    cobra.NoArgs,
			Short:   "A hashicorp tool downloader/installer/updater/uninstaller.",
			Run:     func(ccmd *cobra.Command, args []string) {},
			Version: "0.4.0",
		}
		installCommand = &cobra.Command{
			Use:   "install",
			Args:  cobra.NoArgs,
			Short: "Install one, some, or all of the hashicorp tools",
			Run: func(ccmd *cobra.Command, args []string) {
				var (
					wg                                               = &sync.WaitGroup{}
					CONSUL, NOMAD, PACKER, TERRAFORM, VAGRANT, VAULT *products.Product
					err                                              error
				)

				CONSULV, err := ccmd.Flags().GetString("consul")
				NOMADV, err := ccmd.Flags().GetString("nomad")
				PACKERV, err := ccmd.Flags().GetString("packer")
				TERRAFORMV, err := ccmd.Flags().GetString("terraform")
				VAGRANTV, err := ccmd.Flags().GetString("vagrant")
				VAULTV, err := ccmd.Flags().GetString("vault")

				if err != nil {
					log.Fatal(err)
				}

				if allTools {
					CONSUL, err = products.New("consul", products.SemVer(products.ConsulLatest))
					NOMAD, err = products.New("nomad", products.SemVer(products.NomadLatest))
					PACKER, err = products.New("packer", products.SemVer(products.PackerLatest))
					TERRAFORM, err = products.New("terraform", products.SemVer(products.TerraformLatest))
					VAGRANT, err = products.New("vagrant", products.SemVer(products.VagrantLatest))
					VAULT, err = products.New("vault", products.SemVer(products.VaultLatest))

					if err != nil {
						log.Fatal(err)
					}

					wg.Add(6)
					go CONSUL.Download(wg)
					go NOMAD.Download(wg)
					go PACKER.Download(wg)
					go TERRAFORM.Download(wg)
					go VAGRANT.Download(wg)
					go VAULT.Download(wg)
					wg.Wait()

					if downloadOnly {
						return
					}

					CONSUL.Install()
					NOMAD.Install()
					PACKER.Install()
					TERRAFORM.Install()
					VAGRANT.Install()
					VAULT.Install()
					return
				}
				if CONSULV != "" {
					wg.Add(1)
					CONSUL, err = products.New("consul", products.SemVer(CONSULV))
					CONSUL.Download(wg)
					if downloadOnly {
						return
					}
					CONSUL.Install()
				}
				if NOMADV != "" {
					wg.Add(1)
					NOMAD, err = products.New("nomad", products.SemVer(NOMADV))
					NOMAD.Download(wg)
					if downloadOnly {
						return
					}
					NOMAD.Install()
				}
				if PACKERV != "" {
					wg.Add(1)
					PACKER, err = products.New("packer", products.SemVer(PACKERV))
					PACKER.Download(wg)
					if downloadOnly {
						return
					}
					PACKER.Install()
				}
				if TERRAFORMV != "" {
					wg.Add(1)
					TERRAFORM, err = products.New("terraform", products.SemVer(TERRAFORMV))
					TERRAFORM.Download(wg)
					if downloadOnly {
						return
					}
					TERRAFORM.Install()
				}
				if VAGRANTV != "" {
					wg.Add(1)
					VAGRANT, err = products.New("vagrant", products.SemVer(VAGRANTV))
					VAGRANT.Download(wg)
					if downloadOnly {
						return
					}
					VAGRANT.Install()
				}
				if VAULTV != "" {
					wg.Add(1)
					VAULT, err = products.New("vault", products.SemVer(VAULTV))
					VAULT.Download(wg)
					if downloadOnly {
						return
					}
					VAULT.Install()
				}
			},
		}
		updateCommand = &cobra.Command{
			Use:   "update",
			Args:  cobra.NoArgs,
			Short: "Update one, some, or all of the hashicorp tools",
			Run: func(ccmd *cobra.Command, args []string) {
				var (
					wg                                               = &sync.WaitGroup{}
					CONSUL, NOMAD, PACKER, TERRAFORM, VAGRANT, VAULT *products.Product
					err                                              error
				)

				CONSULV, err := ccmd.Flags().GetString("consul")
				NOMADV, err := ccmd.Flags().GetString("nomad")
				PACKERV, err := ccmd.Flags().GetString("packer")
				TERRAFORMV, err := ccmd.Flags().GetString("terraform")
				VAGRANTV, err := ccmd.Flags().GetString("vagrant")
				VAULTV, err := ccmd.Flags().GetString("vault")

				if err != nil {
					log.Fatal(err)
				}

				if allTools {
					CONSUL, err = products.New("consul", products.SemVer(products.ConsulLatest))
					NOMAD, err = products.New("nomad", products.SemVer(products.NomadLatest))
					PACKER, err = products.New("packer", products.SemVer(products.PackerLatest))
					TERRAFORM, err = products.New("terraform", products.SemVer(products.TerraformLatest))
					VAGRANT, err = products.New("vagrant", products.SemVer(products.VagrantLatest))
					VAULT, err = products.New("vault", products.SemVer(products.VaultLatest))

					if err != nil {
						log.Fatal(err)
					}

					wg.Add(6)
					go CONSUL.Download(wg)
					go NOMAD.Download(wg)
					go PACKER.Download(wg)
					go TERRAFORM.Download(wg)
					go VAGRANT.Download(wg)
					go VAULT.Download(wg)
					wg.Wait()

					if downloadOnly {
						return
					}

					CONSUL.Install()
					NOMAD.Install()
					PACKER.Install()
					TERRAFORM.Install()
					VAGRANT.Install()
					VAULT.Install()
					return
				}
				if CONSULV != "" {
					wg.Add(1)
					CONSUL, err = products.New("consul", products.SemVer(CONSULV))
					CONSUL.Download(wg)
					if downloadOnly {
						return
					}
					CONSUL.Install()
				}
				if NOMADV != "" {
					wg.Add(1)
					NOMAD, err = products.New("nomad", products.SemVer(NOMADV))
					NOMAD.Download(wg)
					if downloadOnly {
						return
					}
					NOMAD.Install()
				}
				if PACKERV != "" {
					wg.Add(1)
					PACKER, err = products.New("packer", products.SemVer(PACKERV))
					PACKER.Download(wg)
					if downloadOnly {
						return
					}
					PACKER.Install()
				}
				if TERRAFORMV != "" {
					wg.Add(1)
					TERRAFORM, err = products.New("terraform", products.SemVer(TERRAFORMV))
					TERRAFORM.Download(wg)
					if downloadOnly {
						return
					}
					TERRAFORM.Install()
				}
				if VAGRANTV != "" {
					wg.Add(1)
					VAGRANT, err = products.New("vagrant", products.SemVer(VAGRANTV))
					VAGRANT.Download(wg)
					if downloadOnly {
						return
					}
					VAGRANT.Install()
				}
				if VAULTV != "" {
					wg.Add(1)
					VAULT, err = products.New("vault", products.SemVer(VAULTV))
					VAULT.Download(wg)
					if downloadOnly {
						return
					}
					VAULT.Install()
				}
			},
		}
		uninstallCommand = &cobra.Command{
			Use:   "remove",
			Args:  cobra.NoArgs,
			Short: "Remove one, some, or all of the hashicorp tools",
			Run:   func(ccmd *cobra.Command, args []string) {},
		}
	)

	hashiupCommand.AddCommand(installCommand, updateCommand, uninstallCommand)
	installCommand.Flags().BoolVar(&allTools, "all", false, "download or install all available hashicorp tools at their latest version")
	installCommand.Flags().BoolVar(&downloadOnly, "download-only", false, "only download the hashicorp tools, but don't install them")
	installCommand.Flags().StringVar(&consulInstallVersion, "consul", "", "download/install consul at a specific version")
	installCommand.Flags().StringVar(&nomadInstallVersion, "nomad", "", "the nomad version to install, leave blank for latest")
	installCommand.Flags().StringVar(&packerInstallVersion, "packer", "", "the packer version to install, leave blank for latest")
	installCommand.Flags().StringVar(&terraformInstallVersion, "terraform", "", "the terraform version to install, leave blank for latest")
	installCommand.Flags().StringVar(&vagrantInstallVersion, "vagrant", "", "the vagrant version to install, leave blank for latest")
	installCommand.Flags().StringVar(&vaultInstallVersion, "vault", "", "the vault version to install, leave blank for latest")
	installCommand.Flags().Lookup("consul").NoOptDefVal = products.ConsulLatest
	installCommand.Flags().Lookup("nomad").NoOptDefVal = products.NomadLatest
	installCommand.Flags().Lookup("packer").NoOptDefVal = products.PackerLatest
	installCommand.Flags().Lookup("terraform").NoOptDefVal = products.TerraformLatest
	installCommand.Flags().Lookup("vagrant").NoOptDefVal = products.VagrantLatest
	installCommand.Flags().Lookup("vault").NoOptDefVal = products.VaultLatest

	updateCommand.Flags().StringVar(&consulUpdatedVersion, "consul", products.ConsulLatest, "the consul version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&nomadUpdatedVersion, "nomad", products.NomadLatest, "the nomad version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&packerUpdatedVersion, "packer", products.PackerLatest, "the packer version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&terraformUpdatedVersion, "terraform", products.TerraformLatest, "the terraform version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&vagrantUpdatedVersion, "vagrant", products.VagrantLatest, "the vagrant version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&vaultUpdatedVersion, "vault", products.VaultLatest, "the vault version to update to, leave blank for latest")
	updateCommand.Flags().Lookup("consul").NoOptDefVal = products.ConsulLatest
	updateCommand.Flags().Lookup("nomad").NoOptDefVal = products.NomadLatest
	updateCommand.Flags().Lookup("packer").NoOptDefVal = products.PackerLatest
	updateCommand.Flags().Lookup("terraform").NoOptDefVal = products.TerraformLatest
	updateCommand.Flags().Lookup("vagrant").NoOptDefVal = products.VagrantLatest
	updateCommand.Flags().Lookup("vault").NoOptDefVal = products.VaultLatest

	uninstallCommand.Flags().BoolVar(&uninstallConsul, "consul", false, "uninstall consul")
	uninstallCommand.Flags().BoolVar(&uninstallNomad, "nomad", false, "uninstall consul")
	uninstallCommand.Flags().BoolVar(&uninstallPacker, "packer", false, "uninstall consul")
	uninstallCommand.Flags().BoolVar(&uninstallTerraform, "terraform", false, "uninstall consul")
	uninstallCommand.Flags().BoolVar(&uninstallVagrant, "vagrant", false, "uninstall consul")
	uninstallCommand.Flags().BoolVar(&uninstallVault, "vault", false, "uninstall consul")

	if err := hashiupCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
