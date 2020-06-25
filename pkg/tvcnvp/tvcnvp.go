package tvcnvp

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const (
	// BaseURL consists of a couple of constructable fields
	// 0) Tool name
	// 1) Tool version
	// 2) Tool name again (as part of the filename)
	// 3) Tool version again (as part of the filename)
	// 4) Tool OS
	// 5) Tool arch
	BaseURL = `https://releases.hashicorp.com/%s/%d.%d.%d/%s_%d.%d.%d_%s_%s.zip`

	TerraformLatest = "0.12.28"
	VaultLatest     = "1.4.2"
	ConsulLatest    = "1.8.0"
	NomadLatest     = "0.11.3"
	VagrantLatest   = "2.2.9"
	PackerLatest    = "1.6.0"
)

type SemVer string

func (S SemVer) Valid() bool {
	v := regexp.MustCompile("^\\d\\.\\d{1-2}\\.\\d{1-2}$")
	return v.MatchString(string(S))
}

func (S SemVer) Major() int {
	v := strings.Split(string(S), ".")
	i, err := strconv.Atoi(v[0])
	if err != nil {
		panic(err)
	}
	return i
}

func (S SemVer) Minor() int {
	v := strings.Split(string(S), ".")
	i, err := strconv.Atoi(v[1])
	if err != nil {
		panic(err)
	}
	return i
}

func (S SemVer) Patch() int {
	v := strings.Split(string(S), ".")
	i, err := strconv.Atoi(v[2])
	if err != nil {
		panic(err)
	}
	return i
}

type OS string

type Arch string

type Name string

func (N Name) Parse() (string, error) {
	switch strings.ToLower(string(N)) {
	case "terraform":
		return strings.ToLower(string(N)), nil
	case "consul":
		return strings.ToLower(string(N)), nil
	case "vault":
		return strings.ToLower(string(N)), nil
	case "nomad":
		return strings.ToLower(string(N)), nil
	case "vagrant":
		return strings.ToLower(string(N)), nil
	case "packer":
		return strings.ToLower(string(N)), nil
	default:
		return "", fmt.Errorf("invalid product")
	}
}

type Product struct {
	Name        Name
	Latest      bool
	Version     SemVer
	downloadURL string
}

func (T *Product) Download() error {
	goos := strings.ToLower(runtime.GOOS)
	goarch := strings.ToLower(runtime.GOARCH)
	name, err := T.Name.Parse()
	if err != nil {
		return err
	}

	switch T.Latest {
	case true:
		T.Version = SemVer(TerraformLatest)
		T.downloadURL = fmt.Sprintf(
			BaseURL,
			name,
			T.Version.Major(),
			T.Version.Minor(),
			T.Version.Patch(),
			name,
			T.Version.Major(),
			T.Version.Minor(),
			T.Version.Patch(),
			goos,
			goarch)
		resp, err := http.Get(T.downloadURL)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return fmt.Errorf("status wasn't 200, was %d, url: %s", resp.StatusCode, T.downloadURL)
		}
		defer resp.Body.Close()

		f, err := os.OpenFile(path.Join(os.TempDir(), fmt.Sprintf("%s.zip", name)), os.O_CREATE|os.O_WRONLY, 0770)
		if err != nil {
			return err
		}
		defer f.Close()

		w, err := io.Copy(f, resp.Body)
		if err != nil {
			return err
		}

		fmt.Printf("successfully wrote %d bytes to %s\n", w, f.Name())
		return nil
	case false:
		if T.Version != "" {
			if T.Version.Valid() {
				T.downloadURL = fmt.Sprintf(
					BaseURL,
					name,
					T.Version.Major(),
					T.Version.Minor(),
					T.Version.Patch(),
					name,
					T.Version.Major(),
					T.Version.Minor(),
					T.Version.Patch(),
					goos,
					goarch)
				resp, err := http.Get(T.downloadURL)
				if err != nil {
					return err
				}
				if resp.StatusCode != 200 {
					return fmt.Errorf("status wasn't 200, was %d, url: %s", resp.StatusCode, T.downloadURL)
				}
				defer resp.Body.Close()

				f, err := os.OpenFile(path.Join(os.TempDir(), fmt.Sprintf("%s.zip", name)), os.O_CREATE|os.O_WRONLY, 0770)
				if err != nil {
					return err
				}
				defer f.Close()

				w, err := io.Copy(f, resp.Body)
				if err != nil {
					return err
				}

				fmt.Printf("Successfully wrote %d bytes to %s\n", w, f.Name())
				return nil
			}
			return fmt.Errorf("%s is not a valid version", T.Version)
		}
		return fmt.Errorf("Version must not be blank")
	}
	return fmt.Errorf("general error")
}
