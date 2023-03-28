![](img/logo.png)

# Quick Start

This chapter shows you how to get started with PuzzleDB quickly: you can start a standalone PuzzleDB server with Docker and use Redis, MongoDB or MySQL CLI commands to insert and read sample data.

## Start PuzzleDB Server

PuzzleDB Docker image is the easiest way; if you do not have Docker installed, go there and install it first. To start the standalone server, use the following command:

```
docker run -d --name puzzledb-server \
  -p 6379:6379 \
  -p 27017:27017 \
  -p 3307:3307 \
  puzzledb/puzzledb-server
```

PuzzleDB listens on three default database ports: Redis, MongoDB, and MySQL.

## Redis

To operate PuzzleDB with the Redis protocol, use the standard Redis command [redis-cli](https://redis.io/docs/ui/cli/) as follows:

```
%  redis-cli 
127.0.0.1:6379> SET mykey "Hello"
OK
127.0.0.1:6379> GET mykey
"Hello"
```

PuzzleDB currently supports the Redis command in phases. See [Redis](doc/redis.md) for current support status.
