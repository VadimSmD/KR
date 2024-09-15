FROM ubuntu:22.04
COPY . repo/
WORKDIR repo
RUN apt update
RUN apt install -yy golang-go
RUN go run cmd/app/main.go