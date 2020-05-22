package config

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/commitdev/zero/pkg/util/exit"
)

const exampleConfig = `name: %s

# Context is normally populated automatically but could be used to inject global params
context:

# module can be in any format the go-getter supports (path, github, url, etc.)
# supports https://github.com/hashicorp/go-getter#url-format
# Example:
# - source: "../development/modules/ci"
# - output: "github-actions"

modules:
  - source: "github.com/commitdev/zero-aws-eks-stack"`

func CreateExample(projectName string) {
	content := []byte(fmt.Sprintf(exampleConfig, projectName))

	err := ioutil.WriteFile(path.Join(projectName, "zero.yml"), content, 0644)
	if err != nil {
		exit.Fatal("Failed to create example zero.yml")
	}
}
