package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/mcdaigle1/slaslosli/modules/fileutils"
)

type Config struct {
    CurrentVendor string                  `yaml:"current_vendor"`
    Vendors       map[string]VendorConfig `yaml:"vendors"`
    Aws           Aws                     `yaml:"aws"`
    LogLevel      string                  `yaml:"log_level"`
}

type VendorConfig struct {
    URL     string            `yaml:"url"`
    Queries map[string]string `yaml:"queries"`
}

type Aws struct {
    SsoProfile string   `yaml:"sso_profile"`
    SecretARN  string   `yaml:"secret_arn"`
    SecretName string   `yaml:"secret_name"`
    Region     string   `yaml:"region"`
}

var Global Config

func Load() {
    workingDir, err := fileutils.GetWorkingDir()
    if err != nil {
        log.Fatalf("Error getting working dir when initially populating config: %v", err)
    }

    file, err := os.ReadFile(workingDir + "/config/config.yaml")
    if err != nil {
        log.Fatalf("Error reading config file " + workingDir + "/config/config.yaml when initially populating config: %v", err)
    }

    err = yaml.Unmarshal(file, &Global)
    if err != nil {
        log.Fatalf("Error unmarshalling config when initially populating config: %v", err)
    }

    fmt.Printf("Current vendor: %s\n", Global.CurrentVendor)
    fmt.Printf("Vendor url: %s\n", Global.Vendors[Global.CurrentVendor].URL)
}
