package architecture

import "github.com/ashiqsabith123/initNinja/pkg/helper"

func CleanCode(project_name string) {

	cleanCode := []string{
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

	helper.CreateFolder(project_name, cleanCode)

}
