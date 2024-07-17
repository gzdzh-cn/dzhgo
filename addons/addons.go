package addons

import (
	"dzhgo/addons/dict"
	"dzhgo/addons/fileUpload"
)

func NewInit() {
	dict.NewInit()
	//space.NewInit()
	//task.NewInit()
	//member.NewInit()
	//crm.NewInit()
	fileUpload.NewInit()

}
