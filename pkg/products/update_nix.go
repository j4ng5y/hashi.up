// +build !windows

package products

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

func (P *Product) checkUpdate(tool string) (needsUpdate bool, err error) {
	err = nil
	needsUpdate = false

	h, err := os.UserHomeDir()
	if err != nil {
		return needsUpdate, err
	}

	p := path.Join(h, ".local", "hashicorp", tool)

	c := exec.Command(p, "-v")
	out, err := c.Output()
	if err != nil {
		return needsUpdate, err
	}

	if strings.Contains(string(out), "out of date!") {
		needsUpdate = true
		return needsUpdate, err
	}
	return needsUpdate, err
}

func (P *Product) Update(tool string) error {
	needsUpdate, err := P.checkUpdate(tool)
	if err != nil {
		return err
	}
	if needsUpdate {
		P.Name = Name(tool)
		P.Download(&sync.WaitGroup{})
	}
	return nil
}
