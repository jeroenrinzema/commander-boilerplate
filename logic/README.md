# Logic

The logic service consumes commands which should always result in a resulting event.

## Getting started

To run the service manually for example during development.
Make sure that you have the `environment` variables set and have `kafka` and `postgres` running.

```bash
$ # cd docker/dependencies
$ docker-compose up -d kafka postgres
$ # cd logic/
$ go run main.go
```

## Environment variables

These are the required environment variables.

```bash
export KAFKA_GROUP=logic
export KAFKA_BROKERS=192.168.99.100:9092
export POSTGRES_HOST=192.168.99.100
export POSTGRES_PORT=5432
export POSTGRES_USER=commander
export POSTGRES_PASSWORD=TFgvT3Pb3bWEhXKAfgMk63bo
export POSTGRES_DB=commander
```
