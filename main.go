/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ashiqsabith123/initNinja/cmd"
	"github.com/ashiqsabith123/initNinja/pkg/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
