package main

import (
	"fmt"
	"os"

	"github.com/MrPhoenix174/mpvctl/cmd"
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
			// launch cava
		case 5:
			// launch EasyEffects
		case 6:
			// settings

		case 0:
			fmt.Println("bye")
			os.Exit(0)
		default:
			fmt.Println("Unknown command!")
		}

	}
}
