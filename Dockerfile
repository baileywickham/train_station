FROM golang

RUN mkdir /train_station
COPY . /train_station
WORKDIR /train_station

cmd go run . -api -cli
