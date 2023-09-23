package ui

import (
	"time"

	"github.com/briandowns/spinner"
)

func Spinner(char_set, sleepDur int, msg, color string, status chan string) {

	s := spinner.New(spinner.CharSets[35], 300*time.Millisecond)
	s.Start()

	s.Color(color)

	s.Suffix = Bold(msg)

	time.Sleep(time.Duration(sleepDur) * time.Second)

	<-status

	s.Stop()

}
