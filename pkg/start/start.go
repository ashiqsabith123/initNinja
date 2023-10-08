package start

import (
	"fmt"

	"github.com/ashiqsabith123/initNinja/pkg/architecture"
	core "github.com/ashiqsabith123/initNinja/pkg/core"
	"github.com/ashiqsabith123/initNinja/pkg/frameworks"
	"github.com/ashiqsabith123/initNinja/pkg/languages"
	ui "github.com/ashiqsabith123/initNinja/pkg/ui"
)

// type Starter struct {
// }

// func NewStarter() Starter {
// 	return Starter{}
// }

func Start() {

	fmt.Println(ui.Magenta(ui.Banner))
	fmt.Println(ui.Bold("Hola... Welcome to initNinja \n"))

	output := core.DisplayAndSelectInputs()

	switch output.Architecture {
	case "Clean Code Architecture":
		architecture.CleanCode(output.Project_name, output.Language)
	case "MVC":
		architecture.Mvc(output.Project_name, output.Language)
	case "DDD":
		architecture.DDD(output.Project_name, output.Language)
	case "three-tier":
		architecture.ThreeTier(output.Project_name, output.Language)
	}

	switch output.Language {
	case "GO":
		languages.Golang(output.Project_name)
	}

	switch output.Framework {
	case "Gin":
		frameworks.Gin(output.Project_name, output.Architecture)
	}

}
