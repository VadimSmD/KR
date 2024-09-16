FROM 1.23.1-alpine3.20
COPY . repo/
WORKDIR repo
RUN go build -o app ./cmd/app/main.go
RUN ./app > log.txt
