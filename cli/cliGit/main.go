package main

//
//import (
//	"fmt"
//	"github.com/bamtols/fn-go/fn/fnPanic"
//	"github.com/bamtols/fn-go/fn/fnParams"
//	"github.com/spf13/cobra"
//	"gopkg.in/yaml.v3"
//	"log"
//	"os"
//	"os/exec"
//	"path/filepath"
//	"strconv"
//	"strings"
//)
//
//type (
//	VersionFormat struct {
//		Version map[string]int `yaml:"version"`
//	}
//
//	Version struct {
//		Major int
//		Minor int
//		Patch int
//	}
//)
//
//const (
//	VersionFileNm = "tags.yaml"
//)
//
//func main() {
//	cmd := &cobra.Command{
//		Use: "cliGit",
//	}
//
//	cmd.AddCommand(patchUp())
//	cmd.AddCommand(saveUp())
//
//	fnPanic.HasError(cmd.Execute())
//}
//
//func saveUp() (root *cobra.Command) {
//	root = &cobra.Command{
//		Use: "save",
//	}
//
//	panic("notImpl")
//}
//
//func minorUp() (root *cobra.Command) {
//	panic("notImpl")
//	return
//}
//
//func majorUp() (root *cobra.Command) {
//	panic("notImpl")
//	return
//}
//
//func getYaml() (format *VersionFormat, err error) {
//	pwd := fnPanic.HasErrorOrValue(os.Getwd())
//
//	fileNm := filepath.Join(pwd, VersionFileNm)
//
//	var file *os.File
//	var res *VersionFormat
//	if _, err = os.Stat(VersionFileNm); err != nil {
//		if file, err = os.Open(fileNm); err != nil {
//			return nil, err
//		}
//		defer file.Close()
//
//		res = &VersionFormat{}
//		if err = yaml.NewDecoder(file).Decode(res); err != nil {
//			return nil, err
//		}
//	} else {
//		if file, err = os.Create(fileNm); err != nil {
//			return nil, err
//		}
//
//		defer file.Close()
//
//		res = &VersionFormat{
//			Version: make(map[string]int),
//		}
//
//		res.Version[getCurrentBranchNm()] = 0
//		if err = yaml.NewEncoder(file).Encode(res); err != nil {
//			return nil, err
//		}
//	}
//
//	return res, nil
//}
//
//func patchUp() (root *cobra.Command) {
//	root = &cobra.Command{
//		Use: "patch",
//	}
//
//	fileName := ""
//	root.Flags().StringVarP(
//		&fileName,
//		"fileName",
//		"v",
//		fileName,
//		fmt.Sprintf("fileName name: example=v1.0"),
//	)
//
//	root.Run = func(cmd *cobra.Command, args []string) {
//		if fileName == "" {
//			log.Panic("fileName not found")
//		}
//
//		f := loadYaml(fileName)
//		f.Patch += 1
//		saveVersion(fileName, f)
//
//		fnPanic.HasError(exec.Command("git", "commit", "--all", "-m", fmt.Sprintf("%s version up", f.Tag())).Run())
//		fnPanic.HasError(exec.Command("git", "push").Run())
//		fnPanic.HasError(exec.Command("git", "tag", f.Tag()).Run())
//		fnPanic.HasError(exec.Command("git", "push", "origin", f.Tag()).Run())
//
//	}
//
//	return
//}
//
//func loadYaml(branch string) (res *VersionFormat) {
//	fp := getFilePath(branch)
//	log.Printf("loadYaml: fp=%s\n", fp)
//
//	file := fnPanic.HasErrorOrValue(os.Open(fp))
//	defer file.Close()
//
//	res = &VersionFormat{}
//	fnPanic.HasError(yaml.NewDecoder(file).Decode(res))
//
//	return
//}
//
//func saveVersion(branch string, yamlVersion *VersionFormat) {
//	fp := getFilePath(branch)
//	fnPanic.HasError(os.Remove(fp))
//	log.Printf("delete file: fp=%s\n", fp)
//
//	file := fnPanic.HasErrorOrValue(os.Create(fp))
//	fnPanic.HasError(yaml.NewEncoder(file).Encode(yamlVersion))
//	log.Printf("saved: branch=%s, major=%d, minor=%d, patch=%d\n", branch, yamlVersion.Major, yamlVersion.Minor, yamlVersion.Patch)
//}
//
//func getFilePath(fileName string) string {
//	pwd := fnPanic.HasErrorOrValue(os.Getwd())
//	return filepath.Join(pwd, fmt.Sprintf("version/%s.yaml", fileName))
//}
//
//func (x *VersionFormat) Tag() string {
//	return fmt.Sprintf("v%d.%d.%d", x.Major, x.Minor, x.Patch)
//}
//
//func NewVersion(version ...string) *Version {
//	v := fnParams.Pick(version)
//
//	if v == "" {
//		v = "v0.0.1"
//	}
//
//	res := &Version{}
//	res = fnPanic.HasErrorOrValue(res.Replace(v))
//	return res
//}
//
//func (x *Version) Replace(version string) (*Version, error) {
//	version = version[1 : len(version)-1]
//
//	split := strings.Split(version, ".")
//	if len(split) != 3 {
//		return nil, fmt.Errorf("invalid version: version=%s", version)
//	}
//
//	var err error
//	if x.Major, err = strconv.Atoi(split[0]); err != nil {
//		return nil, err
//	}
//
//	if x.Minor, err = strconv.Atoi(split[1]); err != nil {
//		return nil, err
//	}
//
//	if x.Patch, err = strconv.Atoi(split[2]); err != nil {
//		return nil, err
//	}
//
//	return x, nil
//}
//
//func getCurrentBranchNm() string {
//	return string(fnPanic.HasErrorOrValue(exec.Command("git", "branch", "--show-current").Output()))
//}
