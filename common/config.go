package common

import (
	"fmt"
	. "github.com/Unknwon/goconfig"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ReadConfig(fileName string) *ConfigFile {
	var root string
	curFilename := os.Args[0]
	binaryPath, err := exec.LookPath(curFilename)
	if err != nil {
		fmt.Errorf("LookPath err: %s ", err)
	}

	binaryPath, err = filepath.Abs(binaryPath)
	if err != nil {
		fmt.Errorf("filepath err: %s ", err)
	}

	root = filepath.Dir(filepath.Dir(binaryPath))

	configPath := root + fileName

	if !fileExist(configPath) {
		curDir, _ := os.Getwd()
		pos := strings.LastIndex(curDir, "src")
		if pos == -1 {
			// panic("can't find " + mainIniPath)
			fmt.Errorf("can't find %s ", fileName)
		} else {
			root = curDir[:pos]
			configPath = root + fileName
		}
	}
	configFile, err := LoadConfigFile(configPath)
	if err != nil {
		fmt.Errorf("read file err: %s", err)
		return nil
	}
	return configFile
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
