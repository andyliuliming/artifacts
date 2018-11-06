package deployments

import (
	"gopkg.in/yaml.v2"
)

const (
	JumpBoxRole        = "builtin/jumpbox"
	GlustrerFSNodeRole = "builtin/gluster"
)

type Role struct {
	Name      string        `yaml:"name"`
	Variables yaml.MapSlice `yaml:"variables"`
}
