package products

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"launchpad.net/xmlpath"
)

func GetLatest(product string) (v string) {
	r, err := http.Get(fmt.Sprintf("https://releases.hashicorp.com/%s", strings.ToLower(string(product))))
	if err != nil {
		return v
	}
	defer r.Body.Close()

	n, err := xmlpath.ParseHTML(r.Body)
	if err != nil {
		return v
	}

	for i := 2; i < 10; i++ {
		path := xmlpath.MustCompile(fmt.Sprintf("/html/body/ul/li[%d]", i))
		if value, ok := path.String(n); ok {
			s := strings.Split(value, "_")
			if SemVer(strings.TrimSpace(s[1])).Valid() {
				v = strings.TrimSpace(s[1])
				break
			}
			continue
		} else {
			continue
		}
	}
	return v
}

func (T *Product) Download(wg *sync.WaitGroup) {
	defer wg.Done()

	goos := strings.ToLower(runtime.GOOS)
	goarch := strings.ToLower(runtime.GOARCH)
	name, err := T.Name.Parse()
	if err != nil {
		log.Println(err)
		return
	}

	if T.Version != "" {
		if T.Version.Valid() {
			switch goos {
			case "windows":
				switch goarch {
				case "amd64":
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
						"x86",
						"64")
				case "386":
					T.downloadURL = strings.TrimSuffix(fmt.Sprintf(
						BaseURL,
						name,
						T.Version.Major(),
						T.Version.Minor(),
						T.Version.Patch(),
						name,
						T.Version.Major(),
						T.Version.Minor(),
						T.Version.Patch(),
						"i686"), "_")
				default:
					log.Println("unsupported OS Architecture")
					return
				}
			default:
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
			}

			resp, err := http.Get(T.downloadURL)
			if err != nil {
				log.Println(err)
				return
			}
			if resp.StatusCode != 200 {
				log.Printf("status wasn't 200, was %d, url: %s\n", resp.StatusCode, T.downloadURL)
				return
			}
			defer resp.Body.Close()

			f, err := os.OpenFile(path.Join(os.TempDir(), fmt.Sprintf("%s.zip", name)), os.O_CREATE|os.O_WRONLY, 0770)
			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()

			_, err = io.Copy(f, resp.Body)
			if err != nil {
				log.Println(err)
				return
			}

			T.downloadedPath = f.Name()
			fmt.Printf("Successfully downloaded %s v%s to %s\n", T.Name, T.Version, f.Name())
			return
		}
		log.Printf("%s is not a valid version\n", T.Version)
		return
	}
	log.Println("Version must not be blank")
	return
}
