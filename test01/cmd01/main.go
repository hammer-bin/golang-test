package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"syscall"
)

func execCommandOutput(name string, arg ...string) (string, error) {
	//os.Chdir("..")
	cmd := exec.Command(name, arg...)

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // 콘솔모드 프로그램실행 시 검은 콘솔창 뜨는것 제거(윈도우만)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
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

func run() {

	cmd := exec.Command("ls", "-a")
	cmd.Dir = "C:/putty"

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}

func main() {
	a, err := execCommandOutput("ls", "-a")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	fmt.Println("----------------------------------------------")
	//output1()
}
