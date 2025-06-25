FROM golang:1.24 AS build

WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o app .

FROM ubuntu:24.04
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/* && apt-get clean
COPY --from=build /app/app /app

ENTRYPOINT ["/app"]
