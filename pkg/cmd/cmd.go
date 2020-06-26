package cmd

import (
	"log"

	"github.com/j4ng5y/hashi.up/pkg/tvcnvp"
	"github.com/spf13/cobra"
)

func Run() {
	var (
		terraformVerFlag, vaultVerFlag, consulVerFlag, nomadVerFlag, vagrantVerFlag, packerVerFlag string
		mainCmd                                                                                    = &cobra.Command{
			Use:     "hashi.up",
			Version: "0.1.0",
			Args:    cobra.NoArgs,
			Run: func(ccmd *cobra.Command, args []string) {
				var TF, VAU, C, N, VAG, P *tvcnvp.Product
				if terraformVerFlag == "latest" {
					TF = &tvcnvp.Product{
						Name:    "terraform",
						Version: tvcnvp.TerraformLatest,
					}
				} else {
					TF = &tvcnvp.Product{
						Name:    "terraform",
						Version: tvcnvp.SemVer(terraformVerFlag),
					}
				}

				if packerVerFlag == "latest" {
					P = &tvcnvp.Product{
						Name:    "packer",
						Version: tvcnvp.PackerLatest,
					}
				} else {
					P = &tvcnvp.Product{
						Name:    "packer",
						Version: tvcnvp.SemVer(packerVerFlag),
					}
				}

				if vagrantVerFlag == "latest" {
					VAG = &tvcnvp.Product{
						Name:    "vagrant",
						Version: tvcnvp.VagrantLatest,
					}
				} else {
					VAG = &tvcnvp.Product{
						Name:    "vagrant",
						Version: tvcnvp.SemVer(vagrantVerFlag),
					}
				}

				if nomadVerFlag == "latest" {
					N = &tvcnvp.Product{
						Name:    "nomad",
						Version: tvcnvp.NomadLatest,
					}
				} else {
					N = &tvcnvp.Product{
						Name:    "nomad",
						Version: tvcnvp.SemVer(nomadVerFlag),
					}
				}

				if consulVerFlag == "latest" {
					C = &tvcnvp.Product{
						Name:    "consul",
						Version: tvcnvp.ConsulLatest,
					}
				} else {
					C = &tvcnvp.Product{
						Name:    "consul",
						Version: tvcnvp.SemVer(consulVerFlag),
					}
				}

				if vaultVerFlag == "latest" {
					VAU = &tvcnvp.Product{
						Name:    "vault",
						Version: tvcnvp.VaultLatest,
					}
				} else {
					VAU = &tvcnvp.Product{
						Name:    "vault",
						Version: tvcnvp.SemVer(vaultVerFlag),
					}
				}

				if err := TF.Download(); err != nil {
					log.Fatal(err)
				}
				if err := VAU.Download(); err != nil {
					log.Fatal(err)
				}
				if err := C.Download(); err != nil {
					log.Fatal(err)
				}
				if err := N.Download(); err != nil {
					log.Fatal(err)
				}
				if err := VAG.Download(); err != nil {
					log.Fatal(err)
				}
				if err := P.Download(); err != nil {
					log.Fatal(err)
				}
			},
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
