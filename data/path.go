package data

import (
	"path/filepath"

	"furina/config"
)

func getCharacterNameMapFile() string {
	return filepath.Join(config.GetBaseDir(), "data", "characterName.json")
}

func getWeaponNameMapFile() string {
	return filepath.Join(config.GetBaseDir(), "data", "weaponName.json")
}

func getArtifactSetNameMapFile() string {
	return filepath.Join(config.GetBaseDir(), "data", "artifactSetName.json")
}

func getCharacterPropertyWeightMapFile() string {
	return filepath.Join(config.GetBaseDir(), "config", "propertyWeight.json")
}

func getCharacterDataFileByName(name string) string {
	return filepath.Join(config.GetBaseDir(), "static", "resources", "character", name, "data.json")
}

func getWeaponDataFileByTypeAndName(weaponType, name string) string {
	return filepath.Join(config.GetBaseDir(), "static", "resources", "weapon", weaponType, name, "data.json")
}

func getArtifactDataFile() string {
	return filepath.Join(config.GetBaseDir(), "static", "resources", "artifact", "data.json")
}
