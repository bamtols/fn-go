package main

import (
	"fmt"
	"github.com/bamtols/fn-go/fn/fnPanic"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type (
	YamlVersion struct {
		Major uint `yaml:"major"`
		Minor uint `yaml:"minor"`
		Patch uint `yaml:"patch"`
		Save  uint `yaml:"save"`
	}
)

func main() {
	cmd := &cobra.Command{
		Use: "cliGit",
	}

	cmd.AddCommand(patchUp())

	fnPanic.HasError(cmd.Execute())
}

func saveUp() (root *cobra.Command) {
	root = &cobra.Command{
		Use: "save",
	}

	panic("notImpl")
}

func getYaml() {
	pwd := fnPanic.HasErrorOrValue(os.Getwd())
	fileNm := string(fnPanic.HasErrorOrValue(exec.Command("git", "branch", "--show-current").Output()))
	fileNm = filepath.Join(pwd, "version", fmt.Sprintf("%s.yaml", fileNm))

	_, err := os.Stat(fileNm)
	if err == nil {
	} else {

	}
}

func patchUp() (root *cobra.Command) {
	root = &cobra.Command{
		Use: "patch",
	}

	fileName := ""
	root.Flags().StringVarP(
		&fileName,
		"fileName",
		"v",
		fileName,
		fmt.Sprintf("fileName name: example=v1.0"),
	)

	root.Run = func(cmd *cobra.Command, args []string) {
		if fileName == "" {
			log.Panic("fileName not found")
		}

		f := loadYaml(fileName)
		f.Patch += 1
		saveVersion(fileName, f)

		fnPanic.HasError(exec.Command("git", "commit", "--all", "-m", fmt.Sprintf("%s version up", f.Tag())).Run())
		fnPanic.HasError(exec.Command("git", "push").Run())
		fnPanic.HasError(exec.Command("git", "tag", f.Tag()).Run())
		fnPanic.HasError(exec.Command("git", "push", "origin", f.Tag()).Run())

	}

	return
}

func loadYaml(branch string) (res *YamlVersion) {
	fp := getFilePath(branch)
	log.Printf("loadYaml: fp=%s\n", fp)

	file := fnPanic.HasErrorOrValue(os.Open(fp))
	defer file.Close()

	res = &YamlVersion{}
	fnPanic.HasError(yaml.NewDecoder(file).Decode(res))

	return
}

func saveVersion(branch string, yamlVersion *YamlVersion) {
	fp := getFilePath(branch)
	fnPanic.HasError(os.Remove(fp))
	log.Printf("delete file: fp=%s\n", fp)

	file := fnPanic.HasErrorOrValue(os.Create(fp))
	fnPanic.HasError(yaml.NewEncoder(file).Encode(yamlVersion))
	log.Printf("saved: branch=%s, major=%d, minor=%d, patch=%d\n", branch, yamlVersion.Major, yamlVersion.Minor, yamlVersion.Patch)
}

func getFilePath(fileName string) string {
	pwd := fnPanic.HasErrorOrValue(os.Getwd())
	return filepath.Join(pwd, fmt.Sprintf("version/%s.yaml", fileName))
}

func (x *YamlVersion) Tag() string {
	return fmt.Sprintf("v%d.%d.%d", x.Major, x.Minor, x.Patch)
}
