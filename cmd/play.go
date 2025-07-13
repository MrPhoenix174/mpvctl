package cmd

import (
	"fmt"
	"os"
	"strings"

	"os/exec"

	"github.com/MrPhoenix174/mpvctl/config"
	//"github.com/MrPhoenix174/mpvctl/cmd"
)

func ShowLinksForPlay(links []config.LinkInfo) {
	fmt.Println("Saved links:")
	if len(links) == 0 {
		fmt.Println("no links")
		return
	}
	for i, link := range links {
		fmt.Printf("[%d] %s - %s\n", i+1, link.LinkName, link.TypeOfEl)
	}
}

func SelectLink(links []config.LinkInfo) (config.LinkInfo, error) {
	ShowLinksForPlay(links)

	if len(links) == 0 {
		return config.LinkInfo{}, fmt.Errorf("no links to choose!")
	}

	fmt.Print("enter link number: ")
	fmt.Print("> ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil || input < 1 || input > len(links) {
		return config.LinkInfo{}, fmt.Errorf("incorrect input")
	}

	return links[input-1], nil
}

func PlayMusicOrVideo() {
	links, err := config.LoadLinks("config/links.json")
	if err != nil {
		fmt.Println("Error when reading:", err)
		return
	}
	fmt.Println("Choose what to play")

	selectedLink, err := SelectLink(links)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("You chose::", selectedLink.LinkName)

	playLink(selectedLink.UserLink, false)

}

func playLink(link string, isVideo bool) {
	var cmd *exec.Cmd
	if isVideo {
		fmt.Println("Do you want disable audio?(y/N)")
		var withAudioStr string
		fmt.Print("> ")
		fmt.Scan(&withAudioStr)
		if strings.HasPrefix((strings.ToLower(withAudioStr)), "y") {
			cmd = exec.Command("mpv", "--no-audio", link)
		} else {
			cmd = exec.Command("mpv", link)
		}
	} else {
		cmd = exec.Command("mpv", "--no-video", link)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error at start MPV:", err)
	}

	fmt.Println("MPV worked without errors.")

}
