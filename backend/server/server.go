package server

import (
	"github.com/1Panel-dev/1Panel/backend/init/db"
	"github.com/1Panel-dev/1Panel/backend/init/viper"
	"github.com/Fighting2520/panelLearn/backend/init/app"
	"github.com/Fighting2520/panelLearn/backend/init/log"
)

func Start() {
	viper.Init()
	log.Init()
	app.Init()
	db.Init()
}
