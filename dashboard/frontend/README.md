# A simple vecty based frontend which shows basic charts

This repo contains a simple frontend which show two ways of rendering charts
with vecty: via simple SVGs or via Apache E-Charts.

# Installation

## Pre-requisites

- modern version of go - this has been tested with go 1.14.2
- modern version of python - this has been tested with python 3.8.2

## Installation Process

Installation is not as smooth as it could be as it is not quite a standard Go
project.

- Clone the repo
    - `git clone https://github.com/seanrmurphy/go-vecty-experiments $GOPATH/src/github.com/seanrmurphy/go-vecty-experiments`
- Install dependencies
    - `cd $GOPATH/src/github.com/seanrmurphy/go-vecty-experiments`
    - `./install_dependencies.sh`


# Running the Application

- Build the frontend
    - `./build_fe.sh`
- Run the server
    - `./server.py`
    - This should run the server on port 4443 on all interfaces on the host
- Point a browser at `<hostname/hostip>:4443/`
    - You will probably have to tell your browser you trust the self signed certificate

# Video

