FROM golang:alpine
RUN apk add --no-cache postgresql postgresql-client postgresql-dev
COPY . repo/
WORKDIR repo
ARG pg_url
ENV PG_URL=$pg_url
RUN go mod download
EXPOSE 8080

