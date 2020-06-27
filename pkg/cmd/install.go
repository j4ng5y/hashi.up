package cmd

import (
	"log"
	"strings"
	"sync"

	"github.com/j4ng5y/hashi.up/pkg/products"
	"github.com/spf13/cobra"
)

func installCMD(ccmd *cobra.Command, args []string) {
	var Terraform, Nomad, Packer, Consul, Vault, Vagrant *products.Product
	var err error
	wg := new(sync.WaitGroup)

	switch {
	case argsContains(args, "all"):
		if len(args) > 1 {
			log.Println("\"all\" can exist all by itself, no need to do extra work :D")
		}

		if tfver, _ := ccmd.Flags().GetString("terraform-version"); strings.ToLower(tfver) == "latest" {
			Terraform, err = products.New(products.Name("terraform"), products.SemVer(products.TerraformLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Terraform, err = products.New(products.Name("terraform"), products.SemVer(tfver))
			if err != nil {
				log.Fatal(err)
			}
		}

		if cver, _ := ccmd.Flags().GetString("consul-version"); strings.ToLower(cver) == "latest" {
			Consul, err = products.New(products.Name("consul"), products.SemVer(products.ConsulLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Consul, err = products.New(products.Name("consul"), products.SemVer(cver))
			if err != nil {
				log.Fatal(err)
			}
		}

		if nver, _ := ccmd.Flags().GetString("nomad-version"); strings.ToLower(nver) == "latest" {
			Nomad, err = products.New(products.Name("nomad"), products.SemVer(products.NomadLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Nomad, err = products.New(products.Name("terraform"), products.SemVer(nver))
			if err != nil {
				log.Fatal(err)
			}
		}

		if vagver, _ := ccmd.Flags().GetString("vagrant-version"); strings.ToLower(vagver) == "latest" {
			Vagrant, err = products.New(products.Name("vagrant"), products.SemVer(products.VagrantLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Vagrant, err = products.New(products.Name("terraform"), products.SemVer(vagver))
			if err != nil {
				log.Fatal(err)
			}
		}

		if vauver, _ := ccmd.Flags().GetString("vault-version"); strings.ToLower(vauver) == "latest" {
			Vault, err = products.New(products.Name("vault"), products.SemVer(products.VaultLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Vault, err = products.New(products.Name("terraform"), products.SemVer(vauver))
			if err != nil {
				log.Fatal(err)
			}
		}

		if pver, _ := ccmd.Flags().GetString("packer-version"); strings.ToLower(pver) == "latest" {
			Packer, err = products.New(products.Name("packer"), products.SemVer(products.PackerLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Packer, err = products.New(products.Name("terraform"), products.SemVer(pver))
			if err != nil {
				log.Fatal(err)
			}
		}

		wg.Add(6)
		go Terraform.Download(wg)
		go Nomad.Download(wg)
		go Consul.Download(wg)
		go Vault.Download(wg)
		go Vagrant.Download(wg)
		go Packer.Download(wg)
		wg.Wait()
		return
	case argsContains(args, "terraform"):
		if tfver, _ := ccmd.Flags().GetString("terraform-version"); strings.ToLower(tfver) == "latest" {
			Terraform, err = products.New(products.Name("terraform"), products.SemVer(products.TerraformLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Terraform, err = products.New(products.Name("terraform"), products.SemVer(tfver))
			if err != nil {
				log.Fatal(err)
			}
		}
		wg.Add(1)
		Terraform.Download(wg)
	case argsContains(args, "nomad"):
		if nver, _ := ccmd.Flags().GetString("nomad-version"); strings.ToLower(nver) == "latest" {
			Nomad, err = products.New(products.Name("nomad"), products.SemVer(products.NomadLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Nomad, err = products.New(products.Name("terraform"), products.SemVer(nver))
			if err != nil {
				log.Fatal(err)
			}
		}
		wg.Add(1)
		Nomad.Download(wg)
	case argsContains(args, "consul"):
		if cver, _ := ccmd.Flags().GetString("consul-version"); strings.ToLower(cver) == "latest" {
			Consul, err = products.New(products.Name("consul"), products.SemVer(products.ConsulLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Consul, err = products.New(products.Name("consul"), products.SemVer(cver))
			if err != nil {
				log.Fatal(err)
			}
		}
		wg.Add(1)
		Consul.Download(wg)
	case argsContains(args, "vagrant"):
		if vagver, _ := ccmd.Flags().GetString("vagrant-version"); strings.ToLower(vagver) == "latest" {
			Vagrant, err = products.New(products.Name("vagrant"), products.SemVer(products.VagrantLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Vagrant, err = products.New(products.Name("terraform"), products.SemVer(vagver))
			if err != nil {
				log.Fatal(err)
			}
		}
		wg.Add(1)
		Vagrant.Download(wg)
	case argsContains(args, "vault"):
		if vauver, _ := ccmd.Flags().GetString("vault-version"); strings.ToLower(vauver) == "latest" {
			Vault, err = products.New(products.Name("vault"), products.SemVer(products.VaultLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Vault, err = products.New(products.Name("terraform"), products.SemVer(vauver))
			if err != nil {
				log.Fatal(err)
			}
		}
		wg.Add(1)
		Vault.Download(wg)
	case argsContains(args, "packer"):
		if pver, _ := ccmd.Flags().GetString("packer-version"); strings.ToLower(pver) == "latest" {
			Packer, err = products.New(products.Name("packer"), products.SemVer(products.PackerLatest))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			Packer, err = products.New(products.Name("terraform"), products.SemVer(pver))
			if err != nil {
				log.Fatal(err)
			}
		}
		wg.Add(1)
		Packer.Download(wg)
	}
}
