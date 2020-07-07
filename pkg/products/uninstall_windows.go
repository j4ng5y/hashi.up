// +build windows

package products

import (
	"fmt"
	"os"
)

func (P *Product) Uninstall() {
	h, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("could not determine user directory, err: %+v\n", err)
		return
	}

	p := fmt.Sprintf("C:\\Program Files\\hashicorp\\%s", string(P.Name))

	if err := os.Remove(p); err != nil {
		fmt.Printf("could not remove %s due to error: %+v\n", p, err)
	}
}
