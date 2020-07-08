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
		ConsulLatest, NomadLatest, PackerLatest, TerraformLatest, VagrantLatest, VaultLatest                                                 string
		downloadOnly, allTools                                                                                                               bool
		consulInstallVersion, nomadInstallVersion, packerInstallVersion, terraformInstallVersion, vagrantInstallVersion, vaultInstallVersion string
		consulUpdatedVersion, nomadUpdatedVersion, packerUpdatedVersion, terraformUpdatedVersion, vagrantUpdatedVersion, vaultUpdatedVersion string
		uninstallConsul, uninstallNomad, uninstallPacker, uninstallTerraform, uninstallVagrant, uninstallVault                               bool

		hashiupCommand = &cobra.Command{
			Use:     "hashi.up",
			Args:    cobra.NoArgs,
			Short:   "A hashicorp tool downloader/installer/updater/uninstaller.",
			Run:     func(ccmd *cobra.Command, args []string) {},
			Version: "0.4.3",
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

				if consulInstallVersion == "" && nomadInstallVersion == "" && packerInstallVersion == "" && terraformInstallVersion == "" && vagrantInstallVersion == "" && vaultInstallVersion == "" {
					allTools = true
				}

				if allTools {
					CONSUL, err = products.New("consul", products.SemVer(ConsulLatest))
					NOMAD, err = products.New("nomad", products.SemVer(NomadLatest))
					PACKER, err = products.New("packer", products.SemVer(PackerLatest))
					TERRAFORM, err = products.New("terraform", products.SemVer(TerraformLatest))
					VAGRANT, err = products.New("vagrant", products.SemVer(VagrantLatest))
					VAULT, err = products.New("vault", products.SemVer(VaultLatest))

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

				if consulUpdatedVersion == "" && nomadUpdatedVersion == "" && packerUpdatedVersion == "" && terraformUpdatedVersion == "" && vagrantUpdatedVersion == "" && vaultUpdatedVersion == "" {
					allTools = true
				}

				if allTools {
					CONSUL, err = products.New("consul", products.SemVer(ConsulLatest))
					NOMAD, err = products.New("nomad", products.SemVer(NomadLatest))
					PACKER, err = products.New("packer", products.SemVer(PackerLatest))
					TERRAFORM, err = products.New("terraform", products.SemVer(TerraformLatest))
					VAGRANT, err = products.New("vagrant", products.SemVer(VagrantLatest))
					VAULT, err = products.New("vault", products.SemVer(VaultLatest))

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
			Run: func(ccmd *cobra.Command, args []string) {
				CONSUL, err := products.New("consul", products.SemVer("0.0.0"))
				NOMAD, err := products.New("nomad", products.SemVer("0.0.0"))
				PACKER, err := products.New("packer", products.SemVer("0.0.0"))
				TERRAFORM, err := products.New("terraform", products.SemVer("0.0.0"))
				VAGRANT, err := products.New("vagrant", products.SemVer("0.0.0"))
				VAULT, err := products.New("vault", products.SemVer("0.0.0"))

				if err != nil {
					log.Fatal(err)
				}

				if !uninstallConsul && !uninstallNomad && !uninstallPacker && !uninstallTerraform && !uninstallVagrant && !uninstallVault {
					CONSUL.Uninstall()
					NOMAD.Uninstall()
					PACKER.Uninstall()
					TERRAFORM.Uninstall()
					VAGRANT.Uninstall()
					VAULT.Uninstall()
					log.Printf("uninstalled all tools")
				}
				if uninstallConsul {
					CONSUL.Uninstall()
					log.Printf("uninstalled consul")
				}
				if uninstallNomad {
					NOMAD.Uninstall()
					log.Printf("uninstalled nomad")
				}
				if uninstallPacker {
					PACKER.Uninstall()
					log.Printf("uninstalled packer")
				}
				if uninstallTerraform {
					TERRAFORM.Uninstall()
					log.Printf("uninstalled terraform")
				}
				if uninstallVagrant {
					VAGRANT.Uninstall()
					log.Printf("uninstalled vagrant")
				}
				if uninstallVault {
					VAULT.Uninstall()
					log.Printf("uninstalled vault")
				}
			},
		}
	)

	ConsulLatest = products.GetLatest("consul")
	NomadLatest = products.GetLatest("nomad")
	PackerLatest = products.GetLatest("packer")
	TerraformLatest = products.GetLatest("terraform")
	VagrantLatest = products.GetLatest("vagrant")
	VaultLatest = products.GetLatest("vault")

	hashiupCommand.AddCommand(installCommand, updateCommand, uninstallCommand)
	installCommand.Flags().BoolVar(&downloadOnly, "download-only", false, "only download the hashicorp tools, but don't install them")
	installCommand.Flags().StringVar(&consulInstallVersion, "consul", "", "download/install consul at a specific version")
	installCommand.Flags().StringVar(&nomadInstallVersion, "nomad", "", "the nomad version to install, leave blank for latest")
	installCommand.Flags().StringVar(&packerInstallVersion, "packer", "", "the packer version to install, leave blank for latest")
	installCommand.Flags().StringVar(&terraformInstallVersion, "terraform", "", "the terraform version to install, leave blank for latest")
	installCommand.Flags().StringVar(&vagrantInstallVersion, "vagrant", "", "the vagrant version to install, leave blank for latest")
	installCommand.Flags().StringVar(&vaultInstallVersion, "vault", "", "the vault version to install, leave blank for latest")
	installCommand.Flags().Lookup("consul").NoOptDefVal = ConsulLatest
	installCommand.Flags().Lookup("nomad").NoOptDefVal = NomadLatest
	installCommand.Flags().Lookup("packer").NoOptDefVal = PackerLatest
	installCommand.Flags().Lookup("terraform").NoOptDefVal = TerraformLatest
	installCommand.Flags().Lookup("vagrant").NoOptDefVal = VagrantLatest
	installCommand.Flags().Lookup("vault").NoOptDefVal = VaultLatest

	updateCommand.Flags().StringVar(&consulUpdatedVersion, "consul", "", "the consul version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&nomadUpdatedVersion, "nomad", "", "the nomad version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&packerUpdatedVersion, "packer", "", "the packer version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&terraformUpdatedVersion, "terraform", "", "the terraform version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&vagrantUpdatedVersion, "vagrant", "", "the vagrant version to update to, leave blank for latest")
	updateCommand.Flags().StringVar(&vaultUpdatedVersion, "vault", "", "the vault version to update to, leave blank for latest")
	updateCommand.Flags().Lookup("consul").NoOptDefVal = ConsulLatest
	updateCommand.Flags().Lookup("nomad").NoOptDefVal = NomadLatest
	updateCommand.Flags().Lookup("packer").NoOptDefVal = PackerLatest
	updateCommand.Flags().Lookup("terraform").NoOptDefVal = TerraformLatest
	updateCommand.Flags().Lookup("vagrant").NoOptDefVal = VagrantLatest
	updateCommand.Flags().Lookup("vault").NoOptDefVal = VaultLatest

	uninstallCommand.Flags().BoolVar(&uninstallConsul, "consul", false, "uninstall consul")
	uninstallCommand.Flags().BoolVar(&uninstallNomad, "nomad", false, "uninstall nomad")
	uninstallCommand.Flags().BoolVar(&uninstallPacker, "packer", false, "uninstall packer")
	uninstallCommand.Flags().BoolVar(&uninstallTerraform, "terraform", false, "uninstall terraform")
	uninstallCommand.Flags().BoolVar(&uninstallVagrant, "vagrant", false, "uninstall vagrant")
	uninstallCommand.Flags().BoolVar(&uninstallVault, "vault", false, "uninstall vault")

	if err := hashiupCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
