package config

import (
	"encoding/json"
	"fmt"

	//"fmt"
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

func LoadSettings(path string) (Settings, error) {
	var settings Settings
	data, err := os.ReadFile(path)
	if err != nil {
		return settings, fmt.Errorf("failed to read settings file: %w", err)
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return settings, fmt.Errorf("failed to parse settings JSON: %w", err)
	}
	return settings, nil
}
