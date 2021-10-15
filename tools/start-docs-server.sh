#!/usr/bin/env sh
# TODO write a script for stopping the server for now docker ps, docker kill, and docker container rm will help you
#      clean stuff up.

docker run --name hived-docs-server -d -p 8080:80 theshadow/hived:latest