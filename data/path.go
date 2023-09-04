package data

import (
	"path/filepath"

	"furina/config"
)

func getCharacterNameMapFile() string {
	return filepath.Join(config.GetBaseDir(), "config", "characterName.json")
}

func getCharacterPropertyWeightMapFile() string {
	return filepath.Join(config.GetBaseDir(), "config", "propertyWeight.json")
}

func getSetNameTextHashMapFile() string {
	return filepath.Join(config.GetBaseDir(), "config", "setNameTextHash.json")
}

func getCharacterDataFileByName(name string) string {
	return filepath.Join(config.GetBaseDir(), "static", "resources", "character", name, "data.json")
}

func getWeaponDataFileByType(weapon string) string {
	return filepath.Join(config.GetBaseDir(), "static", "resources", "weapon", weapon, "data.json")
}

func getArtifactDataFile() string {
	return filepath.Join(config.GetBaseDir(), "static", "resources", "artifact", "data.json")
}

func getUserListFile() string {
	return filepath.Join(config.GetBaseDir(), "storage", "user.json")
}
