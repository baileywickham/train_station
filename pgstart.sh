#!/bin/bash
docker build -t pg -f Dockerfile_pg . && docker run -d -p 5432:5432 pg
