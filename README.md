# Spotify Service

This is an example of a simple Spotify service

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)
- [Example](#example)

## Configuration

- FQDN: `go.micro.srv.spotify`
- Type: `srv`
- Alias: `spotify`

In order to use this service you ned to register Spotify application to obtain your client ID and API secret. You can follow the the guid [here](https://developer.spotify.com/documentation/general/guides/app-settings/).

You need to export the following environment variables:
- `SPOTIFY_ID`: Spotify client ID
- `SPOTIFY_SECRET`: Spotify API secret
- `SPOTIFY_DEVICE_ID`: Spotify device ID you want to play the song on

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Build a docker image
```
make docker
```

Run the service
```
./spotify-srv
```

Alternatively you can run the service locally as follows:
```
go run github.com/microhq/spotify-srv
```

Finally run it inside Docker container:
```
docker run -P -e SPOTIFY_ID="XXX" -e SPOTIFY_SECRET="XXX" -e SPOTIFY_DEVICE_ID="XXX" -e REDIRECT_URI="XXX" microhq/spotify-srv
```

**NOTE:** when running inside Docker container the service runs in the container network namespace and that includes DNS! You will only be able to trigger the playback if you call the service from inside that network namespace.

## Example

Start the service:
```
$ go run github.com/microhq/spotify-srv
Log in to Spotify by visiting the following URL in your browser: https://accounts.spotify.com/authorize?client_id=XXXX&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback&response_type=code&scope=user-read-currently-playing+user-read-playback-state+user-modify-playback-state&state=XXX
2019/05/19 17:27:01 Transport [http] Listening on [::]:53000
2019/05/19 17:27:01 Broker [http] Connected to [::]:53001
2019/05/19 17:27:01 Registry [mdns] Registering node: go.micro.srv.spotify-e920b766-8faf-4df8-b494-ee1161764e0b
```

Play a song:
```
$ micro call go.micro.srv.spotify Spotify.Play '{"name": "spotify:track:6bS7DqTuzrpkluB7boWmw2"}'
```

Stop the playback:
```
$ micro call go.micro.srv.spotify Spotify.Stop '{}'
```
