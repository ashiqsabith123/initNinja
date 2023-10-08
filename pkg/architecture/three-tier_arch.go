package architecture

import (
	"github.com/ashiqsabith123/initNinja/pkg/config"
	"github.com/ashiqsabith123/initNinja/pkg/helper"
)

func ThreeTier(project_name, language string) {
	folderStructure := config.GetFolderStructure("three_tier_arch")

	helper.CreateFolder(project_name, folderStructure)

	switch language {
	case "GO":
		helper.CreateFiles(project_name, config.GetFolderStructure("go_three_tier_files"))
	}
}
