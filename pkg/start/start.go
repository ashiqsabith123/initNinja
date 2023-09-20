package start

import (
	"fmt"

	"github.com/ashiqsabith123/initNinja/pkg/architecture"
	core "github.com/ashiqsabith123/initNinja/pkg/core"
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
		architecture.CleanCode(output.Project_name)
	}

}
