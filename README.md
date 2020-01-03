# train_station

![passing?](https://github.com/baileywickham/train_station/workflows/Go/badge.svg)

A go program to simulate a train station. Uses **Github Actions** to run tests, lint, and build binary files.

## Requirements
This repo requires docker to be installed on the host system. Golang is recommended but a docker container is also provided.

## Use
Run `./pgstart.sh` to start a docker container containing the database. This will create the database and related tables for the server. Scripts in `./scripts`: `backup.sh` and `restore.sh` are provided to backup and restore testing databases.

This program also runs in a docker container. User `dockerbuild.sh` to start this container. Useful if golang is not installed

Use `go run .` to start the cli runner. `-api` and `-cli` provide repective functionality. There are a couple api urls avaible:
- /show-all
- /create-account?name=&balance=
- /server-config
- /show?uuid=


Use `go test` to manualy run unit tests

## TODO
- add redis functionality
- make object oriented [doneis]
- better error handling [done]
- add database support [done]
- add fileparsing support (runner)
- fix runner to not panic on all arguments [done]
- write more unit tests
- use rpi to do real rfid reading
