package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
)

type Config struct {
	Service    ServiceConfig  `json:"service"`
	SaveFormat SaveFormat     `json:"save_format"`
	DB         DatabaseConfig `json:"db"`
	Auth       oauth2.Config  `json:"auth"`
}

func FromFile(filePath string) (*Config, error) {
	fh, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	byt, err := ioutil.ReadAll(fh)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(byt, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type ServiceConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (c *ServiceConfig) URL() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type SaveFormat struct {
	OmitDisplayName bool `json:"omit_display_name"`
	SaveAlbums      bool `json:"save_albums"`
	SaveArtists     bool `json:"save_artists"`
	SaveTracks      bool `json:"save_tracks"`
	TrackAnalysis   bool `json:"track_analysis"`
	TrackFeatures   bool `json:"track_features"`
	TrackHistory    bool `json:"track_history"`
	PlaylistHistory bool `json:"playlist_history"`
}

type DatabaseConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Name        string `json:"name"`
	Username    string `json:"user"`
	Password    string `json:"pswd"`
	ConnTimeout int    `json:"conn_timeout"`
}

func (c *DatabaseConfig) Options() *options.ClientOptions {
	return options.
		Client().
		ApplyURI(c.uri()).
		SetAuth(options.Credential{
			Username: c.Username,
			Password: c.Password,
		})
}

func (c *DatabaseConfig) uri() string {
	return fmt.Sprintf("mongodb://%s:%d", c.Host, c.Port)
}
