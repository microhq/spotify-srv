package spotify

import (
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/micro/go-log"
	"github.com/zmb3/spotify"
)

// auth allows to authenticate with Spotify API
type auth struct {
	// Spotify authenticator
	*spotify.Authenticator
	// state is OAuth state
	state string
	// redirectURI is Spotify RedirectURI
	redirectURI string
}

// newAuth returns auth which is used to authenticate with Spotify API
func newAuth(clientID, clientSecret, redirectURI, state string) *auth {
	// Spotify API authenticator
	a := spotify.NewAuthenticator(redirectURI,
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserModifyPlaybackState)
	a.SetAuthInfo(clientID, clientSecret)

	return &auth{&a, state, redirectURI}
}

// URL returns Spotify API OAuth URL
func (a *auth) URL() string {
	// Spotify Login URL
	return a.AuthURL(a.state)
}

// authHandler handles OAuth2 authentication callback from Spotify API
func authHandler(a *auth, ch chan *spotify.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tok, err := a.Token(a.state, r)
		if err != nil {
			http.Error(w, "Couldn't retrieve OAuth token", http.StatusForbidden)
			// TODO: handle the errors better
			log.Fatalf("Spotify auth error: %s", err)
		}
		if s := r.FormValue("state"); s != a.state {
			http.NotFound(w, r)
			log.Fatalf("OAuth State mismatch: %s != %s\n", s, a.state)
		}
		// use the token to get an authenticated client
		client := a.NewClient(tok)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "Spotify login successful!")
		ch <- &client
	})
}

// Client is Spotify API client
type Client struct {
	// Spotify client
	*spotify.Client
	// device is Spotify player device
	device *spotify.PlayerDevice
	// mutex
	*sync.Mutex
}

// NewClient authenticates with Spotify API and returns Spotify API client.
// It returns error if API authentication fails or if the deviceID could not be found.
func NewClient(clientID, clientSecret, redirectURI, deviceID string) (*Client, error) {
	// Spotify client auth channels
	clientChan := make(chan *spotify.Client)
	// Spotify authenticator
	auth := newAuth(clientID, clientSecret, redirectURI, "abc123")
	// Create HTTP muxer
	h := http.NewServeMux()
	h.Handle("/callback", authHandler(auth, clientChan))
	// HTTP server for Spotify OAuth
	server := &http.Server{
		Handler: h,
	}

	errChan := make(chan error, 1)
	// create OAuth listener for RedirectURI callback
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return nil, fmt.Errorf("Failed to create TCP listener: %s", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		errChan <- server.Serve(listener)
	}()

	fmt.Println("Log in to Spotify by visiting the following URL in your browser:", auth.URL())

	var client *spotify.Client
	// wait for auth to complete
	select {
	case client = <-clientChan:
		listener.Close()
	case err = <-errChan:
	}
	wg.Wait()

	if err != nil {
		return nil, fmt.Errorf("Failed to create Spotify client: %s", err)
	}

	// configure Spotify player device
	_deviceID := spotify.ID(deviceID)
	device := &spotify.PlayerDevice{ID: _deviceID}
	_client := &Client{client, device, &sync.Mutex{}}

	return _client, err
}

// Device returns Spotify active playback device
func (c *Client) Device() *spotify.PlayerDevice {
	return c.device
}

// PlaySong plays songURI Spotify song
func (c *Client) PlaySong(songURI string) error {
	c.Lock()
	defer c.Unlock()

	// if empty, return error
	if songURI == "" {
		return fmt.Errorf("Error: empty song URI")
	}

	// set playback options: device ID and song URI
	opts := &spotify.PlayOptions{
		DeviceID: &c.device.ID,
		URIs:     []spotify.URI{spotify.URI(songURI)},
	}

	return c.PlayOpt(opts)
}

// Stop pauses active playback on the currently active Spotify device.
// It returns error if the playback can't be paused
func (c *Client) Stop() error {
	opts := &spotify.PlayOptions{
		DeviceID: &c.device.ID,
	}

	return c.PauseOpt(opts)
}
