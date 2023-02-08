package main

import (
	"fmt"
	"github.com/bamtols/fn-go/fn/fnParams"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type (
	FileMng struct {
		fp string
	}

	Version struct {
		Major int
		Minor int
		Patch int
		Raw   string
	}

	FileFormat struct {
		List VersMap `yaml:"list"`
	}

	VersMap map[string]int
)

func NewFileMng(fileNm string, rootPath ...string) *FileMng {
	p := fnParams.Pick(rootPath)
	if p == "" {
		var pwd, err = os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		p = pwd
	}

	return &FileMng{
		fp: filepath.Join(p, fileNm),
	}
}

func (x *FileMng) Open() (res *FileFormat, err error) {
	res = &FileFormat{
		List: make(VersMap),
	}

	var file *os.File

	if !x.hasFile() {
		return
	}

	if file, err = os.Open(x.fp); err != nil {
		return
	}

	defer file.Close()

	if err = yaml.NewDecoder(file).Decode(res); err != nil {
		return
	}

	return
}

func (x *FileMng) Save(data *FileFormat) (err error) {
	if x.hasFile() {
		if err = os.Remove(x.fp); err != nil {
			return
		}
	}

	var file *os.File
	if file, err = os.Create(x.fp); err != nil {
		return
	}

	defer file.Close()

	if err = yaml.NewEncoder(file).Encode(data); err != nil {
		return
	}

	return
}

func (x *FileMng) hasFile() bool {
	_, err := os.Stat(x.fp)
	return err == nil
}

func (x *FileMng) GetGitBranchNm() (*Version, error) {
	// format v0.0.0 방식
	byteStr, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return nil, err
	}

	return NewVersion(string(byteStr))
}

func NewVersion(vers string) (res *Version, err error) {
	res = &Version{
		Raw: vers,
	}
	vers = strings.Replace(vers, "v", "", 1)
	vers = strings.Replace(vers, "\n", "", -1)
	ls := strings.Split(vers, ".")

	if len(ls) != 3 {
		return nil, fmt.Errorf("invalidFormat: version=%s", vers)
	}

	if res.Major, err = strconv.Atoi(ls[0]); err != nil {
		return nil, err
	}

	if res.Minor, err = strconv.Atoi(ls[1]); err != nil {
		return nil, err
	}

	if res.Patch, err = strconv.Atoi(ls[2]); err != nil {
		return nil, err
	}

	return
}
