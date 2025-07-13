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
			// view all links
		case 2:
			cmd.AddLink()
		case 3:
			// play music
		case 4:
			// show video
		case 5:
			// launch cava
		case 6:
			// launch EasyEffects
		case 7:
			// settings
		case 0:
			fmt.Println("bye")
			os.Exit(0)
		default:
			fmt.Println("Unknown command!")
		}

	}
}
