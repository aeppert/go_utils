#!/bin/bash

export PATH="/opt/go/go/bin:$PATH"
export GOROOT="/opt/go/go"
export GOPATH="/opt/go"

go get github.com/bluele/slack
go build slackmsg.go
