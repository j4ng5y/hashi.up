// +build windows

package products

import (
	"fmt"
	"os"
)

func (P *Product) Uninstall() {
	p := fmt.Sprintf("C:\\Program Files\\hashicorp\\%s", string(P.Name))

	if err := os.Remove(p); err != nil {
		fmt.Printf("could not remove %s due to error: %+v\n", p, err)
	}
}
