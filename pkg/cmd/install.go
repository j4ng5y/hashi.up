package cmd

import (
	"sync"

	"github.com/j4ng5y/hashi.up/pkg/products"
)

func install(install bool, versions map[string]string) {
	wg := &sync.WaitGroup{}
	var TF, VAU, C, N, VAG, P *products.Product

	tfv := versions["terraform-version"]
	vauv := versions["vault-version"]
	cv := versions["consul-version"]
	nv := versions["nomad-version"]
	vagv := versions["vagrant-version"]
	pv := versions["packer-version"]

	if tfv == "latest" {
		wg.Add(1)
		TF = &products.Product{
			Name:    "terraform",
			Version: products.TerraformLatest,
		}
		if install {
			go TF.Download(false, true, wg)
		} else {
			go TF.Download(false, false, wg)
		}
	} else if tfv == "none" {
		wg.Add(1)
		TF = &products.Product{}
		go TF.Download(true, false, wg)
	} else {
		wg.Add(1)
		TF = &products.Product{
			Name:    "terraform",
			Version: products.SemVer(tfv),
		}
		if install {
			go TF.Download(false, true, wg)
		} else {
			go TF.Download(false, false, wg)
		}
	}

	if pv == "latest" {
		wg.Add(1)
		P = &products.Product{
			Name:    "packer",
			Version: products.PackerLatest,
		}
		if install {
			go P.Download(false, true, wg)
		} else {
			go P.Download(false, false, wg)
		}
	} else if pv == "none" {
		wg.Add(1)
		P = &products.Product{}
		go P.Download(true, false, wg)
	} else {
		wg.Add(1)
		P = &products.Product{
			Name:    "packer",
			Version: products.SemVer(pv),
		}
		if install {
			go P.Download(false, true, wg)
		} else {
			go P.Download(false, false, wg)
		}
	}

	if vagv == "latest" {
		wg.Add(1)
		VAG = &products.Product{
			Name:    "vagrant",
			Version: products.VagrantLatest,
		}
		if install {
			go VAG.Download(false, true, wg)
		} else {
			go VAG.Download(false, false, wg)
		}
	} else if vagv == "none" {
		wg.Add(1)
		VAG = &products.Product{}
		go VAG.Download(true, false, wg)
	} else {
		wg.Add(1)
		VAG = &products.Product{
			Name:    "vagrant",
			Version: products.SemVer(vagv),
		}
		if install {
			go VAG.Download(false, true, wg)
		} else {
			go VAG.Download(false, false, wg)
		}
	}

	if nv == "latest" {
		wg.Add(1)
		N = &products.Product{
			Name:    "nomad",
			Version: products.NomadLatest,
		}
		if install {
			go N.Download(false, true, wg)
		} else {
			go N.Download(false, false, wg)
		}
	} else if nv == "none" {
		wg.Add(1)
		N = &products.Product{}
		go N.Download(true, false, wg)
	} else {
		wg.Add(1)
		N = &products.Product{
			Name:    "nomad",
			Version: products.SemVer(nv),
		}
		if install {
			go N.Download(false, true, wg)
		} else {
			go N.Download(false, false, wg)
		}
	}

	if cv == "latest" {
		wg.Add(1)
		C = &products.Product{
			Name:    "consul",
			Version: products.ConsulLatest,
		}
		if install {
			go C.Download(false, true, wg)
		} else {
			go C.Download(false, false, wg)
		}
	} else if cv == "none" {
		wg.Add(1)
		C = &products.Product{}
		go C.Download(true, false, wg)
	} else {
		wg.Add(1)
		C = &products.Product{
			Name:    "consul",
			Version: products.SemVer(cv),
		}
		if install {
			go C.Download(false, true, wg)
		} else {
			go C.Download(false, false, wg)
		}
	}

	if vauv == "latest" {
		wg.Add(1)
		VAU = &products.Product{
			Name:    "vault",
			Version: products.VaultLatest,
		}
		if install {
			go VAU.Download(false, true, wg)
		} else {
			go VAU.Download(false, false, wg)
		}
	} else if vauv == "none" {
		wg.Add(1)
		VAU = &products.Product{}
		go VAU.Download(true, false, wg)
	} else {
		wg.Add(1)
		VAU = &products.Product{
			Name:    "vault",
			Version: products.SemVer(vauv),
		}
		if install {
			go VAU.Download(false, true, wg)
		} else {
			go VAU.Download(false, false, wg)
		}
	}
	wg.Wait()
}
