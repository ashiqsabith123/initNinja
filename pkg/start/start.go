package start

import (
	"fmt"

	core "github.com/ashiqsabith123/initNinja/pkg/core"
	ui "github.com/ashiqsabith123/initNinja/pkg/ui"
)

type Details struct {
	FolderName   string
	Language     string
	Architecture string
}

// type Starter struct {
// }

// func NewStarter() Starter {
// 	return Starter{}
// }

func Start() {

	fmt.Println(ui.Magenta(ui.Banner))

	fmt.Println(ui.Bold("Hola... Welcome to initNinja \n"))

	fmt.Println("")

	core.DisplayAndSelectInputs()

}
