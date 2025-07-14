package main

import (
	"fmt"
	"os"

	"github.com/MrPhoenix174/mpvctl/cmd"
	"github.com/MrPhoenix174/mpvctl/internal"
	"github.com/MrPhoenix174/mpvctl/ui"
)

func main() {
	//defer fmt.Println("bye")

	for {
		ui.ShowMainMenu()
		input := cmd.GetUserChoice()
		switch input {
		case 1:
			cmd.ShowAllLinks()
		case 2:
			cmd.AddLink()
		case 3:
			cmd.PlayMusicOrVideo()
		case 4:
			//internal.LaunchCava() /// -now not works (or no?)
		case 5:
			internal.LaunchEasyEffects()
		case 6:

		case 0:
			fmt.Println("bye")
			os.Exit(0)
		default:
			fmt.Println("Unknown command!")
		}

	}
}
