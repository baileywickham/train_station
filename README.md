# train_station

![passing?](https://github.com/baileywickham/train_station/workflows/Go/badge.svg)

A go program to simulate a train station. Uses **Github Actions** to run tests, lint, and build binary files.

## Requirements
This repo requires docker to be installed on the host system. golang is recommended but a docker container is also provided.

## Use
Run `./pgstart.sh` to start a docker container containing the database. This will create the database and related tables for the server. Scripts in `./scripts`: `backup.sh` and `restore.sh` are provided to backup and restore testing databases.

This program also runs in a docker container. User `dockerbuild.sh` to start this container. Useful if golang is not installed

Use `go run .` to start the program. `-api` and `-cli` provide repective functionality. There are a couple api urls avaible:
- `/show-all`
- `/create-account?name=&balance=`
- `/server-config`
- `/show?uuid=`

This project uses redis like functionality. This means that the most up to date accounts are stored in memory and once a second they are copied (on another thread) to the database, this is done with build-in gorountines and channels. The actual effectiveness of this is probably less than useless, but it was fun to impliment. 


`config.go` has constants that can be changed.

Use `go test` to manualy run unit tests

## TODO
- add redis functionality [sort of done]
- make object oriented [doneis]
- better error handling [done]
- add database support [done]
- add fileparsing support (runner)
- fix runner to not panic on all arguments [done]
- write more unit tests
- use rpi to do real rfid reading

## Notes/improvements
- golang sql package currently [does not](https://github.com/golang/go/issues/18478) support dynamiclly naming tables.
- Many features of this project are global in scope. This is not good.
- Naming and comments are sparse and not consistent.
- Tests are lacking to say the least.
- Creating an Account relies on the database to create the UUID
- "redis" functionality does not use channels like go recommends. It uses `time.Sleep` instead a `ticker` and could probably be implimented in a way that better takes advantage of go's concurrency patterns. 
