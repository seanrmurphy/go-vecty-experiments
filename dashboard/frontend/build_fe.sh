#! /usr/bin/env bash

env GOOS=js GOARCH=wasm /usr/local/go/bin/go build -o build/out.wasm ./src

