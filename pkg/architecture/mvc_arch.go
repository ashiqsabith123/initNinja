package architecture

import (
	"github.com/ashiqsabith123/initNinja/pkg/helper"
)

func Mvc(project_name, language string) {

	folderStructure := []string{
		"app/controllers/models",
		"app/controllers/views",
		"app/config",
		"app/routes",
		"app/static/css",
		"app/static/js",
		"app/templates/user",
		"app/templates/post",
	}

	go_files := []string{
		"app/config/config.go",
		"app/config/database.go",
		"app/routes/routes.go",
		"main.go",
		"Readme.md",
	}

	helper.CreateFolder(project_name, folderStructure)

	switch language {
	case "GO":
		helper.CreateFiles(project_name, go_files)
	}

}
