# train_station

A go program to simulate a train station. Uses Github Actions to run tests, lint, and build binary files.

## Use
This program uses postgres as the database. This can be installed for tesing using the dockerfile


Use `go run .` to start the cli runner


Use `go test` to manualy run unit tests

## TODO
- make object oriented
- better error handling
- add database support
- add fileparsing support (runner)
- fix runner to not panic on all arguments
- write more unit tests
- use rpi to do real rfid reading
