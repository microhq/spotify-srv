package main

import (
	"os"

	"github.com/microhq/spotify-srv/handler"
	"github.com/microhq/spotify-srv/spotify"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	pb "github.com/microhq/spotify-srv/proto/spotify"
)

var (
	redirectURI = "http://localhost:8080/callback"
)

func main() {
	// New Service
	srv := micro.NewService(
		micro.Name("go.micro.srv.spotify"),
		micro.Version("latest"),
	)

	// Initialise service
	srv.Init()

	// read environment variables
	spotifyID := os.Getenv("SPOTIFY_ID")
	if spotifyID == "" {
		log.Fatal("Could not read SPOTIFY_ID environment variable")
	}

	spotifySecret := os.Getenv("SPOTIFY_SECRET")
	if spotifySecret == "" {
		log.Fatal("Could not read SPOTIFY_SECRET environment variable")
	}

	spotifyDeviceID := os.Getenv("SPOTIFY_DEVICE_ID")
	if spotifySecret == "" {
		log.Fatal("Could not read SPOTIFY_DEVICE_ID environment variable")
	}

	_redirectURI := os.Getenv("REDIRECT_URI")
	if _redirectURI != "" {
		redirectURI = _redirectURI
	}

	client, err := spotify.NewClient(spotifyID, spotifySecret, redirectURI, spotifyDeviceID)
	if err != nil {
		log.Fatal("Could not create Spotify client")
	}

	// Register Handler
	pb.RegisterSpotifyHandler(srv.Server(), &handler.Spotify{client})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
