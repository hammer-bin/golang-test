package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func output1() {
	targetDir := "C:/putty"
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--------------------------")
	for _, file := range files {
		fmt.Println(file.Name())
		fmt.Println(fmt.Sprintf("%v/%v", targetDir, file.Name()))
	}
}

func main() {

	cmd := exec.Command("ls", "-a")
	cmd.Dir = "C:/putty"

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	output1()
}
