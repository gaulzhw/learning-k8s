package main

import (
	"log"
)

var (
	BuildVersion = "--"
	BuildHash    = "--"
	BuildTime    = "--"
)

/*
how to use:
add these commands to Makefile, build with ldflags
	GIT_VERSION := $(shell git describe --match "v[0-9]*")
	GIT_BRANCH := $(shell git branch | grep \* | cut -d ' ' -f2)
	GIT_HASH := $(GIT_BRANCH)/$(shell git log -1 --pretty=format:"%H")
	TIMESTAMP := $(shell date '+%Y.%m.%d-%I:%M:%S')
	LD_FLAGS="-X main.BuildVersion=$(GIT_VERSION) -X main.BuildHash=$(GIT_HASH) -X main.BuildTime=$(TIMESTAMP)"
	go build -ldflags=$(LD_FLAGS) main.go
*/
func main() {
	log.Printf("build version: %s\n", BuildVersion)
	log.Printf("build hash: %s\n", BuildHash)
	log.Printf("build time: %s\n", BuildTime)
}
