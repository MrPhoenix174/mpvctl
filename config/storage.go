package config

import (
	"encoding/json"

	"fmt"
	"os"
	//"github.com/MrPhoenix174/mpvctl/cmd"
	//"github.com/MrPhoenix174/mpvctl/cmd"
)

type LinkInfo struct {
	UserLink string
	LinkName string
	LinkID   string
	TypeOfEl string
}

type Settings struct {
	Autoplay         bool `json:"autoplay"`
	DefaultVolume    int  `json:"defaultVolume"`
	StartEasyEffects bool `json:"startEasyEffects"`
	StartCava        bool `json:"startCava"`
}

func LoadLinks(path string) ([]LinkInfo, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var links []LinkInfo
	err = json.Unmarshal(file, &links)
	return links, err
}

func SaveLinks(filename string, links []LinkInfo) error {
	data, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func ShowLinksForPlay(links []LinkInfo) {
	fmt.Println("Saved links:")
	if len(links) == 0 {
		fmt.Println("no links")
		return
	}
	for i, link := range links {
		fmt.Printf("[%d] %s - %s\n", i+1, link.LinkName, link.TypeOfEl)
	}
}
