package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Version  VersionConfig
}

type ServerConfig struct {
	BaseDir         string
	RunMode         string
	LocalStorageDir string
	HttpPort        string
}

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	DBName   string
}

type VersionConfig struct {
	Version     string
	GameVersion string
	DataSource  string
}

var (
	gConfig = Config{}
)

func Load() {
	if dir, err := os.Getwd(); err != nil {
		log.Fatalln(err)
	} else {
		gConfig.Server.BaseDir = dir
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
	if gConfig.Server.LocalStorageDir == "" {
		gConfig.Server.LocalStorageDir = filepath.Join(gConfig.Server.BaseDir, "local")
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
	return gConfig.Server.BaseDir
}

func GetLocalStorageDir() string {
	return gConfig.Server.LocalStorageDir
}
