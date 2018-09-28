package main

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	packagePath = "github.com/grepory/gentest"
	apisPath    = "pkg/apis"
)

func main() {
	apiGroups, err := ioutil.ReadDir(filepath.Join(".", apisPath))
	if err != nil {
		panic(err)
	}

	for _, group := range apiGroups {
		if group.IsDir() {
			groupName := group.Name()
			versions, err := ioutil.ReadDir(filepath.Join(".", apisPath, groupName))
			if err != nil {
				panic(err)
			}

			for _, dir := range versions {
				versionName := dir.Name()
				if strings.HasPrefix(versionName, "v") { // actually a version?
					fqPackage := strings.Join([]string{
						packagePath,
						apisPath,
						groupName,
						versionName,
					}, "/")

					// client-gen

					args := []string{
						"-i", fqPackage,
						"-p", strings.Join([]string{packagePath, "pkg", "client", "clientset"}, "/"),
						"--input", strings.Join([]string{groupName, versionName}, "/"),
						"--input-base", strings.Join([]string{packagePath, apisPath}, "/"),
					}
					exec.Command("client-gen", args...).Run()

					// register-gen
					args = []string{
						"-i", fqPackage,
					}
					exec.Command("register-gen", args...).Run()

					// defaulter-gen
					exec.Command("defaulter-gen", args...).Run()
				}
			}
		}
	}
}
