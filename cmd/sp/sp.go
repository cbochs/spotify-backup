package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/cbochs/spotify-backup-api/backup"
	"github.com/cbochs/spotify-backup-api/backupdb"
	"github.com/cbochs/spotify-backup-api/config"
	"github.com/cbochs/spotify-backup-api/spotify/auth"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// PARSING CLI FLAGS //

	configPath := flag.String("config", "", "Path to configuration file.")
	enableDebug := flag.Bool("debug", false, "Enable debugging output.")

	flag.Parse()

	// SETTING UP LOGGING //

	logLevel := zerolog.InfoLevel
	if *enableDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	log.Info().Msg("Starting Spotify Backup API")

	// CREATING CONFIG //

	if *configPath == "" {
		log.Error().Msg("Must provide a configuration file path")
		return
	}

	config, err := config.FromFile(*configPath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load configuration file")
		return
	}

	// CONNECTING DATABASE //

	log.Info().Msg("Connecting to Database...")

	ctx := context.Background()

	db, err := backupdb.New(ctx, &config.DB)
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}
	if err := db.Connect(); err != nil {
		log.Error().Err(err).Msg("")
		return
	}

	log.Info().Msg("Spotify Backup API started")

	spa := auth.New().
		WithCredentials(config.Auth.ClientID, config.Auth.ClientSecret).
		WithRedirect(config.Auth.RedirectURL).
		WithScopes(config.Auth.Scopes...)

	service := backup.NewService(ctx, db, spa)

	client, err := service.ClientFromSpotifyID("notbobbobby")
	if err != nil {
		log.Error().Err(err).Msg("")
		fmt.Printf("Auth URL: %s\n\n", service.AuthCodeURL(""))
		fmt.Println("Ender code: ")

		var code string
		fmt.Scanln(&code)

		client, err = service.Exchange(code)
		if err != nil {
			log.Error().Err(err).Msg("")
			return
		}
	}

	if err := client.BackupRecentlyPlayed(ctx); err != nil {
		log.Error().Err(err).Msg("")
		return
	}
}
