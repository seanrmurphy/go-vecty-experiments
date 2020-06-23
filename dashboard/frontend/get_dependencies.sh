#!/usr/bin/env bash

# Fought wtih go modules for an hour, but this was easier; it was a bit tricky with
# wasm being essentially a different platform
go get github.com/gopherjs/vecty
go get github.com/nathanhack/svg
go get github.com/fatih/structs
#go get "github.com/seanrmurphy/go-echarts@96725735b42bbfd17f6b4e2b630d50c0f9d2e733"
git clone https://github.com/seanrmurphy/go-echarts -b feature/add_js_option_mapping $GOPATH/src/github.com/seanrmurphy/go-echarts/

