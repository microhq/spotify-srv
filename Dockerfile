FROM alpine:3.2
ADD spotify-srv /spotify-srv
ENTRYPOINT [ "/spotify-srv" ]
