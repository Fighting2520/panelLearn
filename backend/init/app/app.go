package app

import (
	"github.com/Fighting2520/panelLearn/backend/constant"
	"github.com/Fighting2520/panelLearn/backend/global"
	"github.com/Fighting2520/panelLearn/backend/utils/docker"
	"github.com/Fighting2520/panelLearn/backend/utils/files"
)

func Init() {
	dirs := []string{constant.DataDir, constant.ResourceDir, constant.AppResourceDir, constant.AppInstallDir,
		global.CONF.System.Backup, constant.RuntimeDir, constant.LocalAppResourceDir, constant.RemoteAppResourceDir}
	fileOp := files.NewFileOp()
	for _, dir := range dirs {
		createDir(fileOp, dir)
	}

	_ = docker.CreateDefaultDockerNetwork()
}

func createDir(fileOp files.FileOp, dirPath string) {
	if !fileOp.Stat(dirPath) {
		_ = fileOp.CreateDir(dirPath, 0755)
	}
}
