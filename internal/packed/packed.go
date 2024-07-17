package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}