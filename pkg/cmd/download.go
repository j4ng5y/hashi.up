package cmd

import (
	"sync"

	"github.com/j4ng5y/hashi.up/pkg/products"
	"github.com/spf13/cobra"
)

func downloadCMD(ccmd *cobra.Command, args []string) {
	wg := &sync.WaitGroup{}
	var TF, VAU, C, N, VAG, P *products.Product

	tfv, _ := ccmd.LocalFlags().GetString("terraform-version")
	vauv, _ := ccmd.LocalFlags().GetString("vault-version")
	cv, _ := ccmd.LocalFlags().GetString("consul-version")
	nv, _ := ccmd.LocalFlags().GetString("nomad-version")
	vagv, _ := ccmd.LocalFlags().GetString("vagrant-version")
	pv, _ := ccmd.LocalFlags().GetString("packer-version")

	if tfv == "latest" {
		wg.Add(1)
		TF = &products.Product{
			Name:    "terraform",
			Version: products.TerraformLatest,
		}
		go TF.Download(false, wg)
	} else if tfv == "none" {
		wg.Add(1)
		TF = &products.Product{}
		go TF.Download(true, wg)
	} else {
		wg.Add(1)
		TF = &products.Product{
			Name:    "terraform",
			Version: products.SemVer(tfv),
		}
		go TF.Download(false, wg)
	}

	if pv == "latest" {
		wg.Add(1)
		P = &products.Product{
			Name:    "packer",
			Version: products.PackerLatest,
		}
		go P.Download(false, wg)
	} else if pv == "none" {
		wg.Add(1)
		P = &products.Product{}
		go P.Download(true, wg)
	} else {
		wg.Add(1)
		P = &products.Product{
			Name:    "packer",
			Version: products.SemVer(pv),
		}
		go P.Download(false, wg)
	}

	if vagv == "latest" {
		wg.Add(1)
		VAG = &products.Product{
			Name:    "vagrant",
			Version: products.VagrantLatest,
		}
		go VAG.Download(false, wg)
	} else if vagv == "none" {
		wg.Add(1)
		VAG = &products.Product{}
		go VAG.Download(true, wg)
	} else {
		wg.Add(1)
		VAG = &products.Product{
			Name:    "vagrant",
			Version: products.SemVer(vagv),
		}
		go VAG.Download(false, wg)
	}

	if nv == "latest" {
		wg.Add(1)
		N = &products.Product{
			Name:    "nomad",
			Version: products.NomadLatest,
		}
		go N.Download(false, wg)
	} else if nv == "none" {
		wg.Add(1)
		N = &products.Product{}
		go N.Download(true, wg)
	} else {
		wg.Add(1)
		N = &products.Product{
			Name:    "nomad",
			Version: products.SemVer(nv),
		}
		go N.Download(false, wg)
	}

	if cv == "latest" {
		wg.Add(1)
		C = &products.Product{
			Name:    "consul",
			Version: products.ConsulLatest,
		}
		go C.Download(false, wg)
	} else if cv == "none" {
		wg.Add(1)
		C = &products.Product{}
		go C.Download(true, wg)
	} else {
		wg.Add(1)
		C = &products.Product{
			Name:    "consul",
			Version: products.SemVer(cv),
		}
		go C.Download(false, wg)
	}

	if vauv == "latest" {
		wg.Add(1)
		VAU = &products.Product{
			Name:    "vault",
			Version: products.VaultLatest,
		}
		go VAU.Download(false, wg)
	} else if vauv == "none" {
		wg.Add(1)
		VAU = &products.Product{}
		go VAU.Download(true, wg)
	} else {
		wg.Add(1)
		VAU = &products.Product{
			Name:    "vault",
			Version: products.SemVer(vauv),
		}
		go VAU.Download(false, wg)
	}
	wg.Wait()
}
