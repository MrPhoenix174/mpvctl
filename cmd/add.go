package cmd

import (
	//"github.com/MrPhoenix174/mpvctl/cmd"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MrPhoenix174/mpvctl/config"
)

func getNewLink() (config.LinkInfo, error) {
	fmt.Println("Hi, pls write new link")
	fmt.Print("> ")
	var newLink config.LinkInfo
	fmt.Scan(&newLink.UserLink)
	if s := strings.TrimSpace(newLink.UserLink); s == "" {
		return newLink, fmt.Errorf("error: link cannot be empty!")
	}
	fmt.Println("Okay, let's add a name")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	newLink.LinkName = scanner.Text()
	if l := strings.TrimSpace(newLink.LinkName); l == "" {
		fmt.Println("name is empty, generate random")
		newLink.LinkName = generateRandomString(6)
	}
	newLink.LinkID = generateRandomString(10)
	newLink.TypeOfEl = "video"
	fmt.Println("And lastly, will it be music or video?(m,V)")
	fmt.Print("> ")
	var morv string
	fmt.Scan(&morv)
	morv = strings.ToLower(strings.TrimSpace(morv))
	if strings.HasPrefix(morv, "m") {
		newLink.TypeOfEl = "music"
	}
	fmt.Println("Link added, trying to save")

	return newLink, nil
}

func AddLink() {
	link, err := getNewLink()
	if err != nil {
		fmt.Println("Try again!")
	}
	const filePath = "config/links.json"

	links, err := config.LoadLinks(filePath)
	if err != nil {
		fmt.Println("Error when reading:", err)
		return
	}
	links = append(links, link)
	err = config.SaveLinks(filePath, links)
	if err != nil {
		fmt.Println("error when saving", err)
		return
	}

	fmt.Println("Link Saved!")

}
