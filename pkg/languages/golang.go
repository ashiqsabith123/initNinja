package languages

import (
	"fmt"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

func Golang(project_name string) {

	module_name := ""

	qs := survey.Input{
		Message: "Enter you mod name: ",
	}

	survey.AskOne(&qs, &module_name)


	cmd := exec.Command("go", "mod", "init", module_name)
	cmd.Dir = project_name + "/"

	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

}
