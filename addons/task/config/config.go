package config

import (
	"dzhgo/addons/fileUpload"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	dzhcore.SetVersions("task", fileUpload.Version)
}
