FROM golang:1.12.9-alpine3.9 AS build
RUN apk update && apk add --no-cache git

WORKDIR /work
ADD ./container .
RUN GOOS=linux GOARCH=amd64 go build -o tex-docker main.go

FROM thomasweise/docker-texlive-full
COPY --from=build /work/tex-docker /usr/local/bin/
RUN chmod +x /usr/local/bin/tex-docker

WORKDIR /home/work
