package config

import (
	"encoding/json"
	//"io/ioutil"
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
