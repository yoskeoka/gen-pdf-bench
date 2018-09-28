#!/bin/bash

(cd puppeteer && yarn install)

go get -v ./...
go build -o testwk ./cmd/pdf
