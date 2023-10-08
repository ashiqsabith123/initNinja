package architecture

import (
	"github.com/ashiqsabith123/initNinja/pkg/config"
	"github.com/ashiqsabith123/initNinja/pkg/helper"
)

func DDD(project_name, language string) {

	folderStructure := config.GetFolderStructure("ddd_arch")

	helper.CreateFolder(project_name, folderStructure)

	switch language {
	case "GO":
		helper.CreateFiles(project_name, config.GetFolderStructure("go_ddd_files"))
	}

}
