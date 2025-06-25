> More comprehensive README is coming later

## Getting it to run:

1. Define env variables
   - `PIEFED_INSTANCE` - the hostname of the Piefed instance this proxy uses
   - `PORT` (optional)—the port the app runs on, defaults to 8080
   - `SIMULATE` (optional)—whether the Proxy should pretend it's a Lemmy server (when returning version, software etc.)

2. Run
   - `go run main.go`


### Docker

You can build a production ready Docker image from the [Dockerfile](Dockerfile) by running:

`docker build -t lemmy-piefed-proxy .`
