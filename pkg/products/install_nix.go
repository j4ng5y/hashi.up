// +build !windows

package products

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func (P *Product) setPath() error {
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
		log.Printf("unable to determine shell, writing to %s", path.Join(h, ".profile"))
		f, err := os.OpenFile(path.Join(h, ".profile"), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "bash":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "csh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "ksh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "zsh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "tcsh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "fish":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "ion":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "dash":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "eshell":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "rc":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "scsh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "xonsh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	case "oh":
		f, err := os.OpenFile(path.Join(h, fmt.Sprintf(".%src", sh)), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if strings.Contains(string(b), "# hashi.up") {
			return nil
		} else {
			if _, err := f.WriteString(fmt.Sprintf(p, h)); err != nil {
				return err
			}
			log.Println("successfully updated path")
		}
	}
	return nil
}

func (P *Product) Install() {
	h, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return
	}

	dest := fmt.Sprintf("%s/.local/hashicorp", h)
	r, err := zip.OpenReader(P.downloadedPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := r.Close(); err != nil {
			log.Println(err)
			return
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("successful installed %s in %s", P.Name, dest)
	}

	if err := P.setPath(); err != nil {
		log.Println(err)
		return
	}

	return
}
