package architecture

import (
	"fmt"
	"os/exec"
)

func CleanCode() {

	cmd := exec.Command("mkdir", "hhee")

	cmd.Dir = "/"

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the output of the command
	fmt.Println("Command Output:")
	//fmt.Println(string(output))
}
