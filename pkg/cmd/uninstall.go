package cmd

import (
	"fmt"

	"github.com/j4ng5y/hashi.up/pkg/products"
)

func uninstall(uninstallTerraform, uninstallVault, uninstallConsul, uninstallNomad, uninstallVagrant, uninstallPacker bool) {
	P := &products.Product{}
	switch {
	case uninstallTerraform && uninstallVault && uninstallConsul && uninstallNomad && uninstallVagrant && uninstallPacker:
		P.UninstallAll()
	case uninstallTerraform:
		P.Uninstall("terraform")
	case uninstallVault:
		P.Uninstall("vault")
	case uninstallConsul:
		P.Uninstall("consul")
	case uninstallNomad:
		P.Uninstall("nomad")
	case uninstallVagrant:
		P.Uninstall("vagrant")
	case uninstallPacker:
		P.Uninstall("packer")
	default:
		fmt.Println("I don't understand...")
	}
}
