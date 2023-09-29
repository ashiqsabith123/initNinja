package architecture

import (
	"github.com/ashiqsabith123/initNinja/pkg/helper"
)

func CleanCode(project_name, language string) {

	folderStructure := []string{
		"cmd/api",
		"pkg/api/handler",
		"pkg/api/middleware",
		"pkg/api/routes",
		"pkg/config",
		"pkg/db",
		"pkg/di",
		"pkg/repository/interfaces",
		"pkg/usecases/interfaces",
	}

	go_files := []string{
		"cmd/api/main.go",
		"pkg/api/server.go",
		"pkg/config/config.go",
		"pkg/di/wire.go",
	}

	helper.CreateFolder(project_name, folderStructure)

	switch language {
	case "GO":
		helper.CreateFiles(project_name, go_files)
	}

}
