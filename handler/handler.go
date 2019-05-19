package handler

import (
	"context"

	pb "github.com/microhq/spotify-srv/proto/spotify"

	"github.com/microhq/spotify-srv/spotify"

	"github.com/micro/go-log"
)

// Spotify implements spotify microsevice handler.
// It communicates with Spotify API to play songs.
type Spotify struct {
	// Client is Spotify API client
	Client *spotify.Client
}

// Play plays a song on a Spotify device
func (s *Spotify) Play(ctx context.Context, req *pb.PlayRequest, resp *pb.PlayResponse) error {
	songURI := req.Name

	log.Logf("Attempting to play song: %s on Spotify device: %s", req.Name, s.Client.Device().ID)

	return s.Client.PlaySong(songURI)
}

// Stop stops Spotify playback
func (s *Spotify) Stop(ctx context.Context, req *pb.StopRequest, resp *pb.StopResponse) error {

	log.Logf("Attempting to stop Spotify playback")

	return s.Client.Stop()
}
