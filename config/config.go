package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	BaseDir  string
	Server   ServerConfig
	Database DatabaseConfig
	Version  VersionConfig
}

type ServerConfig struct {
	RunMode  string
	HttpPort string
}

type DatabaseConfig struct {
	LocalStorageDir string
	Username        string
	Password        string
	Addr            string
	DBName          string
}

type VersionConfig struct {
	Version     string
	GameVersion string
	DataSource  string
}

var (
	gConfig Config
)

func init() {
	if dir, err := os.Getwd(); err != nil {
		log.Fatalln(err)
	} else {
		gConfig.BaseDir = dir
	}
	vp := viper.New()
	vp.AddConfigPath("config/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	if err := vp.UnmarshalKey("Server", &gConfig.Server); err != nil {
		log.Fatalln(err)
	}
	if gConfig.Database.LocalStorageDir == "" {
		gConfig.Database.LocalStorageDir = filepath.Join(gConfig.BaseDir, "local")
	}
	if err := vp.UnmarshalKey("Database", &gConfig.Database); err != nil {
		log.Fatalln(err)
	}
	gConfig.Version = VersionConfig{
		Version:     "1.0",
		GameVersion: "4.0",
		DataSource:  "Enka",
	}
}

func GetConfig() Config {
	return gConfig
}

func GetBaseDir() string {
	return gConfig.BaseDir
}

func GetLocalStorageDir() string {
	return gConfig.Database.LocalStorageDir
}
