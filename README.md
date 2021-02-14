# Spotify Backup

A service for backing up your **listening history**, **liked songs**, and **playlists**.

---

**NOTE**: many (if not all) features are in-progress right now. Until [v0.1](https://github.com/cbochs/spotify-backup/projects/1) is complete, most of these features are likely to be unavailable.

**Features**
* Full-featured CLI to manage and run backups
* Store your listening history, liked songs, and playlists to a database
* Store incremental snapshots of liked songs or playlists
* Store track details, track features, and track analysis
* Schedule individual cron-like backups
* Export your data to JSON (or, nd-JSON)

## Self-Host

* Build from source with `build/build.sh`
* Start mongo with `docker-compose -f build/docker-compose.yml -d up`
* Run backup with `bin/sp`