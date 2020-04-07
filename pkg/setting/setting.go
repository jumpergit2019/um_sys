package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type App struct {
	PrefixUrl string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	RuntimeRootPath string
}

var AppSetting = &App{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalln("load conf failed, err: ", err)
		return
	}

	mapTo("app", AppSetting)

	return
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalln("load conf failed, err: ", err)
	}
}
