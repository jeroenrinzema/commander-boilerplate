# Connector

The connector exposes a REST api in order to execute commands en await (if requested) their responding events.
A timeout is thrown on sync requests if no responding event is created within the timeout period (default: 5s).
As an example is it also plausible to consume/await events via the websocket connection.

## Environment variables

This connector requires the following environment variables to be set.

```bash
export KAFKA_GROUP=command
export KAFKA_BROKERS=192.168.99.100:9092
```
