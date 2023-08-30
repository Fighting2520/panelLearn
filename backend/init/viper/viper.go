package viper

import (
	"bytes"
	"fmt"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"

	"github.com/1Panel-dev/1Panel/backend/utils/cmd"

	"github.com/1Panel-dev/1Panel/cmd/server/conf"
	"github.com/Fighting2520/panelLearn/backend/configs"
	"github.com/Fighting2520/panelLearn/backend/global"
	"github.com/Fighting2520/panelLearn/backend/utils/files"
	"github.com/spf13/viper"
	"sigs.k8s.io/yaml"
)

func Init() {
	baseDir := "/opt"
	port := "9999"
	mode := ""
	version := "v1.0.0"
	username, password, entrance := "", "", ""
	fileOp := files.NewFileOp()
	v := viper.NewWithOptions()
	v.SetConfigType("yaml")

	config := configs.ServerConfig{}
	if err := yaml.Unmarshal(conf.AppYaml, &config); err != nil {
		panic(err)
	}
	if config.System.Mode != "" {
		mode = config.System.Mode
	}
	if mode == "dev" && fileOp.Stat("/opt/1panel/conf/app.yaml") {
		v.SetConfigName("app")
		v.AddConfigPath(path.Join("/opt/1panel/conf"))
		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	} else {
		baseDir = loadParams("BASE_DIR")
		port = loadParams("ORIGINAL_PORT")
		version = loadParams("ORIGINAL_VERSION")
		username = loadParams("ORIGINAL_USERNAME")
		password = loadParams("ORIGINAL_PASSWORD")
		entrance = loadParams("ORIGINAL_ENTRANCE")

		reader := bytes.NewReader(conf.AppYaml)
		if err := v.ReadConfig(reader); err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.CONF); err != nil {
			panic(err)
		}
	})
	serverConfig := configs.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	if mode == "dev" && fileOp.Stat("/opt/1panel/conf/app.yaml") {
		if serverConfig.System.BaseDir != "" {
			baseDir = serverConfig.System.BaseDir
		}
		if serverConfig.System.Port != "" {
			port = serverConfig.System.Version
		}
	}
}

func loadParams(param string) string {
	stdout, err := cmd.Execf("grep '^%s=' /usr/bin/1pctl | cut -d'=' -f2", param)
	if err != nil {
		panic(err)
	}
	info := strings.ReplaceAll(stdout, "\n", "")
	if len(info) == 0 || info == `""` {
		panic(fmt.Sprintf("error `%s` find in /usr/bin/1pctl", param))
	}
	return info
}
