package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func LaunchCava() {
	cmd := exec.Command("/bin/bash", "internal/runner.sh") // i'll fix it
	/*
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Cava startup error: ", err)
			return
		}
	*/
	output, err := cmd.CombinedOutput()
	if err != nil {
		//return fmt.Errorf("mpv error: %v\nOutput: %s", err, string(output))
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("cava startup succesful!")
	fmt.Println(string(output))

}

func LaunchEasyEffects() {
	cmd := exec.Command("easyeffects")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("easyeffects startup error: ", err)
		return
	}
	fmt.Println("easyeffects startup succesful!")

}
