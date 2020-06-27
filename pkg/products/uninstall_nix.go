// +build !windows

package products

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func (P *Product) unsetPath() error {
	h, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	p := `
# hashi.up
export PATH=$PATH:%s/.local/hashicorp
`

	_, sh := path.Split(os.Getenv("SHELL"))

	switch strings.ToLower(sh) {
	default:
		log.Printf("unable to determine shell, checking %s", path.Join(h, ".profile"))
		f, err := os.OpenFile(path.Join(h, ".profile"), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "bash":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "csh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "ksh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "zsh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "tcsh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "fish":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "ion":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "dash":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "eshell":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "rc":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "scsh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "xonsh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	case "oh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			s := strings.ReplaceAll(string(b), p, "")
			_, err := f.Write([]byte(s))
			if err != nil {
				return err
			}
			fmt.Printf("PATH is already set, skipping setting PATH\n")
			return nil
		} else {
			fmt.Printf("PATH is unset, exiting\n")
			return nil
		}
	}
	return nil
}

func (P *Product) Uninstall(tool string) {
	h, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("could not determine user directory, err: %+v\n", err)
		return
	}

	p := path.Join(h, ".local", "hashicorp", tool)

	if err := os.Remove(p); err != nil {
		fmt.Printf("could not remove %s due to error: %+v\n", tool, err)
	}
}

func (P *Product) UninstallAll() {
	h, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("could not determine user directory, err: %+v\n", err)
		return
	}

	p := path.Join(h, ".local", "hashicorp")

	if err := os.RemoveAll(p); err != nil {
		fmt.Printf("could not remove %s due to error: %+v\n", p, err)
		return
	}

	if err := P.unsetPath(); err != nil {
		fmt.Printf("could not remove PATH directive, you may have to clean this up yourself: %+v\n", err)
		return
	}
}
