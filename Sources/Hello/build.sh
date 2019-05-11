#!/usr/bin/env bash
go build -buildmode=c-shared -o libhello.so hello.go
