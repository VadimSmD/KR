FROM golang:alpine
COPY . repo/
WORKDIR repo
RUN go build -o app ./cmd/app/main.go
RUN ./app > log.txt
