package config

import (
	"dzhgo/addons/fileUpload"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	dzhcore.SetVersions("space", fileUpload.Version)
}
