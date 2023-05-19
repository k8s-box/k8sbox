package formatters

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/k8s-box/k8sbox/pkg/k8sbox/structs"
)

type TomlFormatter struct {
	GetEnvironmentFromToml func(string) (structs.Environment, string, error)
}

func NewTomlFormatter() TomlFormatter {
	return TomlFormatter{
		GetEnvironmentFromToml: getEnvironmentFromToml,
	}
}

func getEnvironmentFromToml(tomlFile string) (structs.Environment, string, error) {
	var environment structs.Environment

	info, err := os.Stat(tomlFile)
	if err != nil {
		return structs.Environment{}, "", errors.New(fmt.Sprintf("File %s not found", tomlFile))
	}
	boxesPath := strings.ReplaceAll(tomlFile, info.Name(), "")
	data, err := os.ReadFile(tomlFile)

	err = toml.Unmarshal(data, &environment)
	if err != nil {
		panic(err)
	}
	return environment, boxesPath, nil
}
