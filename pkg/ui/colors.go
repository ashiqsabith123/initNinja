package ui

import "github.com/fatih/color"

var Magenta = color.New(color.FgHiMagenta).SprintFunc()
var Green = color.New(color.FgGreen).SprintfFunc()
var Blue = color.New(color.FgBlue).SprintFunc()
var Red = color.New(color.FgRed).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()

var ColorArray = []func(a ...interface{}) string{Blue, Yellow, Blue, Red, Yellow, Blue, Yellow}
