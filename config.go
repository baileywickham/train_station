package main

const (
	// Size of queue of changed accounts to be written to db.
	BUFSIZE = 10
	// Connection string for database
	CONNSTR = "user=postgres password=postgres dbname=station sslmode=disable"
	// Time between writes to database, measured in milisecond.
	DBWRITE = 50
)
