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
)

func (T *Product) Download(skip bool, wg *sync.WaitGroup) {
	defer wg.Done()
	if skip {
		return
	}

	goos := strings.ToLower(runtime.GOOS)
	goarch := strings.ToLower(runtime.GOARCH)
	name, err := T.Name.Parse()
	if err != nil {
		log.Println(err)
		return
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
			log.Print(err)
			return
		}
		defer f.Close()

		w, err := io.Copy(f, resp.Body)
		if err != nil {
			log.Println(err)
			return
		}

		T.downloadedPath = f.Name()
		fmt.Printf("successfully wrote %d bytes to %s\n", w, f.Name())
		return
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

				w, err := io.Copy(f, resp.Body)
				if err != nil {
					log.Println(err)
					return
				}

				T.downloadedPath = f.Name()
				fmt.Printf("Successfully wrote %d bytes to %s\n", w, f.Name())
				return
			}
			log.Printf("%s is not a valid version\n", T.Version)
			return
		}
		log.Println("Version must not be blank")
		return
	}
	log.Println("general error")
	return
}
