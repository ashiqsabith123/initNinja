package core

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ashiqsabith123/initNinja/pkg/frameworks"
	"github.com/ashiqsabith123/initNinja/pkg/languages"
	ui "github.com/ashiqsabith123/initNinja/pkg/ui"
)

type Details struct {
	Project_name string
	Language     string
	Framework    string
}

func DisplayAndSelectInputs() {

	details := Details{}

	language := []string{}
	for i := 0; i < len(languages.Languages); i++ {
		language = append(language, ui.Bold(ui.ColorArray[i](languages.Languages[i])))
	}

	qs := []*survey.Question{

		{
			Name:      "project_name",
			Prompt:    &survey.Input{Message: "Enter your project name: "},
			Validate:  survey.Required,
			Transform: survey.Title,
		},

		{
			Name: "language",
			Prompt: &survey.Select{
				Message: "Select your programming language",
				Options: language,
			},
		},
	}

	err := survey.Ask(qs, &details)
	if err != nil {
		log.Fatal(err)
	}

	var framewrk []string

	fmt.Println(details)
	details.Language += "\033[0m"
	fmt.Println(details)

	for i, v := range frameworks.Frameworks {
		if i == details.Language {
			fmt.Println("hh")
			fmt.Println(v)
		}
	}

	prompt := &survey.Select{
		Message: "Select an option:",
		Options: framewrk,
	}

	// Ask the user to select an option
	if err := survey.AskOne(prompt, &details.Framework); err != nil {
		log.Fatal(err)
	}

}
