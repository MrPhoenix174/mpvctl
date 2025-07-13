package cmd

import (
	"fmt"

	"github.com/MrPhoenix174/mpvctl/config"
)

func ShowAllLinks() {
	links, err := config.LoadLinks("config/links.json")
	if err != nil {
		fmt.Println("Error when reading:", err)
		return
	}
	fmt.Println("All Your Saved Links:")
	if len(links) == 0 {
		fmt.Println("no links")
		return
	}
	for i, link := range links {
		fmt.Printf("[%d] %s - %s\nLinkID -> %s\nLink -> %s\n", i+1, link.LinkName, link.TypeOfEl, link.LinkID, link.UserLink)
	}
}
