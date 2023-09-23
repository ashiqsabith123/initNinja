package core

import (
	"errors"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/acarl005/stripansi"
	"github.com/ashiqsabith123/initNinja/pkg/architecture"
	"github.com/ashiqsabith123/initNinja/pkg/frameworks"
	"github.com/ashiqsabith123/initNinja/pkg/languages"
	ui "github.com/ashiqsabith123/initNinja/pkg/ui"
)

type Details struct {
	Project_name string
	Language     string
	Framework    string
	Architecture string
}

func DisplayAndSelectInputs() Details {

	details := Details{}

	language := []string{}
	for i := 0; i < len(languages.Languages); i++ {
		language = append(language, (ui.ColorArray[i](languages.Languages[i])))
	}

	architectures := []string{}
	architectures = append(architectures, architecture.Architectures...)

	qs1 := []*survey.Question{

		{
			Name:   "project_name",
			Prompt: &survey.Input{Message: "Enter your project name: "},
			Validate: func(ans interface{}) error {
				if _, err := os.Stat(ans.(string)); os.IsNotExist(err) {
					return nil
				}
				return errors.New(ui.Bold(ui.Yellow(ans.(string)) + ui.Red(" -named folder already exist in this directory")))
			},
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

	err := survey.Ask(qs1, &details)
	if err != nil {
		log.Fatal(err)
	}

	var framework []string

	details.Language = stripansi.Strip(details.Language)

	for i, v := range frameworks.Frameworks {
		if i == details.Language {
			framework = append(framework, v...)
			break
		}
	}

	qs2 := []*survey.Question{
		{
			Name: "framework",
			Prompt: &survey.Select{
				Message: "Select your framework",
				Options: framework,
			},
		},
		{
			Name: "architecture",
			Prompt: &survey.Select{
				Message: "Select your architeture",
				Options: architectures,
			},
		},
	}

	// Ask the user to select an option
	if err := survey.Ask(qs2, &details); err != nil {
		log.Fatal(err)
	}

	return details

}
